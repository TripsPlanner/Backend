package service

import "gomod/dao"

func GetDetail(tid int) (re *[]dao.Detail, err error) {
	var d dao.Detail
	re, err = d.GetDetailbyTID(tid)
	return
}
