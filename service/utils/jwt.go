package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewJWT() string {
	key, err := getKey()
	if err != nil {
		log.Fatal(err)
	}

	return getSignedJwtString(key[:])
}

func getSignedJwtString(key []byte) (signedJwtString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": &jwt.NumericDate{Time: time.Now()},
	})

	signedJwtString, err := token.SignedString(key[:])
	if err != nil {
		log.Fatal(err)
	}

	return
}
