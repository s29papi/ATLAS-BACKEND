package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/s29papi/wag3r-bot/worker"
)

type HttpError struct {
	ErrorString string
	ErrorCode   int
}

func AddBotHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	AddBot(2)

	return []byte{0x01, 0x02, 0x03}, nil
}

var SUCCESS_MESSAGE = []byte("success")

func StartBotRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	startBot()
	return SUCCESS_MESSAGE, nil
}

// puts a deposit in the pending array
func DepositRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &HttpError{ErrorString: "Error reading request body", ErrorCode: http.StatusInternalServerError}
	}
	var requestData worker.DepositRequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return nil, &HttpError{ErrorString: "Error decoding JSON", ErrorCode: http.StatusBadRequest}
	}

	depositRequest(requestData)

	return SUCCESS_MESSAGE, nil
}
