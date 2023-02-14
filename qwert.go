package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS example_db")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("USE example_db")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT NOT NULL AUTO_INCREMENT,
            name VARCHAR(50) NOT NULL,
            email VARCHAR(50) NOT NULL,
            PRIMARY KEY (id)
        )
    `)
	if err != nil {
		fmt.Println(err)
		return
	}
}
