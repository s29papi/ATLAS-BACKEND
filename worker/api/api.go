package api

import (
	"log"

	rDWD "github.com/s29papi/atlas-backend/service/renderDoNotWindDown"
	"github.com/s29papi/atlas-backend/worker"
	client "github.com/s29papi/atlas-backend/worker/client"
)

var (
	bot                 *worker.Worker
	renderDoNotWindDown *rDWD.RenderDoNotWindDown
)

func startBot() {
	if bot != nil {
		log.Println("Cant start Bot. Bot is currently running")
		return
	}
	bot = worker.NewWorker()
	log.Println("Starting")
	go bot.Start()
}

func stopBot() {
	if bot == nil {
		log.Println("Cant stop Bot. Bot is not running")
		return
	}
	bot.Stop()
	bot = nil
	log.Println("Stopped")
}

func startRenderDoNotWindDown() {
	if renderDoNotWindDown != nil {
		log.Println("Cant start renderDoNotWindDown. renderDoNotWindDown is currently running")
		return
	}
	s := client.NewHTTPService()
	renderDoNotWindDown = rDWD.NewRenderDoNotWindDown(false, s.C)
	log.Println("Starting Render Do Not Wind Down")
	go renderDoNotWindDown.Start()
}
func stopRenderDoNotWindDown() {
	if renderDoNotWindDown == nil {
		log.Println("Cant stop renderDoNotWindDown. renderDoNotWindDown is not currently running")
		return
	}
	go renderDoNotWindDown.Stop()
	renderDoNotWindDown = nil
	log.Println("Render Do Not Wind Down Stopped")
}
