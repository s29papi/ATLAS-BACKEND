package api

import (
	"log"

	"github.com/s29papi/wag3r-bot/worker"
)

var bot *worker.Worker

func StartBot() {
	if bot != nil {
		log.Println("Cant start Bot. Bot is currently running")
		return
	}
	bot = worker.NewWorker()
	log.Println("Starting")
	go bot.Start()
}

func StopBot() {
	if bot == nil {
		log.Println("Cant stop Bot. Bot is not running")
		return
	}
	bot.Stop()
	bot = nil
	log.Println("Stopped")
}

type CallBack struct {
	Fn     func()
	ArgsNo int
}

func Register() map[string]interface{} {
	patternFuncs := make(map[string]interface{})
	patternFuncs["/api/worker/start-bot"] = CallBack{
		Fn:     StartBot,
		ArgsNo: 0,
	}
	patternFuncs["/api/worker/stop-bot"] = CallBack{
		Fn:     StopBot,
		ArgsNo: 0,
	}
	return patternFuncs
}
