package abac

import (
	"chaincode-go/model"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// 用户相关

// CreateUser issues a new user to the world state with given details.
func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, request string) error {

	// Get ID of submitting client identity
	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return err
	}
	exists, err := s.UserExists(ctx, clientID)

	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the user %s already exists", clientID)
	}

	var user model.User
	err = json.Unmarshal([]byte(request), &user)
	if err != nil {
		return err
	}

	user.ID = clientID

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(clientID, userJSON)
}

func (s *SmartContract) DeleteUser(ctx contractapi.TransactionContextInterface) error {

	// Get ID of submitting client identity
	clientID, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return err
	}

	exists, err := s.UserExists(ctx, clientID)

	if err != nil {
		return err
	}
	if exists {
		ctx.GetStub().DelState(clientID)
		return nil
	}

	return fmt.Errorf("用户不存在")
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

func (s *SmartContract) FindUserById(ctx contractapi.TransactionContextInterface, userId string) (*model.User, error) {

	userAsByte, err := ctx.GetStub().GetState(userId)
	if err != nil {
		return nil, fmt.Errorf("查询资源失败")
	}
	var user model.User
	err = json.Unmarshal(userAsByte, &user)

	return &user, err
}

// returns all users found in world state
func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]*model.User, error) {

	startKey := "user:"
	endKey := string(BytesPrefix([]byte(startKey)))
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var users []*model.User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var user model.User

		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (s *SmartContract) GetUserHistory(ctx contractapi.TransactionContextInterface) ([]string, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	clientIdentity, err := s.GetSubmittingClientIdentity(ctx)
	if err != nil {
		return nil, err
	}
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(clientIdentity)
	var result []string
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		resJson, err := json.Marshal(queryResponse)
		result = append(result, string(resJson))
	}
	return result, nil
}

// UserExists returns true when asset with given ID exists in world state
func (s *SmartContract) UserExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return userJSON != nil, nil
}

// GetSubmittingClientIdentity returns the name and issuer of the identity that
// invokes the smart contract. This function base64 decodes the identity string
// before returning the value to the client or smart contract.
func (s *SmartContract) GetSubmittingClientIdentity(ctx contractapi.TransactionContextInterface) (string, error) {
	b64ID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	decodeID, err := base64.StdEncoding.DecodeString(b64ID)
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode clientID: %v", err)
	}
	return "user:" + string(decodeID), nil
}

//属性相关

// 增加公有属性
func (s *SmartContract) AddAttribute(ctx contractapi.TransactionContextInterface, request string) error {

	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return err
	}

	var attribute model.Attribute
	err = json.Unmarshal([]byte(request), &attribute)
	if err != nil {
		return err
	}

	//判断是否 增加公有属性 还是私有属性
	if attribute.Type == "PUBLIC" {
		//公有属性 直接放进去
		userJson, err := ctx.GetStub().GetState(clientID)
		if err != nil {
			return err
		}
		var user model.User
		err = json.Unmarshal(userJson, &user)
		if err != nil {
			return err
		}
		user.Attributes = append(user.Attributes, attribute)
		newUser, err := json.Marshal(user)
		ctx.GetStub().PutState(clientID, newUser)
	} else {
		//私有属性
		//1.
		return fmt.Errorf("add private attribute")
	}
	return err
}

// 发布私有属性
func (s *SmartContract) PublishPrivateAttribute(ctx contractapi.TransactionContextInterface, request string) error {

	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return err
	}

	var attribute model.Attribute
	err = json.Unmarshal([]byte(request), &attribute)
	if err != nil {
		return err
	}

	//判断是否 增加公有属性 还是私有属性
	if attribute.Type == "PRIVATE" {
		// 判断该资源是否存在
		exist, err2 := s.ResourceExists(ctx, attribute.ResourceId)
		if err2 != nil {
			return err2
		}
		if !exist {
			return fmt.Errorf("要增加的私有属性 该资源不存在")
		}
		//直接放进去
		attribute.Owner = clientID
		attributeAsJsonByte, err := json.Marshal(attribute)
		if err != nil {
			return err
		}
		ctx.GetStub().PutState(attribute.Id, attributeAsJsonByte)
	} else {
		//公有属性
		return fmt.Errorf("can only add private attribute")
	}
	return err
}

