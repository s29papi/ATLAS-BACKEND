package service

import (
	"net/http"
	"reflect"

	wApi "github.com/s29papi/wag3r-bot/worker/api"
)

var Mux http.Handler

func muxHandlerFunc(h interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := reflect.ValueOf(h).FieldByIndex([]int{0}).Call([]reflect.Value{reflect.ValueOf(r)})
		if len(result) > 0 {
			if httpError, ok := result[1].Interface().(wApi.HttpError); ok {
				http.Error(w, httpError.ErrorString, httpError.ErrorCode)
			}

			if bytesResult, ok := result[0].Interface().([]byte); ok {
				w.Header().Set("Access-Control-Allow-Methods", "POST")

				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(http.StatusOK)
				w.Write(bytesResult)
			}

		}
	}
}

func init() {
	mux := http.NewServeMux()
	var patternHandlers = make([]map[string]interface{}, 0, 10)
	workerPatternHandler := wApi.Register()
	patternHandlers = append(patternHandlers, workerPatternHandler)
	for _, patternHandler := range patternHandlers {
		for pattern, handler := range patternHandler {
			handlerFunc := muxHandlerFunc(handler)
			mux.HandleFunc(pattern, handlerFunc)
		}
	}
	Mux = mux
}
