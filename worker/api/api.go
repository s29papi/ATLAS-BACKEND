package api

import (
	"log"
)

func StartBot() {
	log.Println("Start")
}
func StopBot() {
	log.Println("Stop")
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
