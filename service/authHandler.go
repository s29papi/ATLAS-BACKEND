package service

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	KeyFunc     func(token *jwt.Token) (interface{}, error)
	HttpHandler http.Handler
}

func (a AuthHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var (
		tokenStr string
		claims   jwt.RegisteredClaims
	)
	if auth := req.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		tokenStr = strings.TrimPrefix(auth, "Bearer ")
	}
	if len(tokenStr) == 0 {
		http.Error(resp, "missing token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(tokenStr, &claims, a.KeyFunc, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		http.Error(resp, err.Error(), http.StatusUnauthorized)
		return
	}

	if token.Valid {
		a.HttpHandler.ServeHTTP(resp, req)
	} else {
		http.Error(resp, "invalid token", http.StatusUnauthorized)
	}
}
