package api

import (
	"log"
	"net/http"
)

type HttpError struct {
	ErrorString string
	ErrorCode   int
}
type CallBack struct {
	Fn     func(r *http.Request) ([]byte, *HttpError)
	ArgsNo int
}

var SUCCESS_MESSAGE = []byte("success")

func Register() map[string]interface{} {
	patternFuncs := make(map[string]interface{})

	patternFuncs["/api/worker/stop-bot"] = CallBack{
		Fn: StopBotRequestHandleFunc,
	}
	patternFuncs["/api/worker/start-bot"] = CallBack{
		Fn: StartBotRequestHandleFunc,
	}
	patternFuncs["/api/worker/start-rdWD"] = CallBack{
		Fn: StartRenderDoNotWindDownHandleFunc,
	}
	patternFuncs["/api/worker/stop-rdWD"] = CallBack{
		Fn: StopRenderDoNotWindDownHandleFunc,
	}
	patternFuncs["/api/worker/dummy-rdWD"] = CallBack{
		Fn: dummyRenderDoNotWindDown,
	}

	return patternFuncs
}

func StartBotRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	startBot()
	return SUCCESS_MESSAGE, nil
}

func StopBotRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	stopBot()
	return SUCCESS_MESSAGE, nil
}

func dummyRenderDoNotWindDown(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodGet {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	log.Println("DummyRenderDoNotWindDown: Fired 🔥🔥🔥")
	return SUCCESS_MESSAGE, nil
}

func StartRenderDoNotWindDownHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	startRenderDoNotWindDown()
	return SUCCESS_MESSAGE, nil
}
func StopRenderDoNotWindDownHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	stopRenderDoNotWindDown()
	return SUCCESS_MESSAGE, nil
}
