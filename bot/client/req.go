package client

import (
	"log"
	"net/http"

	"github.com/s29papi/wag3r-bot/bot/env"
)

func ChannelCastRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.CHANNEL_CAST_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", "NEYNAR_API_DOCS")
	return req
}
