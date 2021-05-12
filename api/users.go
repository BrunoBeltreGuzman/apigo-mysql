package api

import (
	"encoding/json"
	"gomysql/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllUsers(response http.ResponseWriter, request *http.Request) {
	users := model.GetAllUsers()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func FindByIdUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
	}
	users := model.GetByIdUsers(id)
	if users != nil {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(users)
	}
}

func UpdateUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	results := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Users update succeeded",
		"id":       params["id"],
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(results)

}

func DeleteUsers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	results := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Users delete succeeded",
		"id":       params["id"],
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(results)

}
