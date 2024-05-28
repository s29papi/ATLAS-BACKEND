package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/s29papi/atlas-backend/worker/client"
	"github.com/s29papi/atlas-backend/worker/types"
)

type HttpError struct {
	ErrorString string
	ErrorCode   int
}

type CallBack struct {
	Fn func(r *http.Request) ([]byte, *HttpError)
}

func Register() map[string]interface{} {
	patternFuncs := make(map[string]interface{})
	patternFuncs["/api/service/get-trending-frames"] = CallBack{
		Fn: GetTrendingFramesHandleFunc,
	}
	patternFuncs["/api/service/get-recommended-frames"] = CallBack{
		Fn: GetRecommendedFramesHandleFunc,
	}
	patternFuncs["/api/service/save-user-frames"] = CallBack{
		Fn: SaveUserFrameByDataIdHandleFunc,
	}
	patternFuncs["/api/service/rm-user-frames"] = CallBack{
		Fn: RMUserFrameByDataIdHandleFunc,
	}
	patternFuncs["/api/service/saved-user-frames-by-fid"] = CallBack{
		Fn: GetSavedUserFramesByFid,
	}

	return patternFuncs
}

func GetTrendingFramesHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodGet {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	frames, err := getTrendingFrames()
	if err != nil {
		log.Println(err)
		return nil, &HttpError{ErrorString: err.Error(), ErrorCode: 0}
	}
	return frames, nil
}

func GetRecommendedFramesHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	var reqBody types.BulkFollowingRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return nil, &HttpError{ErrorString: "Invalid request body", ErrorCode: http.StatusBadRequest}
	}
	frames, err := getRecommendedFrames(reqBody.ViewerFid, client.NewHTTPService())
	if err != nil {
		log.Println(err)
		return nil, &HttpError{ErrorString: err.Error(), ErrorCode: 0}
	}
	return frames, nil
}

func GetSavedUserFramesByFid(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}

	var reqBody types.SavedUserFrameByFidRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return nil, &HttpError{ErrorString: "Invalid request body", ErrorCode: http.StatusBadRequest}
	}

	frames, err := getSavedUserFramesByFid(reqBody.UserFid)
	if err != nil {
		log.Println(err)
		return nil, &HttpError{ErrorString: err.Error(), ErrorCode: 0}
	}
	return frames, nil
}

func SaveUserFrameByDataIdHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	var reqBody types.SaveUserFrameRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return nil, &HttpError{ErrorString: "Invalid request body", ErrorCode: http.StatusBadRequest}
	}
	isSuccess, err := putUserFrameByDataId(reqBody.DataId, reqBody.UserFid)
	if err != nil {
		log.Println(err)
		return nil, &HttpError{ErrorString: err.Error(), ErrorCode: 0}
	}

	return isSuccess, nil
}
func RMUserFrameByDataIdHandleFunc(r *http.Request) ([]byte, *HttpError) {
	if r.Method != http.MethodPost {
		return nil, &HttpError{ErrorString: "Method not allowed", ErrorCode: http.StatusMethodNotAllowed}
	}
	var reqBody types.SaveUserFrameRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return nil, &HttpError{ErrorString: "Invalid request body", ErrorCode: http.StatusBadRequest}
	}
	isSuccess, err := rmUserFrameByDataId(reqBody.DataId, reqBody.UserFid)
	if err != nil {
		log.Println(err)
		return nil, &HttpError{ErrorString: err.Error(), ErrorCode: 0}
	}

	return isSuccess, nil
}
