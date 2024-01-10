package api

import (
	"client-go-gateway/contract"
	"client-go-gateway/model"
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

type AttributeApi struct {
	respUtil utils.RespUtil
}

// 根据用户id 查询该用户的所有属性
func (attributeAPi *AttributeApi) FindAttributeByUserId(ctx *gin.Context) {
	userId, _ := ctx.GetQuery("user_id")

	if len(userId) == 0 {
		setting.MyLogger.Info("userId为空")
		attributeAPi.respUtil.ErrorResp(502, "userId为空", ctx)
		return
	}
	attributes, err := contract.FindAttributeByUserId(setting.ClientInfoMap["org1MSP"].Contract, userId)
	if err != nil {
		attributeAPi.respUtil.ErrorResp(502, "获取用户属性错误", ctx)
		return
	}
	attributeAPi.respUtil.SuccessResp(attributes, ctx)
}

// 增加属性
func (attributeAPi *AttributeApi) AddAttribute(ctx *gin.Context) {

	var attribute model.Attribute
	err := ctx.BindJSON(&attribute)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		attributeAPi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	uuid, err := utils.GenerateUUID()
	attribute.Id = fmt.Sprintf("attribute:%s", uuid)
	attribute.Type = "PUBLIC"

	attributes, err := contract.AddAttribute(setting.ClientInfoMap["org1MSP"].Contract, &attribute)
	if err != nil {
		attributeAPi.respUtil.ErrorResp(502, "获取用户属性错误", ctx)
		return
	}
	attributeAPi.respUtil.SuccessResp(attributes, ctx)
}

var endorseTransaction *client.Transaction

// 增加属性 只进行背书，不提交到orderer节点
func (attributeAPi *AttributeApi) AddAttributeOnlyPeer(ctx *gin.Context) {

	var attribute model.Attribute
	err := ctx.BindJSON(&attribute)
	if err != nil {
		setting.MyLogger.Info("传入信息错误,err =", err)
		attributeAPi.respUtil.IllegalArgumentErrorResp("传入信息错误", ctx)
		return
	}
	uuid, err := utils.GenerateUUID()
	attribute.Id = fmt.Sprintf("attribute:%s", uuid)
	attribute.Type = "PUBLIC"
	attributeJsonByte, err := json.Marshal(attribute)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}

	proposal, err := setting.ClientInfoMap["org1MSP"].
		Contract.NewProposal("AddAttribute", client.WithArguments(string(attributeJsonByte)))
	if err != nil {
		attributeAPi.respUtil.ErrorResp(502, "生成proposal错误", ctx)
		return
	}
	// 背书和提交给orderer分离
	endorseTransaction, err = proposal.Endorse()
	if err != nil {
		attributeAPi.respUtil.ErrorResp(502, "背书错误", ctx)
		return
	}
	log.Printf("背书结果，res=%v", string(endorseTransaction.Result()))

	if err != nil {
		fmt.Println("反序列化错误")
		return
	}

	attributeAPi.respUtil.SuccessResp("ok1111", ctx)
}

// 增加属性 只提交到orderer节点
func (attributeAPi *AttributeApi) AddAttributeOnlyOrderer(ctx *gin.Context) {

	// 背书和提交给orderer分离
	comit, err := endorseTransaction.Submit()
	if err != nil {
		attributeAPi.respUtil.ErrorResp(502, "转发给orderer错误", ctx)
		return
	}
	log.Printf("排序结果，res=%v", comit.TransactionID())
	attributeAPi.respUtil.SuccessResp("ok1111", ctx)
}
