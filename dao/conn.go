//mysql struct
//

package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "*******"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "blog"
)

var (
	err    error
	dbConn *sql.DB
)

func Init() {
	//dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%v)/%s?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		//panic("Open mysql failed,err:%v\n", err)
		return
	}

	return
}
