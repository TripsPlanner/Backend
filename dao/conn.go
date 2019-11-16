//mysql struct
//

package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	USERNAME = "root"
// 	PASSWORD = "m?CpS347#H}^"
// 	NETWORK  = "tcp"
// 	SERVER   = "49.235.244.252"
// 	PORT     = 23306
// 	DATABASE = "tripsPlanner"
// )

var (
	USERNAME = os.Getenv("DB_USER")
	PASSWORD = os.Getenv("DB_PASSWD")
	NETWORK  = "tcp"
	SERVER   = os.Getenv("DB_SERVER")
	PORT     = os.Getenv("DB_PORT")
	DATABASE = "tripsPlanner"
)

var (
	err    error
	dbConn *sql.DB
)

func init() {
	//dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%v)/%s?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	log.Printf("%v", dsn)
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		//panic("Open mysql failed,err:%v\n", err)
		return
	}

	log.Printf("dbcon  is %v", dbConn)
	return
}
