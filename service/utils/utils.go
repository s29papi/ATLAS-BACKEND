package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func KeyFunc(token *jwt.Token) (interface{}, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	return key[:], nil
}

func getKey() ([32]byte, error) {
	data := os.Getenv("JWT_KEY")
	var bytes32data [32]byte
	copy(bytes32data[:], data[:32])

	return bytes32data, nil
}
