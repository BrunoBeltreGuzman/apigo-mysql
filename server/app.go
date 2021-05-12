package server

import (
	"fmt"
	"gomysql/routers"
	"log"
	"net/http"
	"time"
)

func App() {
	myrouters := routers.GetRouters()

	server := &http.Server{
		Addr:           ":3000",
		Handler:        myrouters,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server starting successfully")
	fmt.Println("http://localhost:3000")
	log.Fatal(server.ListenAndServe())

}
