package controller

import (
	"client-go-gateway/contract"
	"client-go-gateway/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, request *http.Request) {
	users := contract.GetAllUsers(util.ClientInfoMap["org4MSP"])
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, users)
}

func AddUser(w http.ResponseWriter, request *http.Request) {
	var addUserRequest AddUserRequest
	bodyByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	err = json.Unmarshal(bodyByte, &addUserRequest)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	record := contract.AddUser(util.ClientInfoMap[addUserRequest.MspID])

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, record)
}

func AddAllUser(w http.ResponseWriter, request *http.Request) {

	for _, val := range util.ClientInfoMap {
		contract.AddUser(val)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "注册全部用户身份成功")
}

func GetSubmittingClientIdentity(w http.ResponseWriter, request *http.Request) {
	var addUserRequest AddUserRequest
	bodyByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	err = json.Unmarshal(bodyByte, &addUserRequest)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	identity := contract.GetSubmittingClientIdentity(util.ClientInfoMap[addUserRequest.MspID])
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, identity)
}
