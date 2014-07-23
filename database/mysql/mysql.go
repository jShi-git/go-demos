package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, e := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dht?charset=utf8")
	if e != nil {
		print("ERROR")
		return
	}
	println("Conn DB OK")
	rows, e := db.Query("select * from peerip")
	if e != nil {
		fmt.Print("error:%v\n", e)
		return
	}
	if rows == nil {
		print("Rows is nil")
		return
	}
	for rows.Next() {
		var u, p string
		err := rows.Scan(&u, &p)
		if err != nil {
			print("Row error!")
			return
		}
		fmt.Println(u, " ", p)
	}
	println("WIN!")
}
