package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/s29papi/wag3r-bot/service"
	"github.com/s29papi/wag3r-bot/service/utils"
)

func main() {
	dev := os.Args[len(os.Args)-1]
	if dev == "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	startServer()
}

var id = 502736

func startServer() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	a := service.AuthHandler{
		KeyFunc:     utils.KeyFunc,
		HttpHandler: service.Mux,
	}
	corsHandler := cors.Handler(a)
	server := http.Server{
		Addr:    ":8181",
		Handler: corsHandler,
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
