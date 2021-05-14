package database

import (
	"database/sql"
	"fmt"
	"gomysql/config"
	"log"
)

func GetConnection() (db *sql.DB) {

	const conn string = config.USER + ":" + config.PASSWORD + config.HOST + "/" + config.DATABASE

	db, err := sql.Open(config.DRIVER, conn)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database `" + config.DATABASE + "` connection successfully")
	}
	return db
}
