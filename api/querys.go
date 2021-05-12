package api

import (
	"encoding/json"
	"gomysql/database"
	"gomysql/types"
	"log"
	"net/http"
)

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

func Select(response http.ResponseWriter, request *http.Request) {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users"
	results, error := connection.Query(sql)
	defer results.Close()

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
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}
