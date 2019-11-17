package dao

import (
	"fmt"
	"log"
)

type Journey struct {
	ID          int     `json:"id"`
	TID         int     `json:"tid"`
	Destination string  `json:"destination"` //景点
	Address     string  `json:"address"`     //地址
	Favorite    int     `json:"favorite"`    //受欢迎程度
	Stime       int     `json:"stime"`       //推荐游览时间
	CoordinateX float64 `json:"coordinateX"` //精度
	CoordinateY float64 `json:"coordinateY"` //维度
	City        string  `json:"city"`        //目的地
	Photo       string  `json:"photo"`       //景区图片
}

func (j *Journey) TableName() string {
	return "journey"
}

//Query by name asc
func (j *Journey) GetJourneybyName(name string) (re *[]Journey, err error) {
	//sql := fmt.Sprintf("SELECT * FROM %s WHERE  city = '%s' order by favorite desc", j.TableName(), name)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE  city = '%s' order by favorite asc", j.TableName(), name)
	log.Printf("%v", sql)
	if re, err = querybySql(sql); err != nil {
		log.Printf("query Journey by name error:%v", err)
		return nil, err
	}

	return re, nil
}

//Query by name DESC
func (j *Journey) GetJourneybyNameDesc(name string) (re *[]Journey, err error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE  city = '%s' order by favorite desc", j.TableName(), name)
	//sql := fmt.Sprintf("SELECT * FROM %s WHERE  city = '%s' order by favorite asc", j.TableName(), name)
	log.Printf("%v", sql)
	if re, err = querybySql(sql); err != nil {
		log.Printf("query Journey by name error:%v", err)
		return nil, err
	}
  return re, nil
}

//query by stime
func (j *Journey) GetJourneybyTime(stime int) (re *[]Journey, err error) {
	//var j Journey
	sql := fmt.Sprintf("SELECT * FROM %s WHERE stime = %v", j.TableName(), stime)
	if re, err = querybySql(sql); err != nil {
		log.Printf("query Journey by name error:%v", err)
		return nil, err
	}

	return re, nil
}

//sql query
func querybySql(sql string) (*[]Journey, error) {
	rows, err := dbConn.Query(sql)
	if err != nil {
		log.Printf("con journey error:%v", err)
		return nil, err
	}

	defer rows.Close()

	var re []Journey

	for rows.Next() {
		var jo Journey
		err := rows.Scan(&jo.ID, &jo.TID, &jo.Destination, &jo.Address, &jo.Favorite, &jo.Stime, &jo.CoordinateX, &jo.CoordinateY, &jo.City, &jo.Photo)
		if err != nil {
			log.Printf("query sql error:%v", err)
			continue
		}

		re = append(re, jo)
	}

	return &re, nil
}
