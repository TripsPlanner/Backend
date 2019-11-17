package service

import (
	"log"
	"time"
)

func CalcSapn(startTime string, startStage string, endTime string, endStage string) int {
	log.Printf("starttime : %v", startTime)
	log.Printf("endtime : %v", endTime)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	start, _ := time.ParseInLocation("2006-01-02", startTime, loc)
	end, _ := time.ParseInLocation("2006-01-02", endTime, loc)
	log.Printf("start time is :%v", start.String())
	log.Printf("end time is :%v", end.String())
	if startStage == endStage {
		//span := end.Sub(start).Hours()
		return (int(end.Sub(start).Hours())/24)*2 + 1
	}

	if startStage == "am" {
		log.Printf("am is :%v", end.Sub(start).Hours())
		return ((int(end.Sub(start).Hours()) / 24) + 1) * 2
	}

	if startStage == "pm" {
		return (int(end.Sub(start).Hours()) / 24) * 2
	}

	return 0
}
