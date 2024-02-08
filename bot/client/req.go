package client

import (
	"log"
	"net/http"
	"os"

	"github.com/s29papi/wag3r-bot/bot/env"
)

func ChannelCastRequest() *http.Request {
	url_env_key := os.Getenv("CHANNEL_CAST_URL")
	if len(url_env_key) == 0 {
		url_env_key = env.CHANNEL_CAST_URL
	}
	req, err := http.NewRequest(http.MethodGet, url_env_key, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", "NEYNAR_API_DOCS")
	return req
}
