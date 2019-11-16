//web handler
//
package server

import (
	_ "gomod/dao"
	"gomod/model" //"gomod/dir"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	re := model.Response{
		Code:    0,
		Message: "success",
		Data:    1000000,
	}

	c.JSON(200, re)
}
