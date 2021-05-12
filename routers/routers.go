package routers

import (
	"github.com/gorilla/mux"
	"gomysql/api"
)

func GetRouters() (r *mux.Router) {

	routers := mux.NewRouter()

	//index
	routers.HandleFunc("/", api.HomeHandler).Methods("GET")

	//Users
	routers.HandleFunc("/users", api.FindAllUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", api.FindByIdUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", api.UpdateUsers).Methods("PUT")
	routers.HandleFunc("/users/{id}", api.DeleteUsers).Methods("DELETE")

	//Querys
	routers.HandleFunc("/querys", api.Select).Methods("GET")
	routers.HandleFunc("/querys", api.Insert).Methods("POST")

	//test
	routers.HandleFunc("/test/db", api.TestDataBase).Methods("GET")

	return routers
}