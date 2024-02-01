package abac

import (
	"chaincode-go/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

//策略决策部分 不写入访问记录
func (s *SmartContract) DecideNoRecordWithPolicy(ctx contractapi.TransactionContextInterface, request string) (string, error) {
	//1. 查询资源
	//2. 根据资源取回策略
	//3. 查询策略中需要的主体属性和客体属性
	//4. 根据策略进行决策

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

	//验证资源控制器身份
	controllerId, err := s.GetSubmittingClientIdentity(ctx)
	controllers := resource.Controllers
	i := 0
	for i < len(controllers) {
		//log.Printf("controllerId: %v", controllerId)
		//log.Printf("currentId: %v", controllers[i])
		if controllers[i] == controllerId {
			log.Printf("当前资源控制器身份验证通过")
			break
		}
		i++
	}
	if i == len(controllers) {
		//log.Printf("i=%v\n", i)
		return "false", fmt.Errorf("资源控制器身份验证不通过")
	}
	policyId := fmt.Sprintf("policy:%s", resource.Id)
	// 获取资源的策略
	policy, err := s.findPolicyById(ctx, policyId)
	if err != nil || policy == nil {
		return "false", fmt.Errorf("获取资源的相关策略失败，err=%v", err)
	}

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
	//判断访问者的公有属性是否满足策略
	for index, key := range policy.RequesterAttributeKeys {
		if attributeMap[key] != policy.RequesterAttributeValues[index] {
			return "false", fmt.Errorf("访问者公有属性不满足策略")
		}
	}

	// 判断是否是同一个组织
	requester, err := s.FindUserById(ctx, decideRequest.RequesterId)
	if err != nil {
		return "false", fmt.Errorf("查找请求者身份失败")
	}
	resurceOwner, err := s.FindUserById(ctx, resource.Owner)
	if requester.Org != resurceOwner.Org {
		//获取私有属性 attribute:private:userid:resourceid:key
		privateAttributeStartKey := fmt.Sprintf("attribute:private:%s:%s:", decideRequest.RequesterId, decideRequest.ResourceId)
		privateAttributeEndKey := string(BytesPrefix([]byte(publicAttributeStartKey)))
		privateAttributeRange, err := ctx.GetStub().GetStateByRange(privateAttributeStartKey, privateAttributeEndKey)
		if err != nil {
			return "", fmt.Errorf("获取私有属性失败")
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
		//判断访问者的私有属性是否满足策略
		for index, key := range policy.PrivateKeys {
			if attributeMap[key] != policy.PrivateValues[index] {
				return "false", fmt.Errorf("访问者私有属性不满足策略")
			}
		}
	}

	return "true", nil

}

//策略决策部分 写入访问记录
func (s *SmartContract) DecideWithRecordWithPolicy(ctx contractapi.TransactionContextInterface, request string) (string, error) {
	//1. 查询资源
	//2. 根据资源取回策略
	//3. 查询策略中需要的主体属性和客体属性
	//4. 根据策略进行决策

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

	//验证资源控制器身份
	controllerId, err := s.GetSubmittingClientIdentity(ctx)
	controllers := resource.Controllers
	i := 0
	for i < len(controllers) {
		//log.Printf("controllerId: %v", controllerId)
		//log.Printf("currentId: %v", controllers[i])
		if controllers[i] == controllerId {
			log.Printf("当前资源控制器身份验证通过")
			break
		}
		i++
	}
	if i == len(controllers) {
		//log.Printf("i=%v\n", i)
		return "false", fmt.Errorf("资源控制器身份验证不通过")
	}
	policyId := fmt.Sprintf("policy:%s", resource.Id)
	// 获取资源的策略
	policy, err := s.findPolicyById(ctx, policyId)
	if err != nil || policy == nil {
		return "false", fmt.Errorf("获取资源的相关策略失败，err=%v", err)
	}

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
	//判断访问者的公有属性是否满足策略
	for index, key := range policy.RequesterAttributeKeys {
		if attributeMap[key] != policy.RequesterAttributeValues[index] {
			return "false", fmt.Errorf("访问者公有属性不满足策略")
		}
	}

	// 判断是否是同一个组织
	requester, err := s.FindUserById(ctx, decideRequest.RequesterId)
	if err != nil {
		return "false", fmt.Errorf("查找请求者身份失败")
	}
	resurceOwner, err := s.FindUserById(ctx, resource.Owner)
	if requester.Org != resurceOwner.Org {
		//获取私有属性 attribute:private:userid:resourceid:key
		privateAttributeStartKey := fmt.Sprintf("attribute:private:%s:%s:", decideRequest.RequesterId, decideRequest.ResourceId)
		privateAttributeEndKey := string(BytesPrefix([]byte(publicAttributeStartKey)))
		privateAttributeRange, err := ctx.GetStub().GetStateByRange(privateAttributeStartKey, privateAttributeEndKey)
		if err != nil {
			return "", fmt.Errorf("获取私有属性失败")
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
		//判断访问者的私有属性是否满足策略
		for index, key := range policy.PrivateKeys {
			if attributeMap[key] != policy.PrivateKeys[index] {
				return "false", fmt.Errorf("访问者私有属性不满足策略")
			}
		}
	}
	var record model.Record
	record.Id = decideRequest.Id
	record.RequesterId = decideRequest.RequesterId
	record.ResourceId = decideRequest.ResourceId
	record.Response = "true"
	recordJsonAsByte, _ := json.Marshal(record)
	err = s.CreateRecord(ctx, string(recordJsonAsByte))
	if err != nil {
		return "", err
	}
	return "true", nil

}
