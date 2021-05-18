package api

import (
	"encoding/json"
	"gomysql/src/config"
	"gomysql/src/database"
	"gomysql/src/types"
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

func InsertSession(response http.ResponseWriter, request *http.Request) {
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}

	user := request.FormValue("user")
	token := request.FormValue("token")
	ip := request.FormValue("ip")

	if user == "" && token == "" && ip == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	connection := database.GetConnection()
	const sql = "INSERT INTO sessions (user, token, ip, expiredAt) VALUES (?,?,?,?)"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Session")
		return
	}
	_, error = insert.Exec(
		user,
		token,
		ip,
		"DATE_ADD(now(), INTERVAL"+strconv.Itoa(config.TIMESESSION)+"MINUTE)")

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in Insert Session")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Session Insert successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func UpdateSession(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	error := request.ParseForm()
	if error != nil {
		log.Fatal(error)
		return
	}
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid Sessions Id")
		return
	}

	connection := database.GetConnection()
	const sql = "UPDATE sessions set expiredAt=? WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Sessions")
		return
	}
	_, error = insert.Exec(
		"DATE_ADD(now(), INTERVAL"+strconv.Itoa(config.TIMESESSION)+"MINUTE)",
		id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in update Sessions")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Sessions update successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func DeleteSession(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Invalid Sessions Id")
		return
	}

	connection := database.GetConnection()
	const sql = "DELETE FROM sessions WHERE id=?"
	insert, error := connection.Prepare(sql)
	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Sessions")
		return
	}
	_, error = insert.Exec(id)

	if error != nil {
		log.Fatal(error)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Has a error in delete Sessions")
		return
	}

	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "Sessions delete successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}
