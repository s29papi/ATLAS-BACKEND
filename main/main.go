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

// i would need to get the logs of this service
/**
* we have a single go routine
* bot.Start() --> 3 go routines ---> a channel and an os.Exit()
* 1st go routine is called tick. has a for loop and select statement which blocks
* pending when its channels recieve a value
*
*
*
 */
//  curl http://hub.freefarcasterhub.com:2281/v1/castsByMention?fid=318902

// we have a free hub
// a frontend vercel
//
// curl http://arena.wield.co/v1/castsByMention?fid=318902
