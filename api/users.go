package api

import (
	"encoding/json"
	"gomysql/database"
	"gomysql/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllUsers(response http.ResponseWriter, request *http.Request) {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users"
	results, error := connection.Query(sql)
	if error != nil {
		log.Fatal(error)
		return
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func FindByIdUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
		return
	}

	connection := database.GetConnection()
	const sql = "SELECT * FROM users WHERE id=?"
	results, error := connection.Query(sql, id)
	if error != nil {
		log.Fatal(error)
		return
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func InsertUsers(response http.ResponseWriter, request *http.Request) {
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}

	name := request.FormValue("name")
	email := request.FormValue("email")
	password := request.FormValue("password")

	if name == "" && email == "" && password == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	connection := database.GetConnection()
	const sql = "INSERT INTO users (name, email, password) VALUES(?,?,?)"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Users")
		return
	}
	_, error = insert.Exec(name, email, password)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Users")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Users Insert successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func UpdateUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}
	name := request.FormValue("name")
	email := request.FormValue("email")
	password := request.FormValue("password")
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
		return
	}

	if name == "" && email == "" && password == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	connection := database.GetConnection()
	const sql = "UPDATE users set name=?, email=?, password=? WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Users")
		return
	}
	_, error = insert.Exec(name, email, password, id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Users")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Users update successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func DeleteUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
		return
	}

	connection := database.GetConnection()
	const sql = "DELETE FROM users WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Users")
		return
	}
	_, error = insert.Exec(id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Users")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Users delete successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}
