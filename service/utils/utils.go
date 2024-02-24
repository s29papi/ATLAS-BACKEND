package utils

import (
	"log"
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
	data, err := os.ReadFile("jwt.txt")
	if err != nil {
		log.Println(err)
		return [32]byte{}, err
	}
	var bytes32data [32]byte
	copy(bytes32data[:], data[:32])

	return bytes32data, nil
}
