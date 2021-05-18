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

func GetAllSessions(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid User Id")
		return
	}

	connection := database.GetConnection()
	const sql string = "SELECT * FROM sessions WHERE id=?"
	results, err := connection.Query(sql, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err)
		return
	}

	var session = types.Sessions{}
	var sessions = []types.Sessions{}
	for results.Next() {
		err := results.Scan(&session.Id, &session.User, &session.Token, &session.Ip, &session.CreatedAt, &session.ExpiredAt, &session.Created_timeAt, &session.Updated_timeAt)
		if err != nil {
			log.Fatal(err)
		}
		sessions = append(sessions, session)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sessions)
}
func GetByIdSessions(response http.ResponseWriter, request *http.Request) {
	connection := database.GetConnection()
	const sql string = "SELECT * FROM sessions"
	results, err := connection.Query(sql)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err)
		return
	}

	var session = types.Sessions{}
	var sessions = []types.Sessions{}
	for results.Next() {
		err := results.Scan(&session.Id, &session.User, &session.Token, &session.Ip, &session.CreatedAt, &session.ExpiredAt, &session.Created_timeAt, &session.Updated_timeAt)
		if err != nil {
			log.Fatal(err)
		}
		sessions = append(sessions, session)
	}

	defer results.Close()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sessions)
}
