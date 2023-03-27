package api

import (
	"client-go-gateway/utils"
	"github.com/gin-gonic/gin"
)

type HelloApi struct {
	respUtil utils.RespUtil
}

func (c *HelloApi) Hello(ctx *gin.Context) {

	c.respUtil.SuccessResp("hello world", ctx)
}
