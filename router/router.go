package router

import (
	"github.com/gin-gonic/gin"
	"gomod/server"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", server.Test)

	//搜索目的地的景点
	router.GET("/popular/search", server.JourneyList)
	//
	router.GET("/guideline/search", server.GuideRecommand)
	return router
}
