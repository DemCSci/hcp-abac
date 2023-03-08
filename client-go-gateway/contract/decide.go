package contract

import (
	"client-go-gateway/model"
	"client-go-gateway/request"
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

func DecideNoRecord(contract *client.Contract, decideRequest request.DecideRequest) string {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
	}
	result, err := contract.EvaluateTransaction("DecideNoRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用Decide合约失败：%v\n", err)
	}
	return string(result)
}

func DecideWithRecord(contract *client.Contract, decideRequest request.DecideRequest) string {

	jsonByte, err := json.Marshal(decideRequest)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
	}
	result, err := contract.SubmitTransaction("DecideWithRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用DecideWithRecord合约失败：%v\n", err)

	}
	return string(result)
}

func CreateRecord(contract *client.Contract, record model.Record) string {

	jsonByte, err := json.Marshal(record)
	if err != nil {
		log.Println("序列化失败,%v\n", err)
	}
	result, err := contract.SubmitTransaction("CreateRecord", string(jsonByte))
	if err != nil {
		log.Printf("调用CreateRecord合约失败：%v\n", err)

	}
	return string(result)
}