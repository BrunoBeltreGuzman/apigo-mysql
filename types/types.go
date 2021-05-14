package types

import (
	"github.com/dgrijalva/jwt-go"
)

type Users struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Roles struct {
	Id         int    `json:"id"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

//data
type Data struct {
	User Users `json:"user"`
	Role Roles `json:"role"`
	jwt.StandardClaims
}
