package server

import (
	"github.com/gin-gonic/gin"
	"gomod/service"
	"log"
)

func Guide(c *gin.Context) {
	params := new(struct {
		Target     string `form:"target`
		StartTime  string `form:"start_time"`
		StartStage string `form:"start_stage"`
		EndTime    string `form:"end_time"`
		EndStage   string `form:"end_stage"`
	})

	if err := c.Bind(params); err != nil {
		log.Printf("input parameter error:%v", err)
		JsonResponse(c, -400, "input parameter error")
		return
	}

	//make new tourist line

}

func JourneyList(c *gin.Context) {
	// params := new(struct {
	// 	Target string `form:"target`
	// })

	// if err := c.Bind(params); err != nil {
	// 	log.Printf("input parameter error:%v", err)
	// 	JsonResponse(c, -400, "input parameter error")
	// 	return
	// }

	target := c.Query("target")
	log.Printf("%v", target)
	//get journey list
	re, err := service.GetJourneyList(target)
	if err != nil {
		JsonResponse(c, 10001, "获取景点列表失败")
		return
	}

	JsonResponse(c, 0, "success", re)
}
