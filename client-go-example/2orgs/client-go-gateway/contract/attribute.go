package contract

import (
	"client-go-gateway/model"
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

// 增加公有属性
func AddAttribute(contract *client.Contract, attribute *model.Attribute) (string, error) {
	attributeJsonByte, err := json.Marshal(attribute)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}
	result, err := contract.SubmitTransaction("AddAttribute", string(attributeJsonByte))
	if err != nil {
		log.Printf("调用AddAttribute合约失败：%v\n", err)
	}
	return string(result), err
}
func PublishPrivateAttribute(contract *client.Contract, attribute *model.Attribute) (string, error) {
	attributeJsonByte, err := json.Marshal(attribute)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}
	result, err := contract.SubmitTransaction("PublishPrivateAttribute", string(attributeJsonByte))
	if err != nil {
		log.Printf("调用PublishPrivateAttribute合约失败：%v\n", err)
	}
	return string(result), err
}

// 添加私有属性
func BuyPrivateAttribute(contract *client.Contract, request *model.BuyAttributeRequest) (string, error) {
	attributeJsonByte, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}
	result, err := contract.SubmitTransaction("BuyPrivateAttribute", string(attributeJsonByte))
	if err != nil {
		log.Printf("调用BuyPrivateAttribute合约失败：%v\n", err)
	}
	return string(result), err
}

// 根据用户id 查询属性
func FindAttributeByUserId(contract *client.Contract, userId string) ([]model.Attribute, error) {

	result, err := contract.EvaluateTransaction("FindAttributeByUserId", userId)
	if err != nil {
		log.Printf("调用FindAttributeByUserId合约失败：%v\n", err)
		return nil, err
	}
	var attributes []model.Attribute
	json.Unmarshal(result, &attributes)
	return attributes, err
}

//
//func GetSubmittingClientIdentity(contract *client.Contract) string {
//	res, err := contract.EvaluateTransaction("GetSubmittingClientIdentity")
//	if err != nil {
//		log.Printf("调用GetSubmittingClientIdentity合约失败：%v\n", err)
//	}
//	return string(res)
//}
