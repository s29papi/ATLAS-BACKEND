package main

import (
	"os"
	"os/signal"
	"syscall"

	worker "github.com/s29papi/wag3r-bot/bot"
	// _ "github.com/s29papi/wag3r-bot/bot/env"
)

func main() {

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	bot := worker.NewWorker(signalCh)
	bot.Start()
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
//  curl http://arena.wield.co:2281/v1/castsByFid?fid=2
