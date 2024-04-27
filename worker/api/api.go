package api

import (
	"log"

	dbHC "github.com/s29papi/wag3r-bot/service/dbHealthChecker"
	rDWD "github.com/s29papi/wag3r-bot/service/renderDoNotWindDown"
	"github.com/s29papi/wag3r-bot/worker"
	client "github.com/s29papi/wag3r-bot/worker/client"
)

var (
	bot                 *worker.Worker
	dbHealthChecker     *dbHC.DBHealthChecker
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

func startDBHealthChecker() {
	if dbHealthChecker != nil {
		log.Println("Cant start Database Health Checker. Database Health Checker is currently running")
		return
	}
	dbHealthChecker = dbHC.NewDBHealthChecker()
	log.Println("Starting Database Health Checker")
	go dbHealthChecker.Start()
}

func stopDBHealthChecker() {
	if dbHealthChecker == nil {
		log.Println("Cant stop Database Health Checker. Database Health Checker is not running")
		return
	}
	go dbHealthChecker.Stop()
	dbHealthChecker = nil
	log.Println("Database Health Checker Stopped")
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
