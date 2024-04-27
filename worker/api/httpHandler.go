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

func Register() map[string]interface{} {
	patternFuncs := make(map[string]interface{})

	patternFuncs["/api/worker/stop-bot"] = CallBack{
		Fn:     StopBotRequestHandleFunc,
		ArgsNo: 0,
	}
	patternFuncs["/api/worker/start-bot"] = CallBack{
		Fn:     StartBotRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/start-db-healthchecker"] = CallBack{
		Fn:     StartDBHealthCheckerRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/stop-db-healthchecker"] = CallBack{
		Fn:     StopDBHealthCheckerRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/start-db-healthchecker"] = CallBack{
		Fn:     StartDBHealthCheckerRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/stop-db-healthchecker"] = CallBack{
		Fn:     StopDBHealthCheckerRequestHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/start-rdWD"] = CallBack{
		Fn:     StartRenderDoNotWindDownHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/stop-rdWD"] = CallBack{
		Fn:     StopRenderDoNotWindDownHandleFunc,
		ArgsNo: 2,
	}
	patternFuncs["/api/worker/dummy-rdWD"] = CallBack{
		Fn:     dummyRenderDoNotWindDown,
		ArgsNo: 2,
	}
	return patternFuncs
}

var SUCCESS_MESSAGE = []byte("success")

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

func StartDBHealthCheckerRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	startDBHealthChecker()
	return SUCCESS_MESSAGE, nil
}

func StopDBHealthCheckerRequestHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	stopDBHealthChecker()
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
