package routers

import (
	"client-go-gateway/middleware"
	"client-go-gateway/routers/api"
	"github.com/gin-gonic/gin"
)

var (
	helloAPi api.HelloApi
)

func InitRouter(contextPath string) *gin.Engine {
	router := gin.New()

	router.Use(middleware.HandleError)
	router.Use(middleware.EnableTraceIdHook)

	currency := router.Group(contextPath + "/hello")
	currency.GET("", helloAPi.Hello)

	return router
}
