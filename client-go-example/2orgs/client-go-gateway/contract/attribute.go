package contract

import (
	"client-go-gateway/model"
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

type Attribute struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	ResourceId string  `json:"resourceId"`
	Owner      string  `json:"ownerId"`
	Key        string  `json:"key"`
	Value      string  `json:"value"`
	NotBefore  string  `json:"notBefore"`
	NotAfter   string  `json:"notAfter"`
	Money      float64 `json:"money"`
}

// 增加公有属性
func AddAttribute(contract *client.Contract, attribute *Attribute) (string, error) {
	attributeJsonByte, err := json.Marshal(attribute)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}
	result, err := contract.SubmitTransaction("AddAttribute", string(attributeJsonByte))
	if err != nil {
		log.Printf("调用AddUser合约失败：%v\n", err)
	}
	return string(result), err
}
func PublishPrivateAttribute(contract *client.Contract, attribute *Attribute) (string, error) {
	attributeJsonByte, err := json.Marshal(attribute)
	if err != nil {
		log.Fatalf("序列化失败：%v\n", err)
	}
	result, err := contract.SubmitTransaction("PublishPrivateAttribute", string(attributeJsonByte))
	if err != nil {
		log.Printf("调用AddUser合约失败：%v\n", err)
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
		log.Printf("调用AddUser合约失败：%v\n", err)
	}
	return string(result), err
}

//
//func GetSubmittingClientIdentity(contract *client.Contract) string {
//	res, err := contract.EvaluateTransaction("GetSubmittingClientIdentity")
//	if err != nil {
//		log.Printf("调用GetSubmittingClientIdentity合约失败：%v\n", err)
//	}
//	return string(res)
//}
