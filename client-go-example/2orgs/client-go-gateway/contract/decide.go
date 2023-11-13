package contract

import (
	"client-go-gateway/model"
	"client-go-gateway/request"
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

func DecideNoRecord(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.EvaluateTransaction("DecideNoRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用Decide合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideWithRecord(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.SubmitTransaction("DecideWithRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用DecideWithRecord合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func CreateRecord(contract *client.Contract, record model.Record) (string, error) {

	jsonByte, err := json.Marshal(record)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.SubmitTransaction("CreateRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用CreateRecord合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

//测试不同的属性使用

func DecideNoRecord4Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.EvaluateTransaction("DecideNoRecord4Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用Decide合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideWithRecord4Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.SubmitTransaction("DecideWithRecord4Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用DecideWithRecord合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideNoRecord8Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.EvaluateTransaction("DecideNoRecord8Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用Decide合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideWithRecord8Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.SubmitTransaction("DecideWithRecord8Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用DecideWithRecord合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideNoRecord16Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.EvaluateTransaction("DecideNoRecord16Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用Decide合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}

func DecideWithRecord16Attributes(contract *client.Contract, decideRequest request.DecideRequest) (string, error) {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
		return "", err
	}
	result, err := contract.SubmitTransaction("DecideWithRecord16Attributes", string(jsonByte))
	if err != nil {
		log.Printf("调用DecideWithRecord合约失败：%v\n", err)
		return "", err
	}
	return string(result), nil
}
