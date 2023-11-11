package utils

import (
	"github.com/go-basic/uuid"
	"log"
	"strings"
)

func GenerateUUID() (string, error) {
	generateUUID, err := uuid.GenerateUUID()
	if err != nil {
		log.Fatal("UUID生成错误")
	}
	generateUUID = strings.Replace(generateUUID, "-", "", -1)
	return generateUUID, err
}
