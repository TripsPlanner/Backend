package server

import (
	"github.com/gin-gonic/gin"
	"gomod/model"
)

func JsonResponse(c *gin.Context, code int, message string, data ...interface{}) {
	re := model.Response{
		Code:    code,
		Message: message,
		Data:    data[0],
	}

	c.JSON(200, re)
	c.Abort()
}
