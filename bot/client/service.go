package client

import (
	"bytes"
	"log"
	"net"
	"net/http"
	"time"
)

type HTTPService struct {
	C *http.Client

	Resp chan []byte
}

func NewHTTPService() *HTTPService {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			// re - using a single tcp connection
			MaxIdleConnsPerHost: 1,
		},
	}

	return &HTTPService{
		C:    c,
		Resp: make(chan []byte, 1),
	}
}

func (h *HTTPService) SendRequest(method string, request *http.Request) {
	switch method {
	case http.MethodGet:
		if !(len(request.Header.Values("accept")) > 0) {
			log.Println("Error: Accept Header not set")
			return
		}
		if !(len(request.Header.Values("api_key")) > 0) {
			log.Println("Error: api_key not set")
			return
		}
	default:
		log.Printf("Error: Passed in HTTP Method %v does not exist", method)
		return
	}
	response, err := h.C.Do(request)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v", err)
		return
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Printf("Error: HTTP Method %v Errored. StatuCode: %v, Status: %v", method, response.StatusCode, response.Status)
		return
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %+v", err)
		return
	}
	h.Resp <- buf.Bytes()
}
