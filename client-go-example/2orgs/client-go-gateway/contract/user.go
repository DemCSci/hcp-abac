package contract

import (
	"encoding/json"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

type User struct {
	Id    string  `json:"id"`
	Money float64 `json:"money"`
}

func GetAllUsers(contract *client.Contract) (string, error) {
	result, err := contract.EvaluateTransaction("GetAllUsers")

	return string(result), err
}

func AddUser(contract *client.Contract) string {
	user := &User{
		Id:    "",
		Money: 200,
	}
	userJsonByte, err2 := json.Marshal(user)
	if err2 != nil {
		log.Fatalf("序列化失败：%v\n", err2)
	}
	result, err := contract.SubmitTransaction("CreateUser", string(userJsonByte))
	if err != nil {
		log.Printf("调用AddUser合约失败：%v\n", err)
	}
	return string(result)
}

func GetSubmittingClientIdentity(contract *client.Contract) string {
	res, err := contract.EvaluateTransaction("GetSubmittingClientIdentity")
	if err != nil {
		log.Printf("调用GetSubmittingClientIdentity合约失败：%v\n", err)
	}
	return string(res)
}
