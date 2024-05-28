package db

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
)

type RMSuccess struct {
	Success bool   `json:"isRMSuccessful"`
	Msg     string `json:"RMMsg"`
}

func (d *DB) RMUserFrameByDataId(userfid, dataid string) ([]byte, error) {
	return json.Marshal(rmUserFrameByDataId(d.fc, userfid, dataid))
}

func rmUserFrameByDataId(fc *firestore.Client, UserFid, DataId string) RMSuccess {
	weekStr := getCurrentWeek()
	_, err := fc.Collection("trending-frames").Doc("users-frame-by-dataid-bool").
		Update(context.Background(), []firestore.Update{
			{
				FieldPath: []string{weekStr, UserFid, DataId},
				Value:     firestore.Delete,
			},
		})
	if err != nil {
		log.Println(err)
		return RMSuccess{Success: false}
	}
	_, err = fc.Collection("trending-frames").Doc("users-frame-by-dataid").
		Update(context.Background(), []firestore.Update{
			{
				FieldPath: []string{"users-frame-by-dataid", UserFid, DataId},
				Value:     firestore.Delete,
			},
		})
	if err != nil {
		log.Println(err)
		return RMSuccess{Success: false}
	}
	return RMSuccess{Success: true}
}
