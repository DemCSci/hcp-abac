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
	"time"
)

type DecideApi struct {
	respUtil utils.RespUtil
}

func (decideApi *DecideApi) DecideNoRecord(ctx *gin.Context) {

	var decideRequest request.DecideRequest
	err := ctx.BindJSON(&decideRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		decideApi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	contractResponse1 := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"], decideRequest)
	contractResponse2 := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"], decideRequest)
	contractResponse3 := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"], decideRequest)
	contractResponse4 := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"], decideRequest)
	contractResponse5 := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"], decideRequest)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
		return
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}

	contract.CreateRecord(setting.ClientInfoMap["softMSP"], *record)
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
	contractResponse1 := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"], decideRequest)
	contractResponse2 := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"], decideRequest)
	contractResponse3 := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"], decideRequest)
	contractResponse4 := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"], decideRequest)
	contractResponse5 := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"], decideRequest)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
		return
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}
	err = setting.GoroutinePool.Submit(func() {
		contract.CreateRecord(setting.ClientInfoMap["softMSP"], *record)
		time.Sleep(time.Second * 1)
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
	contractResponse1 := contract.DecideNoRecord(setting.ClientInfoMap["softMSP"], decideRequest)
	contractResponse2 := contract.DecideNoRecord(setting.ClientInfoMap["webMSP"], decideRequest)
	contractResponse3 := contract.DecideNoRecord(setting.ClientInfoMap["hardMSP"], decideRequest)
	contractResponse4 := contract.DecideNoRecord(setting.ClientInfoMap["org4MSP"], decideRequest)
	contractResponse5 := contract.DecideNoRecord(setting.ClientInfoMap["org5MSP"], decideRequest)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		decideApi.respUtil.ErrorResp(403, "结果不一致，拒绝该请求", ctx)
		return
	}
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
	record := contract.DecideWithRecord(setting.ClientInfoMap["softMSP"], decideRequest)
	decideApi.respUtil.SuccessResp(record, ctx)
}

/**
使用一致性hash，负载均衡请求节点
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
	contractResponse1 := contract.DecideNoRecord(setting.ClientInfoMap[ger], decideRequest)
	//contractResponse1 := contract.DecideNoRecord(util.ClientInfoMap["softMSP"], request)

	//异步发送record
	record := &model.Record{
		Id:          "record:" + decideRequest.ResourceId + ":" + decideRequest.RequesterId + ":" + uuid.New(),
		RequesterId: decideRequest.RequesterId,
		ResourceId:  decideRequest.RequesterId,
		Response:    contractResponse1,
	}
	err = setting.GoroutinePool.Submit(func() {
		//contract.CreateRecord(util.ClientInfoMap["webMSP"], *record)
		contract.CreateRecord(setting.ClientInfoMap[ger], *record)
		time.Sleep(time.Millisecond * 500)
	})
	if err != nil {
		setting.MyLogger.Fatal("放入协程池错误")
		decideApi.respUtil.ErrorResp(403, "放入协程池错误", ctx)
	}

	decideApi.respUtil.SuccessResp(contractResponse1, ctx)
}
