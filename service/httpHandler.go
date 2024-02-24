package service

import (
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
			handlerFunc := func(h interface{}) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					reflect.ValueOf(h).FieldByIndex([]int{0}).Call([]reflect.Value{})
				}
			}(handler)
			mux.HandleFunc(pattern, handlerFunc)
		}
	}
	Mux = mux
}
