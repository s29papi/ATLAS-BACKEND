package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/s29papi/atlas-backend/worker/types"
)

func filterOutNftFramesCast(c types.Casts) types.Casts {
	var casts types.Casts
	excludedDomains := getExcludedDomains()
	for _, cast := range c.Data {
		if len(cast.Frames) == 0 {
			continue
		}
		frameUrl, err := url.Parse(cast.Frames[0].FramesUrl)
		if err != nil {
			log.Println(err)
			return types.Casts{}
		}
		splitedUrl := strings.Split(frameUrl.Hostname(), ".")
		if len(splitedUrl) == 2 {
			if !(isExcludedDomain(splitedUrl[0], excludedDomains)) {
				casts.Data = append(casts.Data, cast)
			}
		}
		if len(splitedUrl) == 3 {
			if !(isExcludedDomain(splitedUrl[1], excludedDomains)) {
				casts.Data = append(casts.Data, cast)
			}
		}
	}
	return casts
}

func filterOutNonPowerBadge(c types.Casts) types.Casts {
	var casts types.Casts
	for _, cast := range c.Data {
		if cast.Author.PowerBadge == true {
			casts.Data = append(casts.Data, cast)
		}
	}
	return casts
}

func filterOutRecentFramesCast(c types.Casts, fc *firestore.Client) types.Casts {
	currReqInfo := fetchCurrentReqInfo(fc)
	var casts types.Casts
	for _, cast := range c.Data {
		if timestamp2secs(cast.Timestamp) < currReqInfo.TimeStamp {
			continue
		}
		if timestamp2secs(cast.Timestamp) == currReqInfo.TimeStamp {
			if strings.EqualFold(cast.Hash, currReqInfo.Hash) {
				continue
			}
		}
		casts.Data = append(casts.Data, cast)
	}
	return casts
}

func fetchCurrentReqInfo(fc *firestore.Client) types.CurrentReqInfo {
	dsnap, err := fc.Collection("trending-frames").Doc("current_req_info").Get(context.Background())
	if err != nil {
		log.Println(err)
		return types.CurrentReqInfo{}
	}
	data := dsnap.Data()
	return types.CurrentReqInfo{
		TimeStamp: data["current_timestamp"].(int64),
		DataId:    data["current_data_id"].(int64),
		Hash:      data["current_cast_hash"].(string),
	}
}

func updateCurrentReqInfo(fc *firestore.Client, t int64, hash, dataId string) {

	criRef := fc.Collection("trending-frames").Doc("current_req_info")
	_, err := criRef.Update(context.Background(), []firestore.Update{
		{Path: "current_timestamp", Value: t},
		{Path: "current_cast_hash", Value: hash},
		{Path: "current_data_id", Value: dataIdToInt(dataId)},
	})
	if err != nil {
		log.Println(err)
	}
}

func resetCurrentReqDataId(fc *firestore.Client) {
	criRef := fc.Collection("trending-frames").Doc("current_req_info")
	_, err := criRef.Update(context.Background(), []firestore.Update{
		{Path: "current_data_id", Value: int64(0)},
	})
	if err != nil {
		log.Println(err)
	}
}

func updateRecentFramesCast(c types.Casts, fc *firestore.Client) string {
	batch := fc.Batch()
	tfRef := fc.Collection("trending-frames").Doc("trending-frames")
	week := getCurrentWeek()
	dataId := getCurrentDataId(fc)

	for _, cast := range c.Data {
		dataId = getTranscientDataId(dataId)
		frameData := types.FrameData{
			ImageUrl:       cast.Frames[0].Image,
			FramesTitle:    cast.Frames[0].Title,
			AuthorUserName: cast.Author.Username,
			Text:           cast.Text,
			Week:           week,
			DataId:         dataId,
			Timestamp:      toUnixInt64(cast.Timestamp),
			AuthorFid:      cast.Author.Fid,
		}

		batch.Set(tfRef, map[string]interface{}{
			week: map[string]interface{}{
				dataId: frameData,
			},
		}, firestore.MergeAll)

	}

	_, err := batch.Commit(context.Background())
	if err != nil {
		log.Println(err)
		return ""
	}
	return dataId
}

func toUnixInt64(s string) int64 {
	timestamp, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	fmt.Println(timestamp.Unix())
	return timestamp.Unix()
}

func getCurrentWeek() string {
	var strBuilder strings.Builder
	_, weekInt := time.Now().ISOWeek()
	strBuilder.WriteString("week")
	strBuilder.WriteString("_")
	strBuilder.WriteString(strconv.Itoa(weekInt))
	return strBuilder.String()
}

func getTranscientDataId(d string) string {
	prefix := "dataId_"
	numStr, _ := strings.CutPrefix(d, prefix)
	num, _ := strconv.Atoi(numStr)
	num += 1
	numStr = strconv.Itoa(int(num))
	return prefix + numStr
}

func dataIdToInt(d string) int64 {
	prefix := "dataId_"
	numStr, _ := strings.CutPrefix(d, prefix)
	num, _ := strconv.Atoi(numStr)
	return int64(num)
}

func getCurrentDataId(fc *firestore.Client) string {
	tfDoc := fc.Collection("trending-frames").Doc("trending-frames")
	dsnap, err := tfDoc.Get(context.Background())
	if err != nil {
		log.Println(err)
		return ""
	}

	data := dsnap.Data()
	currReqInfo := fetchCurrentReqInfo(fc)
	weekStr := getCurrentWeek()

	if data[weekStr] == nil {
		doc := make(map[string]interface{})
		doc[weekStr] = map[string]interface{}{
			"dataId_0": struct{}{},
		}
		_, err = tfDoc.Set(context.Background(), doc)
		if err != nil {
			log.Println(err)
			return ""
		}
		resetCurrentReqDataId(fc)

		return "dataId_0"
	}

	prefix := "dataId_"

	numStr := strconv.Itoa(int(currReqInfo.DataId))

	return prefix + numStr
}

func timestamp2secs(t string) int64 {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	return parsedTime.Unix()
}

func sortHighestTimeAndHash(c []types.CastData) (int64, string) {
	var highestIdx int
	for i, castData := range c {
		if timestamp2secs(c[highestIdx].Timestamp) < timestamp2secs(castData.Timestamp) {
			highestIdx = i
		}
	}
	return timestamp2secs(c[highestIdx].Timestamp), c[highestIdx].Hash
}

func getExcludedDomains() []string {
	var excludedDomains types.DomainData
	nftdomainBytes, err := os.ReadFile("nftdomain.json")
	if err != nil {
		log.Println(err)
		return []string{}
	}
	if err := json.Unmarshal(nftdomainBytes, &excludedDomains); err != nil {
		log.Println(err)
		return []string{}
	}
	return excludedDomains.NFTDomain
}

func isExcludedDomain(domain string, excludedDomains []string) bool {
	for _, excludedDomain := range excludedDomains {
		if !(strings.EqualFold(domain, excludedDomain)) {
			continue
		}
		return true
	}
	return false
}
