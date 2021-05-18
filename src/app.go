package main

import (
	"gomysql/src/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server.App()
}
