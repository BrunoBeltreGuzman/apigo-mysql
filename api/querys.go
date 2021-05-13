package api

import (
	"encoding/json"
	"gomysql/database"
	"gomysql/types"
	"log"
	"net/http"
)

func Select(response http.ResponseWriter, request *http.Request) {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users"
	results, error := connection.Query(sql)

	if error != nil {
		log.Fatal(error)
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

func Insert(response http.ResponseWriter, request *http.Request) {

	connection := database.GetConnection()
	const sql = "INSERT INTO users (name, email, password) VALUES('sa', 'sa@gmail.com', '123')"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Query")
	} else {
		_, error := insert.Exec()
		if error != nil {
			log.Fatal(error)
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode("Has a error in Insert Query")
		} else {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode("Query Insert successfully")
		}

	}

}

func Update(response http.ResponseWriter, request *http.Request) {

	connection := database.GetConnection()
	const sql = "UPDATE users set name='saUpdate', email='saUpdate@gmail.com', password='123Update' WHERE id='1'"

	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Query")
	} else {
		_, error := insert.Exec()
		if error != nil {
			log.Fatal(error)
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode("Has a error in Insert Query")
		} else {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode("Query Insert successfully")
		}

	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	connection := database.GetConnection()
	const sql = "DELETE FROM users WHERE id='1'"

	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Query")
	} else {
		_, error := insert.Exec()
		if error != nil {
			log.Fatal(error)
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode("Has a error in Insert Query")
		} else {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode("Query Insert successfully")
		}

	}

}
