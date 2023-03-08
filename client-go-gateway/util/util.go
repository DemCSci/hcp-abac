package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strings"
)

import "github.com/panjf2000/ants/v2"

var Pool ants.Pool
var Rdb *redis.Client
var Ctx = context.Background()

func FormatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, " ", ""); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

func GetUUID() string {
	u := uuid.New()
	uuidString := strings.ReplaceAll(u.String(), "-", "")
	return uuidString
}
