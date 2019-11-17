// web handler
// detail

package server

import (
	_ "gomod/dao" //"gomod/dir"
	"gomod/service"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func Detail(c *gin.Context) {
	// params := new(struct {
	// 	Target string `form:"target`
	// })

	tid := c.Query("tid")
	Tid, _ := StrToInt(tid)
	log.Printf("%v", tid)
	//get journey list
	re, err := service.GetDetail(Tid)
	if err != nil {
		JsonResponse(c, 10001, "获取景点详细信息失败")
		return
	}

	JsonResponse(c, 0, "success", re)
}
