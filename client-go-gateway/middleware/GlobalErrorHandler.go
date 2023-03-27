package middleware

import (
	"client-go-gateway/setting"
	"client-go-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var respUtil utils.RespUtil

func HandleError(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			setting.MyLogger.Info(r)

			respUtil.ErrorResp(http.StatusForbidden, "内部错误", ctx)

			ctx.Abort()
		}
	}()
	ctx.Next()
}
