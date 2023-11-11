package contract

import (
	"client-go-gateway/model"
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

// GetAllResource 获取所有的资源
func GetAllResource(contract *client.Contract) (*[]model.Resource, error) {
	result, err := contract.EvaluateTransaction("GetAllResource")
	//resourceJson := string(result)
	var resourceArr []model.Resource
	err = json.Unmarshal(result, &resourceArr)

	return &resourceArr, err
}

// 添加资源
func AddResource(contract *client.Contract, resource *model.Resource) string {

	resourceJsonByte, err2 := json.Marshal(resource)
	if err2 != nil {
		log.Fatalf("序列化失败：%v\n", err2)
	}
	result, err := contract.SubmitTransaction("CreateResource", string(resourceJsonByte))
	if err != nil {
		log.Printf("调用CreateResource合约失败：%v\n", err)
	}
	return string(result)
}

// 添加资源控制器
func AddResourceController(contract *client.Contract, resource *model.AddResourceControllerRequest) string {
	controllerJsonByte, err2 := json.Marshal(resource)
	if err2 != nil {
		log.Fatalf("序列化失败：%v\n", err2)
	}
	result, err := contract.SubmitTransaction("AddResourceController", string(controllerJsonByte))
	if err != nil {
		log.Printf("调用AddResourceController合约失败：%v\n", err)
	}
	return string(result)
}
