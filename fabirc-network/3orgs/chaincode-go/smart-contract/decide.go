package abac

import (
	"chaincode-go/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

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
func (s *SmartContract) DecideNoRecord(ctx contractapi.TransactionContextInterface, request string) (string, error) {
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
	//log.Printf("decideRequest:%v", decideRequest)
	//查询资源
	resourceAsByte, err := ctx.GetStub().GetState(decideRequest.ResourceId)
	if err != nil {
		return "查询资源失败", fmt.Errorf("查询资源失败")
	}
	var resource model.Resource
	err = json.Unmarshal(resourceAsByte, &resource)
	//resource, err := s.FindResourceById(ctx, decideRequest.ResourceId)
	//log.Printf("resource:%v", resource)
	//log.Printf("resource2:%v", *resource)
	//log.Printf("resource controllers:%v", resource.Controllers)
	//验证资源控制器身份
	controllerId, err := s.GetSubmittingClientIdentity(ctx)
	controllers := resource.Controllers
	i := 0
	for i < len(controllers) {
		log.Printf("controllerId: %v", controllerId)
		log.Printf("currentId: %v", controllers[i])
		if controllers[i] == controllerId {
			break
		}
		i++
	}
	if i == len(controllers) {
		return "false", fmt.Errorf("资源控制器身份验证不通过")
	}
	//验证请求者身份

	//requesterId, err := s.FindUserById(ctx, decideRequest.RequesterId)
	//if err != nil {
	//	return "false", err
	//}

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

	// 先来个简易版
	var decideRequest model.DecideRequest
	err := json.Unmarshal([]byte(request), &decideRequest)
	if err != nil {
		return "false", err
	}

	//查询资源
	//查询资源
	resourceAsByte, err := ctx.GetStub().GetState(decideRequest.ResourceId)
	if err != nil {
		return "查询资源失败", fmt.Errorf("查询资源失败")
	}
	var resource model.Resource
	err = json.Unmarshal(resourceAsByte, &resource)

	//验证资源控制器身份
	controllerId, err := s.GetSubmittingClientIdentity(ctx)
	controllers := resource.Controllers
	i := 0
	for i < len(controllers) {
		if controllers[i] == controllerId {
			break
		}
		i++
	}
	if i == len(controllers) {
		return "false", fmt.Errorf("资源控制器身份验证不通过")
	}
	//验证请求者身份

	//requesterId, err := s.FindUserById(ctx, decideRequest.RequesterId)
	//if err != nil {
	//	return "false", err
	//}

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
		var record model.Record
		record.Id = decideRequest.Id
		record.RequesterId = decideRequest.RequesterId
		record.ResourceId = decideRequest.ResourceId
		record.Response = "true"
		recordJsonAsByte, _ := json.Marshal(record)
		err := s.CreateRecord(ctx, string(recordJsonAsByte))
		if err != nil {
			return "", err
		}
		return "true", nil
	}

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
