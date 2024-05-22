package task

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/s29papi/atlas-backend/worker/client"
	"github.com/s29papi/atlas-backend/worker/types"
)

func fetchCastWithFrames(service *client.HTTPService) types.Casts {
	req := client.FramesInCast()
	resp := service.SendRequest(http.MethodGet, req)
	var casts types.Casts
	if err := json.Unmarshal(resp, &casts); err != nil {
		log.Println(err)
		return types.Casts{}
	}
	return casts
}
