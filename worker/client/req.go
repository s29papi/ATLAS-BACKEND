package client

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/s29papi/wag3r-bot/worker/env"
)

func ChannelCastRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.CHANNEL_CAST_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}

func UserMentionsRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.USER_MENTIONS_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}
func NotificationsInStadiumRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.USER_MENTIONS_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}
func DataForCastHash(castHash string) *http.Request {
	url := fmt.Sprintf(env.DataForCastHash_URL, castHash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}

func NotificationsForToshiPayBot() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.NotificationsForToshiPayBot_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}

func MentionReplyRequest(s *strings.Reader) *http.Request {
	req, err := http.NewRequest(http.MethodPost, env.MENTIONS_REPLY_URL, s)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY_TOSHI)
	req.Header.Add("content-type", "application/json")
	return req
}
