package setting

import (
	"bytes"
	"client-go-gateway/constants"
	"client-go-gateway/model"
	"client-go-gateway/utils"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/ini.v1"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	cfg              *ini.File
	WebSetting       = &WebConfig{}
	logOutSetting    = &LogOutputConfig{}
	MyLogger         = &logrus.Logger{}
	MetadataLogger   *logrus.Logger
	GoroutinePool    *ants.Pool
	redisConfig      = &RedisConfig{}
	RedisClient      *redis.Client
	RateLimitSetting = &RateLimitConfig{}
	ClientInfoMap    = make(map[string]*model.ClientInfo)
	GlobalConsistent = utils.NewConsistent()
)

func Setup() {
	var err error
	cfg, err = ini.Load("setting/my.ini")
	if err != nil {
		fmt.Println("failed while load setting file setting/my.ini,err: ", err)
	}

	mapToConfig("web", WebSetting)

	mapToConfig("log", logOutSetting)

	mapToConfig("redis", redisConfig)

	mapToConfig("rate", RateLimitSetting)

	setupLogOutput()
	setupGoroutinePool()
	// 先临时不使用Redis
	//setupRedis()
}

func setupGoroutinePool() {

	pool, err := ants.NewPool(300, ants.WithNonblocking(false))
	if err != nil {
		log.Fatal("goroutine 池子创建失败")
	}
	GoroutinePool = pool
}

func mapToConfig(section string, value interface{}) {
	err := cfg.Section(section).MapTo(value)
	if err != nil {
		fmt.Println("failed while cfg.MapTo "+section+",err: ", err)
	}
}

type WebConfig struct {
	Port        int
	ContextPath string
}

type MyLogFormatter struct {
}

type LogOutputConfig struct {
	Dir string
}

func (m *MyLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	var ctx *gin.Context
	var traceId = ""
	for k, v := range entry.Data {
		if k == constants.Ctx {
			ctx = v.(*gin.Context)
			continue
		}
		if k == "traceId" {
			traceId = v.(string)
		}
	}

	if ctx != nil {
		ctx.Request.Header.Add(constants.TraceId, traceId)
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog = fmt.Sprintf("%s|%s|%s|%s\n", timestamp, entry.Level, traceId, entry.Message)
	buffer.WriteString(newLog)
	return buffer.Bytes(), nil
}

func setupLogOutput() {
	// 打印请求中业务日志
	MyLogger = initLog(logOutSetting.Dir, "-access.log")
	// 打印请求的元数据信息
	MetadataLogger = initLog(logOutSetting.Dir, "-metadata.log")
}

func initLog(path string, filename string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &MyLogFormatter{}

	filepath := path + filename
	writer, err := rotatelogs.New(
		filepath+".%Y%m%d",
		rotatelogs.WithLinkName(filepath),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	if err != nil {
		fmt.Println("fail to open log file " + filepath)
	}

	log.SetOutput(writer)
	log.Level = logrus.InfoLevel

	return log
}

type RedisConfig struct {
	Host     string
	Port     int
	Database int
}

func setupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
		Password: "",
		DB:       redisConfig.Database,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		panic("redis初始化失败")
	}
}

type RateLimitConfig struct {
	Qps      int
	Interval int64
}