//查询属性
func (s *SmartContract) FindAttributeById(ctx contractapi.TransactionContextInterface, attributeId string) (*model.Attribute, error) {

	attributeAsByte, err := ctx.GetStub().GetState(attributeId)
	if err != nil {
		return nil, fmt.Errorf("查询属性失败")
	}
	var attribute model.Attribute
	err = json.Unmarshal(attributeAsByte, &attribute)

	return &attribute, err
}

//查看某资源对应的 可买属性
func (s *SmartContract) FindAttributeByResourceId(ctx contractapi.TransactionContextInterface, resourceId string) ([]*model.Attribute, error) {

	startKey := "attribute:" + resourceId + ":"
	endKey := string(BytesPrefix([]byte(startKey)))
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var attributes []*model.Attribute
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var attribute model.Attribute

		err = json.Unmarshal(queryResponse.Value, &attribute)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, &attribute)
	}

	return attributes, nil

}

// 购买私有属性
func (s *SmartContract) BuyPrivateAttribute(ctx contractapi.TransactionContextInterface, request string) error {

	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return err
	}

	var buyAttributeRequest model.BuyAttributeRequest
	err = json.Unmarshal([]byte(request), &buyAttributeRequest)
	if err != nil {
		return err
	}
	//验证买家的身份
	if clientID != buyAttributeRequest.Buyer {
		return fmt.Errorf("您的身份不匹配")
	}

	//查找出 买家 和卖家
	buyer, err := s.FindUserById(ctx, clientID)
	//
	seller, err := s.FindUserById(ctx, buyAttributeRequest.Seller)

	//根据属性id 查询属性
	attribute, err := s.FindAttributeById(ctx, buyAttributeRequest.AttributeId)

	//验证该属性的资源是否是卖家拥有
	resource, err := s.FindResourceById(ctx, attribute.ResourceId)

	if seller.ID != resource.Owner {
		return fmt.Errorf("该资源不是该卖家所有")
	}

	//转账
	seller.Money += attribute.Money
	buyer.Money -= attribute.Money

	buyer.Attributes = append(seller.Attributes, *attribute)

	//分别存储buyer 和seller
	buyerAsJsonByte, err := json.Marshal(buyer)
	sellerAsJsonByte, err := json.Marshal(seller)

	ctx.GetStub().PutState(seller.ID, sellerAsJsonByte)
	ctx.GetStub().PutState(buyer.ID, buyerAsJsonByte)
	return err
}

/**
资源
*/
func (s *SmartContract) CreateResource(ctx contractapi.TransactionContextInterface, request string) error {

	//// Get ID of submitting client identity
	clientID, err := s.GetSubmittingClientIdentity(ctx)

	if err != nil {
		return fmt.Errorf("获取用户id失败")
	}

	var resource model.Resource
	err = json.Unmarshal([]byte(request), &resource)

	resource.Owner = clientID
	exists, err := s.ResourceExists(ctx, resource.Id)

	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the resource %s already exists", resource.Id)
	}

	resourceJSON, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(resource.Id, resourceJSON)
}

//查询资源
func (s *SmartContract) FindResourceById(ctx contractapi.TransactionContextInterface, resourceId string) (*model.Resource, error) {

	resourceAsByte, err := ctx.GetStub().GetState(resourceId)
	if err != nil {
		return nil, fmt.Errorf("查询资源失败")
	}
	var resource model.Resource
	err = json.Unmarshal(resourceAsByte, &resource)

	return &resource, err
}

func (s *SmartContract) ResourceExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	resource, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return resource != nil, nil
}

// returns all resource found in world state
func (s *SmartContract) GetAllResource(ctx contractapi.TransactionContextInterface) ([]*model.Resource, error) {

	startKey := "resource:"
	endKey := string(BytesPrefix([]byte(startKey)))
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var resources []*model.Resource
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var resource model.Resource

		err = json.Unmarshal(queryResponse.Value, &resource)
		if err != nil {
			return nil, err
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}

//策略决策部分
func (s *SmartContract) Decide(ctx contractapi.TransactionContextInterface, request string) ([]*model.Resource, error) {

	startKey := "resource:"
	endKey := string(BytesPrefix([]byte(startKey)))
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var resources []*model.Resource
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var resource model.Resource

		err = json.Unmarshal(queryResponse.Value, &resource)
		if err != nil {
			return nil, err
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}
