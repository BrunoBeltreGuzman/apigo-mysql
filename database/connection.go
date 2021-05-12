package database

import (
	"database/sql"
	"fmt"
	"log"
)

func GetConnection() (db *sql.DB) {

	const driver string = "mysql"
	const database string = "api-go"
	const user string = "bobo"
	const password string = "123"

	db, err := sql.Open(driver, user+":"+password+"@tcp(127.0.0.1:3306)/"+database)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database `" + database + "` connection successfull")
	}
	return db
}
