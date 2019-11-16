package service

import (
	"gomod/dao"
)

//get journey list by name
func GetJourneyList(name string) (re *[]dao.Journey, err error) {
	var jo dao.Journey
	re, err = jo.GetJourneybyName(name)
	return
}
