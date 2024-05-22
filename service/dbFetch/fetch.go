package dbfetch

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/s29papi/atlas-backend/worker/env"
	"google.golang.org/api/option"
)

type DBFetch struct {
	fc *firestore.Client
}

func NewDBFetch() *DBFetch {
	ctx := context.Background()
	fireStoreCred, err := hex.DecodeString(env.SERVICE_ACCOUNT_JSON_HEX)
	if err != nil {
		log.Fatalln("Error decoding hexadecimal representation:", err)
	}
	opt := option.WithCredentialsJSON(fireStoreCred)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	fc, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &DBFetch{
		fc: fc,
	}
}

func (d *DBFetch) GetTrendingFrames() ([]byte, error) {
	fd := fetchSortedFrames(d.fc)
	return json.Marshal(fd)
}

func (d *DBFetch) GetRecommendedFrames() {

}

func (d *DBFetch) FetchSortedFrames() []FrameData {
	return fetchSortedFrames(d.fc)
}

func fetchSortedFrames(fc *firestore.Client) []FrameData {
	tfDoc := fc.Collection("trending-frames").Doc("trending-frames")
	dsnap, err := tfDoc.Get(context.Background())
	if err != nil {
		log.Println(err)
	}
	data := dsnap.Data()
	weekStr := getCurrentWeek()

	if data[weekStr] == nil {
		return []FrameData{}
	}
	var iDataIds = data[weekStr].(map[string]interface{})
	var framesData []FrameData
	for _, dataMap := range iDataIds {
		if data, ok := dataMap.(map[string]interface{}); ok {
			var frameData FrameData
			var dataCount int64

			if data["image_url"] != nil {
				dataCount += 1
				frameData.ImageUrl = data["image_url"].(string)
			}

			if data["frames_title"] != nil {
				dataCount += 1
				frameData.FramesTitle = data["frames_title"].(string)
			}

			if data["author_username"] != nil {
				dataCount += 1
				frameData.AuthorUserName = data["author_username"].(string)
			}

			if data["text"] != nil {
				dataCount += 1
				frameData.Text = data["text"].(string)
			}
			if data["author_fid"] != nil {
				dataCount += 1
				frameData.AuthorFid = data["author_fid"].(int64)
			}

			if data["week"] != nil {
				dataCount += 1
				frameData.Week = data["week"].(string)
			}

			if data["dataid"] != nil {
				dataCount += 1
				frameData.DataId = data["dataid"].(string)
			}

			if data["timestamp"] != nil {
				dataCount += 1
				frameData.Timestamp = data["timestamp"].(int64)
			}

			if dataCount > 0 {
				framesData = append(framesData, frameData)
			}
		}
	}

	sort.Slice(framesData, func(i, j int) bool {
		if framesData[i].Timestamp == framesData[j].Timestamp {
			return i < j
		}
		return framesData[i].Timestamp > framesData[j].Timestamp
	})

	return framesData
}

func getCurrentWeek() string {
	var strBuilder strings.Builder
	_, weekInt := time.Now().ISOWeek()
	strBuilder.WriteString("week")
	strBuilder.WriteString("_")
	strBuilder.WriteString(strconv.Itoa(weekInt))
	return strBuilder.String()
}

type FrameData struct {
	ImageUrl       string `json:"image_url"`
	FramesTitle    string `json:"frames_title"`
	AuthorUserName string `json:"author_username"`
	AuthorFid      int64  `json:"author_fid"`
	Text           string `json:"text"`
	Week           string `json:"week"`
	DataId         string `json:"dataId"`
	Timestamp      int64  `json:"timestamp"`
}
