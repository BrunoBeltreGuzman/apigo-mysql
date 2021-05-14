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

func FindAllRoles(response http.ResponseWriter, request *http.Request) {
	/*
		authToken := request.Header.Get("Authorization")

		boo := token.DecodeToken(authToken)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(boo)

	*/
	connection := database.GetConnection()
	const sql = "SELECT * FROM roles"
	results, error := connection.Query(sql)
	if error != nil {
		log.Fatal(error)
		return
	}

	var role = types.Roles{}
	var roles = []types.Roles{}
	for results.Next() {
		err := results.Scan(&role.Id, &role.Role, &role.Created_at, &role.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		roles = append(roles, role)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(roles)

}

func FindByIdRoles(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid Role Id")
		return
	}

	connection := database.GetConnection()
	const sql = "SELECT * FROM roles WHERE id=?"
	results, error := connection.Query(sql, id)
	if error != nil {
		log.Fatal(error)
		return
	}

	var role = types.Roles{}
	var roles = []types.Roles{}
	for results.Next() {
		err := results.Scan(&role.Id, &role.Role, role.Role, &role.Created_at, &role.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		roles = append(roles, role)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(roles)
}

func InsertRoles(response http.ResponseWriter, request *http.Request) {
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}

	role := request.FormValue("role")

	if role == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	connection := database.GetConnection()
	const sql = "INSERT INTO roles (role) VALUES(?)"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Roles")
		return
	}
	_, error = insert.Exec(role)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Roles")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Roles Insert successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func UpdateRoles(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}
	role := request.FormValue("role")
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
		return
	}

	if role == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	connection := database.GetConnection()
	const sql = "UPDATE roles set role=? WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Roles")
		return
	}
	_, error = insert.Exec(role, id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Roles")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Roles update successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func DeleteRoles(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid Role Id")
		return
	}

	connection := database.GetConnection()
	const sql = "DELETE FROM roles WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Roles")
		return
	}
	_, error = insert.Exec(id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Roles")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Roles delete successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}
