package controller

import (
	"client-go-gateway/contract"
	"io"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, request *http.Request) {
	users := contract.GetAllUsers(contract.Contract1)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, users)
}

func AddUser(w http.ResponseWriter, request *http.Request) {

	record := contract.AddUser(contract.Contract3)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, record)
}

func GetSubmittingClientIdentity(w http.ResponseWriter, request *http.Request) {
	identity := contract.GetSubmittingClientIdentity(contract.Contract2)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, identity)
}
