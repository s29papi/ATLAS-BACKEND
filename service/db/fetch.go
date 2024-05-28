package db

import (
	"context"
	"encoding/json"
	"log"
	"sort"

	"cloud.google.com/go/firestore"
)

func (d *DB) GetTrendingFrames() ([]byte, error) {
	fd := fetchSortedFrames(d.fc)
	return json.Marshal(fd)
}

func (d *DB) FetchSortedFrames() []FrameData {
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

			if data["frames_url"] != nil {
				dataCount += 1
				frameData.FramesUrl = data["frames_url"].(string)
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

func (d *DB) GetSavedSortedFrames(userfid string) ([]byte, error) {
	return json.Marshal(fetchSavedSortedFrames(d.fc, userfid))
}

func fetchSavedSortedFrames(fc *firestore.Client, userfid string) []FrameData {
	ufbdDoc := fc.Collection("trending-frames").Doc("users-frame-by-dataid")
	dsnap, err := ufbdDoc.Get(context.Background())
	if err != nil {
		log.Println(err)
	}
	data := dsnap.Data()

	if data["users-frame-by-dataid"] == nil {
		return []FrameData{}
	}

	users := data["users-frame-by-dataid"].(map[string]interface{})
	userDataIds := make(map[string]interface{})
	if users[userfid] != nil {
		userDataIds = users[userfid].(map[string]interface{})
	}
	var framesData []FrameData
	for _, dataMap := range userDataIds {
		if data, ok := dataMap.(map[string]interface{}); ok {
			var frameData FrameData
			var dataCount int64

			if data["image_url"] != nil {
				dataCount += 1
				frameData.ImageUrl = data["image_url"].(string)
			}

			if data["frames_url"] != nil {
				dataCount += 1
				frameData.FramesUrl = data["frames_url"].(string)
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
	// v:= []firestore.Update{
	// 	{
	// 		FieldPath: firestore.FieldPath{
	// 			"",
	// 		},
	// 	},
	// }

	sort.Slice(framesData, func(i, j int) bool {
		if framesData[i].Timestamp == framesData[j].Timestamp {
			return i < j
		}
		return framesData[i].Timestamp > framesData[j].Timestamp
	})
	return framesData
}
