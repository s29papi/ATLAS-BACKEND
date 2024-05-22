package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/s29papi/atlas-backend/service"
	"github.com/s29papi/atlas-backend/service/utils"
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

func startServer() {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "http://localhost:3001"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler
	a := service.AuthHandler{
		KeyFunc:     utils.KeyFunc,
		HttpHandler: service.Mux,
	}
	server := http.Server{
		Addr:    ":8181",
		Handler: corsHandler(a),
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
