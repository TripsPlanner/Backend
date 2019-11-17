package server

import (
	"github.com/gin-gonic/gin"
	"gomod/model"
	"gomod/service"
	"log"
)

func GuideRecommand(c *gin.Context) {

	gu := model.GuideResq{
		Target:     c.Query("target"),
		StartTime:  c.Query("start_time"),
		StartStage: c.Query("start_stage"),
		EndTime:    c.Query("end_time"),
		EndStage:   c.Query("end_stage"),
	}

	//make new tourist line for custorize
	span := service.CalcSapn(gu.StartTime, gu.StartStage, gu.EndTime, gu.EndStage)
	re, err := service.RecommandLine(gu.Target, span)
	if err != nil {
		JsonResponse(c, 10002, "获取推荐路线失败")
		return
	}

	JsonResponse(c, 0, "success", re)

}

func JourneyList(c *gin.Context) {

	target := c.Query("target")
	log.Printf("搜索目的地：%v的景点", target)
	if target == "" {
		JsonResponse(c, -400, "input parameter error")
		return
	}
	//get journey list
	re, err := service.GetJourneyListbyHot(target)
	if err != nil {
		JsonResponse(c, 10001, "获取景点列表失败")
		return
	}

	JsonResponse(c, 0, "success", re)
}
