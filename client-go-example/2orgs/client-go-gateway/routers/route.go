package routers

import (
	"client-go-gateway/middleware"
	"client-go-gateway/routers/api"
	"github.com/gin-gonic/gin"
)

var (
	helloAPi  api.HelloApi
	userApi   api.UserApi
	decideApi api.DecideApi
	toolAPi   api.ToolApi
)

func InitRouter(contextPath string) *gin.Engine {
	router := gin.New()

	router.Use(middleware.HandleError)
	router.Use(middleware.EnableTraceIdHook)
	//router.Use(middleware.RecordCostTime())

	currency := router.Group(contextPath + "/hello")
	currency.GET("", helloAPi.Hello)

	user := router.Group(contextPath + "/user")
	{
		user.GET("/users", userApi.GetAllUsers)
		user.POST("/add", userApi.AddUser)
		user.POST("/addAllUser", userApi.AddAllUser)
		user.GET("/my", userApi.GetSubmittingClientIdentity)
	}

	decide := router.Group(contextPath + "/decide")
	{
		decide.POST("/decideNoRecord", decideApi.DecideNoRecord)
		decide.POST("/decideNoRecord2", decideApi.DecideNoRecord2)
		decide.POST("/decideNoRecordPool", decideApi.DecideNoRecordPool)
		decide.POST("/decideNoRecordRedis", decideApi.DecideNoRecordRedis)
		decide.POST("/decideWithRecord", decideApi.DecideWithRecord)

		decide.POST("/DecideHashNoRecordPool", decideApi.DecideHashNoRecordPool)
		decide.POST("/DecideHashNoRecordRedis", decideApi.DecideHashNoRecordRedis)
		decide.POST("/DecideDynamicHashNoRecordRedis", decideApi.DecideDynamicHashNoRecordRedis)

		//测试不同的属性
		decide.POST("/DecideWithRecord4Attributes", decideApi.DecideWithRecord4Attributes)
		decide.POST("/DecideNoRecord4Attributes", decideApi.DecideNoRecord4Attributes)
		decide.POST("/DecideWithRecord8Attributes", decideApi.DecideWithRecord8Attributes)
		decide.POST("/DecideNoRecord8Attributes", decideApi.DecideNoRecord8Attributes)
		decide.POST("/DecideWithRecord16Attributes", decideApi.DecideWithRecord16Attributes)
		decide.POST("/DecideNoRecord16Attributes", decideApi.DecideNoRecord16Attributes)
	}
	tool := router.Group(contextPath + "/tool")
	{
		tool.GET("/init", toolAPi.Init)
		tool.GET("/init4", toolAPi.Init4)
		tool.GET("/init8", toolAPi.Init8)
		tool.GET("/init16", toolAPi.Init16)
	}
	return router
}
