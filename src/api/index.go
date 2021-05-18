package api

import (
	"encoding/json"
	"fmt"
	"gomysql/src/util/valid"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, response *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!!")
}

func Home(response http.ResponseWriter, request *http.Request) {
	authToken := request.Header.Get("Authorization")

	boo := valid.ValidToken(authToken)
	if !boo {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	boo = valid.IsAdmin(authToken)
	if !boo {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Required Admin Role")
		return
	}
	result := map[string]string{
		"code":     "200",
		"message":  "The request has succeeded.",
		"response": "You has accepted!!",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
