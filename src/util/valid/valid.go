package valid

import (
	"gomysql/src/util/token"
)

func ValidToken(authToken string) bool {
	if authToken == "" {
		return false
	}
	data := token.DecodeToken(authToken)
	if data == nil {
		return false
	}
	return true
}

func IsAdmin(authToken string) bool {
	data := token.DecodeToken(authToken)
	if data == nil {
		return false
	}
	if data.Role.Role == "admin" {
		return true
	} else {
		return false
	}
}

func IsReader(authToken string) bool {
	data := token.DecodeToken(authToken)
	if data == nil {
		return false
	}
	if data.Role.Role == "reader" {
		return true
	} else {
		return false
	}
}

func IsUser(authToken string) bool {
	data := token.DecodeToken(authToken)
	if data == nil {
		return false
	}
	if data.Role.Role == "user" {
		return true
	} else {
		return false
	}
}
