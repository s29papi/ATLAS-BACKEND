package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/s29papi/wag3r-bot/worker"
)

func main() {
	startBot()
	startServer()
}

var id = 318902

func startBot() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	bot := worker.NewWorker(signalChan)
	go bot.Start()
}

func startServer() {
	http.ListenAndServe(":8090", nil)
}
