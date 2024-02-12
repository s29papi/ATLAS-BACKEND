package client

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type HTTPService struct {
	C *http.Client

	Resp []byte
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
		C: c,
	}
}

func (h *HTTPService) SendRequest(method string, request *http.Request) []byte {
	if !(len(request.Header.Values("accept")) > 0) {
		log.Println("Error: Accept Header not set")
		return nil
	}
	if !(len(request.Header.Values("api_key")) > 0) {
		log.Println("Error: api_key not set")
		return nil
	}

	if method == http.MethodPost {
		if !(len(request.Header.Values("content-type")) > 0) {
			log.Println("Error: content-type not set")
			return nil
		}
	}

	response, err := h.C.Do(request)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v", err)
		return nil
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Printf("Error: HTTP Method %v Errored. StatuCode: %v, Status: %v", method, response.StatusCode, response.Status)
		return nil
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %+v", err)
		return nil
	}
	fmt.Println(string(buf.Bytes()))

	return buf.Bytes()
}
