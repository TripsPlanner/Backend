package router

import (
	"github.com/gin-gonic/gin"
	"gomod/server"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", server.Test)
	return router
}
