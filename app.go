package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gomysql/server"
)

func main() {
	server.App()
}
