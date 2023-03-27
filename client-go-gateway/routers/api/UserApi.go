package api

import (
	"client-go-gateway/contract"
	"client-go-gateway/request"
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

type UserApi struct {
	respUtil utils.RespUtil
}

func (userAPi *UserApi) GetAllUsers(ctx *gin.Context) {
	ger, err := setting.GlobalConsistent.Ger(uuid.New())
	//log.Printf("本地请求映射到：%s\n", ger)
	if err != nil {
		setting.MyLogger.Info("一致性hash内部错误,err =", err)
		userAPi.respUtil.ErrorResp(502, "一致性hash内部错误", ctx)
		return
	}
	//allUsers, err := contract.GetAllUsers(util.ClientInfoMap["org4MSP"])
	allUsers, err := contract.GetAllUsers(setting.ClientInfoMap[ger])
	users := allUsers
	if err != nil {
		userAPi.respUtil.ErrorResp(500, err.Error(), ctx)
		return
	}
	userAPi.respUtil.SuccessResp(users, ctx)
}

func (userAPi *UserApi) AddUser(ctx *gin.Context) {
	var addUserRequest request.AddUserRequest
	err := ctx.BindJSON(&addUserRequest)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		userAPi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	record := contract.AddUser(setting.ClientInfoMap[addUserRequest.MspID])
	userAPi.respUtil.SuccessResp(record, ctx)
}

func (userAPi *UserApi) GetSubmittingClientIdentity(ctx *gin.Context) {

	identity := contract.GetSubmittingClientIdentity(setting.ClientInfoMap["softMSP"])
	userAPi.respUtil.SuccessResp(identity, ctx)
}

func (userAPi *UserApi) AddAllUser(ctx *gin.Context) {

	for _, val := range setting.ClientInfoMap {
		contract.AddUser(val)
	}
	userAPi.respUtil.SuccessResp("全部用户身份注册成功", ctx)
}
