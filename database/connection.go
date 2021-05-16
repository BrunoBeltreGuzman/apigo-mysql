package database

import (
	"database/sql"
	"fmt"
	"gomysql/config"
	"log"
)

var Connection *sql.DB

func GetConnection() (db *sql.DB) {

	if Connection != nil {
		return Connection
	}

	const conn string = config.USER + ":" + config.PASSWORD + config.HOST + "/" + config.DATABASE

	var err error
	Connection, err = sql.Open(config.DRIVER, conn)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database `" + config.DATABASE + "` connection successfully")
	}
	return Connection
}
