package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	dbfetch "github.com/s29papi/atlas-backend/service/dbFetch"
	"github.com/s29papi/atlas-backend/worker/client"
	"github.com/s29papi/atlas-backend/worker/types"
)

func getTrendingFrames() ([]byte, error) {
	return dbfetch.NewDBFetch().GetTrendingFrames()
}

func getRecommendedFrames(viewerFid string, service *client.HTTPService) ([]byte, error) {
	var trendingFramesByFollowing []dbfetch.FrameData
	var authorFids []string
	trendingFrames := dbfetch.NewDBFetch().FetchSortedFrames()

	for _, tFrame := range trendingFrames {
		authorFids = append(authorFids, strconv.Itoa(int(tFrame.AuthorFid)))
	}

	encodedauthorfids := strings.Join(authorFids, ",")
	encodedQueryString := url.QueryEscape(encodedauthorfids)

	resp := service.SendRequest(http.MethodGet, client.BulkFollowing(encodedQueryString, viewerFid))

	var usersData types.UsersData
	if err := json.Unmarshal(resp, &usersData); err != nil {
		return nil, err
	}

	for idx, tFrame := range trendingFrames {
		if usersData.UserData[idx].ViewerCtx.Following {
			trendingFramesByFollowing = append(trendingFramesByFollowing, tFrame)
		}
	}
	if len(trendingFramesByFollowing) == 0 {
		return nil, nil
	}
	return json.Marshal(trendingFramesByFollowing)
}
