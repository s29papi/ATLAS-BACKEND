package main

import (
	"log"
	"net/http"

	"github.com/s29papi/wag3r-bot/service"
	"github.com/s29papi/wag3r-bot/service/utils"
)

// memory_db
// api
func main() {
	// startBot()
	startServer()
}

var id = 318902

// func startBot() {
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
// 	bot := worker.NewWorker(signalChan)
// 	go bot.Start()
// }

func startServer() {
	a := service.AuthHandler{
		KeyFunc:     utils.KeyFunc,
		HttpHandler: service.Mux,
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: a,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// https://dashboard.render.com/web/srv-cn9tbvuv3ddc73d88a20/settings
