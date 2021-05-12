package main

import (
	"gomysql/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server.App()
}
