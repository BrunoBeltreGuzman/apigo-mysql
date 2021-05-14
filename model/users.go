package model

import (
	"gomysql/database"
	"gomysql/types"
	"log"
)

func FindUsersByName(name string) []types.Users {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users WHERE name=?"
	results, error := connection.Query(sql, name)
	if error != nil {
		log.Fatal(error)
		return []types.Users{}
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	return users
}

func FindUsersByEmail(email string) []types.Users {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users WHERE email=?"
	results, error := connection.Query(sql, email)
	if error != nil {
		log.Fatal(error)
		return []types.Users{}
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	return users
}

func FindUsersByPassword(password string) []types.Users {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users WHERE password=?"
	results, error := connection.Query(sql, password)
	if error != nil {
		log.Fatal(error)
		return []types.Users{}
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	return users
}

func SignInUsers(name string, password string) []types.Users {
	connection := database.GetConnection()
	const sql = "SELECT * FROM users WHERE name=? AND password=?"
	results, error := connection.Query(sql, name, password)
	if error != nil {
		log.Fatal(error)
		return []types.Users{}
	}

	var user = types.Users{}
	var users = []types.Users{}
	for results.Next() {
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer results.Close()
	return users
}
