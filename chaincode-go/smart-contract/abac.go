package abac

import (
	"chaincode-go/model"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

//计算GetStateByRange时的endKey，该函数摘自：github.com/syndtr/goleveldb/leveldb/util
func BytesPrefix(prefix []byte) []byte {
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	return limit
}

//策略决策部分 不写入访问记录
func (s *SmartContract) Decide(ctx contractapi.TransactionContextInterface, request string) (string, error) {
	//1. 查询资源
	//2. 根据资源取回策略
	//3. 查询策略中需要的主体属性和客体属性
	//4. 根据策略进行决策

	// 先来个简易版
	var decideRequest model.DecideRequest
	err := json.Unmarshal([]byte(request), &decideRequest)
	if err != nil {
		return "false", err
	}
	//验证请求者身份
	requesterId := decideRequest.RequesterId
	//验证客户端是否是请求者id
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return "false", err
	}
	if clientID != requesterId {
		return "false", fmt.Errorf("请求者身份验证不通过")
	}

	//查询资源
	//resource, err := s.FindResourceById(ctx, decideRequest.RequesterId)

	attributeMap := make(map[string]interface{})
	//获取主体的公有属性
	publicAttributeStartKey := fmt.Sprintf("attribute:public:%s:", decideRequest.RequesterId)
	publicAttributeEndKey := string(BytesPrefix([]byte(publicAttributeStartKey)))
	publicAttributeRange, err := ctx.GetStub().GetStateByRange(publicAttributeStartKey, publicAttributeEndKey)
	if err != nil {
		return "", fmt.Errorf("获取公有属性失败")
	}
	defer publicAttributeRange.Close()
	for publicAttributeRange.HasNext() {
		result, err := publicAttributeRange.Next()
		if err != nil {
			return "", fmt.Errorf("错误")
		}
		var attribute model.Attribute
		json.Unmarshal(result.Value, &attribute)
		attributeMap[attribute.Key] = attribute.Value
	}
	//获取私有属性 attribute:private:userid:resourceid:key
	privateAttributeStartKey := fmt.Sprintf("attribute:private:%s:%s:", decideRequest.RequesterId, decideRequest.ResourceId)
	privateAttributeEndKey := string(BytesPrefix([]byte(publicAttributeStartKey)))
	privateAttributeRange, err := ctx.GetStub().GetStateByRange(privateAttributeStartKey, privateAttributeEndKey)
	if err != nil {
		return "", fmt.Errorf("获取公有属性失败")
	}
	defer privateAttributeRange.Close()
	for privateAttributeRange.HasNext() {
		result, err := privateAttributeRange.Next()
		if err != nil {
			return "", fmt.Errorf("错误")
		}
		var attribute model.Attribute
		json.Unmarshal(result.Value, &attribute)
		attributeMap[attribute.Key] = attribute.Value
	}

	//根据资源查找对应的策略
	//这里暂时写死一个资源的策略
	if attributeMap["age"] == "40" && attributeMap["occupation"] == "doctor" {
		return "true", nil
	}
	return "false", nil

}

//策略决策部分 写入访问记录
func (s *SmartContract) DecideWithRecord(ctx contractapi.TransactionContextInterface, request string) (string, error) {
	//1. 查询资源
	//2. 根据资源取回策略
	//3. 查询策略中需要的主体属性和客体属性
	//4. 根据策略进行决策
	//5. 写入访问记录

	// 先来个简易版
	var decideRequest model.DecideRequest
	err := json.Unmarshal([]byte(request), &decideRequest)
	if err != nil {
		return "false", err
	}
	//验证请求者身份
	requesterId := decideRequest.RequesterId
	//验证客户端是否是请求者id
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return "false", err
	}
	if clientID != requesterId {
		return "false", fmt.Errorf("请求者身份验证不通过")
	}

	//查询资源
	//resource, err := s.FindResourceById(ctx, decideRequest.RequesterId)
	//根据资源查找对应的策略
	//这里暂时写死一个资源的策略
	//获取主体的属性
	//subject, err := s.FindUserById(ctx, clientID)
	//attributeMap := make(map[string]interface{})
	//for i := 0; i < len(subject.Attributes); i++ {
	//	attributeMap[subject.Attributes[i].Key] = subject.Attributes[i].Value
	//}
	//
	//if attributeMap["age"] == "40" && attributeMap["occupation"] == "doctor" {
	//	var record model.Record
	//	record.Id = decideRequest.Id
	//	record.RequesterId = decideRequest.RequesterId
	//	record.ResourceId = decideRequest.ResourceId
	//	record.Response = "true"
	//	recordJsonAsByte, _ := json.Marshal(record)
	//	s.CreateRecord(ctx, string(recordJsonAsByte))
	//	return "true", nil
	//}
	return "false", nil

}

//访问记录相关
func (s *SmartContract) CreateRecord(ctx contractapi.TransactionContextInterface, request string) error {

	var record model.Record
	err := json.Unmarshal([]byte(request), &record)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(record.Id, []byte(request))
}
