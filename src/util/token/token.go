package token

import (
	"gomysql/src/config"
	"gomysql/src/types"

	"log"

	"github.com/dgrijalva/jwt-go"
)

func DecodeToken(token string) *types.Data {
	//Valid Token
	validToken, _ := jwt.ParseWithClaims(
		token,
		&types.Data{},
		func(validToken *jwt.Token) (interface{}, error) {
			return []byte(config.KEY), nil
		},
	)

	data, _ := validToken.Claims.(*types.Data)

	return data
}

func GetToken(options types.Data) string {
	objectJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, options)
	token, err := objectJWT.SignedString([]byte(config.KEY))
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	//Token
	return token
}
