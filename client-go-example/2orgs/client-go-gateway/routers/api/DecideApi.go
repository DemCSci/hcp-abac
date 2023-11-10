package api

import (
	"client-go-gateway/contract"
	"client-go-gateway/model"
	"client-go-gateway/request"
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"net/http"
	"time"
)

type DecideApi struct {
	respUtil utils.RespUtil
}

func (decideApi *DecideApi) DecideNoRecord2(ctx *gin.Context) {

	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	//now := time.Now()
	contractResponse1, err := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	//fmt.Printf("花费%d us\n", time.Now().Sub(now).Microseconds())
	if err != nil {
		decideApi.respUtil.ErrorResp(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

func (decideApi *DecideApi) DecideNoRecord(ctx *gin.Context) {

	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	//now := time.Now()
	contractResponse1, _ := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	//contractResponse2, _ := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"].Contract, decideRequest)
	//contractResponse3, _ := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"].Contract, decideRequest)
	//contractResponse4, _ := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"].Contract, decideRequest)
	//contractResponse5, _ := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"].Contract, decideRequest)
	//fmt.Printf("花费%d us\n", time.Now().Sub(now).Microseconds())
	//if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
	//	contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
	//	decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
	//	return
	//}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}

	contract.CreateRecord(setting.ClientInfoMap["softMSP"].Contract, *record)
	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

/**
使用携程池的方式异步上链
*/
func (decideApi *DecideApi) DecideNoRecordPool(ctx *gin.Context) {

	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	contractResponse1, _ := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	//contractResponse2, _ := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"].Contract, decideRequest)
	//contractResponse3, _ := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"].Contract, decideRequest)
	//contractResponse4, _ := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"].Contract, decideRequest)
	//contractResponse5, _ := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"].Contract, decideRequest)

	//if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
	//	contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
	//	decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
	//	return
	//}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}
	err = setting.GoroutinePool.Submit(func() {
		contract.CreateRecord(setting.ClientInfoMap["softMSP"].Contract, *record)
		//time.Sleep(time.Millisecond * 100)
	})
	if err != nil {
		setting.MyLogger.Fatal("放入协程池错误")
		decideApi.respUtil.ErrorResp(403, "放入协程池错误", ctx)
	}

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

/**
使用redis消息队列异步上链
*/
func (decideApi *DecideApi) DecideNoRecordRedis(ctx *gin.Context) {

	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	contractResponse1, _ := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	//contractResponse2, _ := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"].Contract, decideRequest)
	//contractResponse3, _ := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"].Contract, decideRequest)
	//contractResponse4, _ := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"].Contract, decideRequest)
	//contractResponse5, _ := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"].Contract, decideRequest)

	//if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
	//	contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
	//	decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
	//	return
	//}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}
	//util.Pool.Submit(func() {
	//	contract.CreateRecord(contract.Contract2, *record)
	//	time.Sleep(time.Second * 1)
	//})

	//contract.CreateRecord(contract.Contract1, *record)
	recordJsonByte, err := json.Marshal(record)

	setting.RedisClient.Publish("test-channel", string(recordJsonByte))
	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

func (decideApi *DecideApi) DecideWithRecord(ctx *gin.Context) {
	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	record, _ := contract.DecideWithRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	decideApi.respUtil.SuccessResp(record, ctx)
}

/**
使用一致性hash，负载均衡请求节点 线程池方法
*/
func (decideApi *DecideApi) DecideHashNoRecordPool(ctx *gin.Context) {
	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	ger, err := setting.GlobalConsistent.Ger(uuid.New())
	//log.Printf("本地请求映射到：%s\n", ger)
	if err != nil {
		setting.MyLogger.Info("一致性hash内部错误,err =", err)
		decideApi.respUtil.ErrorResp(502, "一致性hash内部错误", ctx)
		return
	}
	contractResponse1, err := contract.DecideNoRecord(setting.ClientInfoMap[ger].Contract, decideRequest)
	//contractResponse1 := contract.DecideNoRecord(util.ClientInfoMap["softMSP"], request)
	if err != nil {
		decideApi.respUtil.ErrorResp(403, err.Error(), ctx)
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}
	err = setting.GoroutinePool.Submit(func() {
		//contract.CreateRecord(util.ClientInfoMap["webMSP"], *record)
		contract.CreateRecord(setting.ClientInfoMap[ger].Contract, *record)
		time.Sleep(time.Millisecond * 500)
	})
	if err != nil {
		setting.MyLogger.Fatal("放入协程池错误")
		decideApi.respUtil.ErrorResp(403, "放入协程池错误", ctx)
	}

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

/**
使用一致性hash，负载均衡请求节点 redis方法
*/
func (decideApi *DecideApi) DecideHashNoRecordRedis(ctx *gin.Context) {
	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	ger, err := setting.GlobalConsistent.Ger(uuid.New())
	//log.Printf("本地请求映射到：%s\n", ger)
	if err != nil {
		setting.MyLogger.Info("一致性hash内部错误,err =", err)
		decideApi.respUtil.ErrorResp(502, "一致性hash内部错误", ctx)
		return
	}
	contractResponse1, err := contract.DecideNoRecord(setting.ClientInfoMap[ger].Contract, decideRequest)
	//contractResponse1, err := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
	if err != nil {
		decideApi.respUtil.ErrorResp(403, err.Error(), ctx)
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}

	setting.RedisClient.Publish("test-channel", record)

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}

/**
使用一致性hash，负载均衡请求节点 redis方法
一开始请求单个节点，一旦发现返回的延迟高于阈值，就进行负载均衡
*/
func (decideApi *DecideApi) DecideDynamicHashNoRecordRedis(ctx *gin.Context) {
	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	var contractResponse1 string
	if utils.DynamicHashEnable {
		ger, err := setting.GlobalConsistent.Ger(uuid.New())
		if err != nil {
			setting.MyLogger.Info("一致性hash内部错误,err =", err)
			decideApi.respUtil.ErrorResp(502, "一致性hash内部错误", ctx)
			return
		}
		contractResponse1, err = contract.DecideNoRecord(setting.ClientInfoMap[ger].Contract, decideRequest)
	} else {
		before := time.Now()
		contractResponse1, err = contract.DecideNoRecord(setting.ClientInfoMap["softMSP"].Contract, decideRequest)
		if err != nil {
			decideApi.respUtil.ErrorResp(403, err.Error(), ctx)
		}
		after := time.Now()
		if after.Sub(before).Milliseconds() > 60 {
			utils.DynamicHashEnable = true
		}
	}

	//log.Printf("本地请求映射到：%s\n", ger)

	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}

	setting.RedisClient.Publish("test-channel", record)

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}
