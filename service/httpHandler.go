package service

import (
	"fmt"
	"net/http"
	"reflect"

	wApi "github.com/s29papi/wag3r-bot/worker/api"
)

var Mux http.Handler

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

func muxHandlerFunc(h interface{}) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		result := reflect.ValueOf(h).FieldByIndex([]int{0}).Call([]reflect.Value{reflect.ValueOf(r)})

		if len(result) > 0 {
			fmt.Println(38774646)
			if httpError, ok := result[1].Interface().(wApi.HttpError); ok {

				http.Error(w, httpError.ErrorString, httpError.ErrorCode)
			}

			if bytesResult, ok := result[0].Interface().([]byte); ok {
				// retval := fmt.Sprintf("%v", `{ "ooo": 1 }`)
				// w.Write([]byte(retval))
				// w.Header().Set("Content-Type", "application/octet-stream")
				// w.Header().Set("accept", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3001")
				w.Header().Set("Access-Control-Allow-Methods", "POST")

				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(http.StatusOK)
				w.Write(bytesResult)

				// fmt.Fprintf(w, "Value of 'ooo': %v", bytesResult)
			}

		}
	}
}
