package routers

import (
	"gomysql/src/api"
	"gomysql/src/middlewares"

	"github.com/gorilla/mux"
)

func GetRouters() (r *mux.Router) {

	routers := mux.NewRouter()

	//middlewares
	routers.Use(middlewares.Log)

	//index
	routers.HandleFunc("/", api.HomeHandler).Methods("GET")
	routers.HandleFunc("/home", api.Home).Methods("GET")

	//Users
	routers.HandleFunc("/users", api.FindAllUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", api.FindByIdUsers).Methods("GET")
	routers.HandleFunc("/users", api.InsertUsers).Methods("POST")
	routers.HandleFunc("/users/{id}", api.UpdateUsers).Methods("PUT")
	routers.HandleFunc("/users/{id}", api.DeleteUsers).Methods("DELETE")

	//Roles
	routers.HandleFunc("/roles", api.FindAllRoles).Methods("GET")
	routers.HandleFunc("/roles/{id}", api.FindByIdRoles).Methods("GET")
	routers.HandleFunc("/roles", api.InsertRoles).Methods("POST")
	routers.HandleFunc("/roles/{id}", api.UpdateRoles).Methods("PUT")
	routers.HandleFunc("/roles/{id}", api.DeleteRoles).Methods("DELETE")

	//Sign
	routers.HandleFunc("/signin", api.SignIn).Methods("POST")
	routers.HandleFunc("/signup", api.SignUp).Methods("POST")

	return routers
}
