package api

import (
	"encoding/json"
	"gomysql/database"
	"gomysql/model"
	"gomysql/types"
	"gomysql/util/token"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func SignIn(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	var name string = request.FormValue("name")
	var password string = request.FormValue("password")

	if name == "" || password == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("Bad Request 400")
		return
	}

	user := model.SignInUsers(name, password)

	if len(user) > 0 {
		role := model.GetRole(user[0].Role)
		if len(role) == 0 {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode("No found find roles")
			return
		}
		options := types.Data{
			User: user[0],
			Role: role[0],
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "NameApp",
			},
		}
		result := map[string]string{
			"code":    "200",
			"message": "The request has succeeded.",
			"token":   token.GetToken(options),
		}
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(result)
		json.NewEncoder(response).Encode(user[0])
		return
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode("Incorrect User or Password")
		return
	}

}

func SignUp(response http.ResponseWriter, request *http.Request) {
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
		"response": "Signup successfully",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
