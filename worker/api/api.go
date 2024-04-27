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
	renderDoNotWindDown = rDWD.NewRenderDoNotWindDown(true, s.C)
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

// {
// 	"messages":
// 		[
// 			{
// 			  "data":
// 				{
// 					"type":"MESSAGE_TYPE_CAST_ADD",
// 					"fid":399452,
// 					"timestamp":104737157,
// 					"network":"FARCASTER_NETWORK_MAINNET",
// 					"castAddBody":
// 						{
// 							"embedsDeprecated":[],
// 							"mentions":[502736],
// 							"text":" hey",
// 							"mentionsPositions":[0],
// 							"embeds":[]
// 						}
// 				},
// 					"hash":"0x25e621250906f784a4a6eec22c0bd4d898d4564a",
// 					"hashScheme":"HASH_SCHEME_BLAKE3",
// 					"signature":"YyHQCVOdLfdnF+cgkmtuS7zOEWdcAo+448N8zRfBx2r2GP+4og3cXNPPoRFHr2S0A6tdqJS7UKdj0REV/5g7Aw==",
// 					"signatureScheme":"SIGNATURE_SCHEME_ED25519",
// 					"signer":"0x94542a1c465ad74982cd8874cfa758aa38c450edd65d829ff3da067390e79c10"
// 			}
// 		],
// 	"nextPageToken":""
// }
