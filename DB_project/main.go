package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "web:#Girgis#2004@/snippetbox?parseTime=true")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}

	defer db.Close()
}
