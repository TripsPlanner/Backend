package router

import (
	"gomod/server"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", server.Test)

	//搜索目的地的景点
	router.GET("/popular/search", server.JourneyList)
	router.GET("/popular/detail", server.Detail)

	return router
}
