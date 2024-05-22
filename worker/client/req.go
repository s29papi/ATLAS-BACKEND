package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s29papi/atlas-backend/worker/env"
)

func FramesInCast() *http.Request {
	req, err := http.NewRequest(http.MethodGet, env.CASTWITHFRAMES_URL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}

func BulkFollowing(authorFid, viewerFid string) *http.Request {
	formattedURL := fmt.Sprintf(env.BULKFOLLOWING_URL, authorFid, viewerFid)
	req, err := http.NewRequest(http.MethodGet, formattedURL, nil)
	if err != nil {
		log.Println("Error: couldn't create requests")
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api_key", env.API_KEY)
	return req
}
