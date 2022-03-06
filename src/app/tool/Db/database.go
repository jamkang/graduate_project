package Db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB = nil

//连接数据库
func DbPointer() *sql.DB {
	if Db == nil {
		database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mygraduate")
		if err != nil {
			fmt.Println("数据库科连接失败")
			panic("Db连接失败")
		}
		Db = database
	}
	return Db
}
