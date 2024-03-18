package abac

import (
	"chaincode-go/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/**
创建策略
*/
func (s *SmartContract) CreatePolicy(ctx contractapi.TransactionContextInterface, request string) error {

	// Get ID of submitting client identity
	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return fmt.Errorf("获取用户id失败")
	}

	var policy model.Policy
	err = json.Unmarshal([]byte(request), &policy)
	resource, err := s.FindResourceById(ctx, policy.ResourceId)
	if err != nil {
		return fmt.Errorf("获取资源失败,err =%v", err)
	}
	if resource.Owner != clientID {
		return fmt.Errorf("创建者不是资源的所有人")
	}
	// 生成真正的策略id，一个资源一个
	policy.Id = fmt.Sprintf("policy:%s", policy.ResourceId)

	//持久化
	policyJsonBytes, err := json.Marshal(policy)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(policy.Id, policyJsonBytes)

}

//查询策略 外部不要调用
func (s *SmartContract) findPolicyById(ctx contractapi.TransactionContextInterface, policyId string) (*model.Policy, error) {

	policyAsByte, err := ctx.GetStub().GetState(policyId)
	if err != nil {
		return nil, fmt.Errorf("查询策略失败")
	}
	var policy model.Policy
	err = json.Unmarshal(policyAsByte, &policy)

	return &policy, err
}

//查询策略 外部不要调用
func (s *SmartContract) FindPolicyById(ctx contractapi.TransactionContextInterface, policyId string) (string, error) {

	policyAsByte, err := ctx.GetStub().GetState(policyId)
	if err != nil {
		return "", fmt.Errorf("查询策略失败")
	}
	return string(policyAsByte), err
}

//删除策略
func (s *SmartContract) DeletePolicyById(ctx contractapi.TransactionContextInterface, policyId string) (string, error) {

	err := ctx.GetStub().DelState(policyId)
	if err != nil {
		return "删除策略失败", fmt.Errorf("删除策略失败")
	}

	return "删除策略成功", nil
}
func (s *SmartContract) PolicyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	policy, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return policy != nil, nil
}