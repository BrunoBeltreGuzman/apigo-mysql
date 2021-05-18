package api

import (
	"encoding/json"
	"gomysql/src/database"
	"net/http"
)

func TestDataBase(response http.ResponseWriter, request *http.Request) {

	connection := database.GetConnection()

	if connection != nil {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode("Database connection has successfully")
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Database connection failed")
	}
}
