package dao

import (
	"log"
)

type Detail struct {
	ID        int    `json:id`
	TID       int    `json:tid`
	Transport string `json:transport`
	Stime     int    `json:stime`
	Profile   string `json:profile`
}

func (d *Detail) GetDetailbyTID(TID int) (re *[]Detail, err error) {

	stmt, err := dbConn.Prepare(`SELECT * FROM detail WHERE tid=?`)
	if err != nil {
		log.Printf("con prepare error:%v", err)
		return nil, err
	}
	rows, err := stmt.Query(TID)
	defer stmt.Close()

	if err != nil {
		log.Printf("con detail error:%v", err)
		return nil, err
	}
	defer rows.Close()
	var tmp []Detail

	for rows.Next() {
		var d Detail
		err := rows.Scan(&d.Profile, &d.TID, &d.Stime, &d.Transport, &d.ID)
		if err != nil {
			log.Printf("query sql error:%v", err)
			continue
		}

		tmp = append(tmp, d)
	}
	return &tmp, nil
}
