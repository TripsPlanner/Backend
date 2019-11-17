package server

import (
	"github.com/gin-gonic/gin"
	"gomod/model"
)

func JsonResponse(c *gin.Context, code int, message string, data ...interface{}) {
	var re model.Response
	if len(data) > 0 {
		re = model.Response{
			Code:    code,
			Message: message,
			Data:    data[0],
		}
	} else {
		re = model.Response{
			Code:    code,
			Message: message,
		}
	}
	c.JSON(200, re)
	c.Abort()
}
