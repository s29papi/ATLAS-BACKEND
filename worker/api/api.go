package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s29papi/wag3r-bot/worker"
)

var bot *worker.Worker

func startBot() {
	fmt.Println(337373763)
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

func StoppedFrameEvents() {
	if bot == nil {
		log.Println("Cant stop frame events. Bot is not running")
		return
	}
	bot.StoppedFrameEvents = true
}

func AddBot(ooo int64) {
	ooo += ooo
	fmt.Println(ooo)
}

func DepositRequest() {

}

type CallBack struct {
	Fn     func()
	ArgsNo int
}

type CallBackAddBot struct {
	Fn     func(r *http.Request) ([]byte, *HttpError)
	ArgsNo int
}

func Register() map[string]interface{} {
	patternFuncs := make(map[string]interface{})
	// patternFuncs["/api/worker/start-bot"] = CallBack{
	// 	Fn:     StartBot,
	// 	ArgsNo: 0,
	// }
	patternFuncs["/api/worker/stop-bot"] = CallBack{
		Fn:     StopBot,
		ArgsNo: 0,
	}
	patternFuncs["/api/worker/stop-frame-events"] = CallBack{
		Fn:     StoppedFrameEvents,
		ArgsNo: 0,
	}
	patternFuncs["/api/worker/add-bot"] = CallBackAddBot{
		Fn:     AddBotHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/deposit-request"] = CallBackAddBot{
		Fn:     DepositRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/start-bot"] = CallBackAddBot{
		Fn:     StartBotRequestHandleFunc,
		ArgsNo: 2,
	}
	return patternFuncs
}

func depositRequest(d worker.DepositRequestData) {
	bot.SendDepositRequest(d)
}
