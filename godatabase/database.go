package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:aaaAAA@123@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful")
	defer db.Close()

}
