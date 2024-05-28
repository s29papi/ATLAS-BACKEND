package db

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
)

type PutSuccess struct {
	Success bool   `json:"isPutSuccessful"`
	Msg     string `json:"PutMsg"`
}

func (d *DB) PutUserFrameByDataId(userfid, dataid string) ([]byte, error) {
	return json.Marshal(putUserFrameByDataId(d.fc, userfid, dataid))
}

// 1. check if the user has the dataId, this check is done by week
// 2. if has return false with message, Already exist
// 3. if user doent have put the data
// 4. the update 1. to reflect this

func putUserFrameByDataId(fc *firestore.Client, UserFid, DataId string) PutSuccess {
	ufbbdDoc := fc.Collection("trending-frames").Doc("users-frame-by-dataid-bool")
	ufbbdsnap, err := ufbbdDoc.Get(context.Background())
	if err != nil {
		log.Println(err)
	}
	ufbbdata := ufbbdsnap.Data()
	weekStr := getCurrentWeek()
	users := make(map[string]interface{})
	if ufbbdata[weekStr] != nil {
		users = ufbbdata[weekStr].(map[string]interface{})
	}
	dataIdByFid := make(map[string]interface{})
	if users[UserFid] != nil {
		dataIdByFid = users[UserFid].(map[string]interface{})
	}
	dataIdExist := false
	if dataIdByFid[DataId] != nil {
		dataIdExist = dataIdByFid[DataId].(bool)
	}
	if dataIdExist {
		return PutSuccess{Success: false, Msg: "Already, Exists !!!"}
	}

	tfDoc := fc.Collection("trending-frames").Doc("trending-frames")
	tfdsnap, err := tfDoc.Get(context.Background())
	if err != nil {
		log.Println(err)
	}
	tfdata := tfdsnap.Data()
	if tfdata[weekStr] == nil {
		return PutSuccess{Success: false, Msg: "Trending Frames Doesn't Exist !!!"}
	}
	tfdataIds := tfdata[weekStr].(map[string]interface{})

	tfdataVal := make(map[string]interface{})
	if tfdataIds[DataId] != nil {
		tfdataVal = tfdataIds[DataId].(map[string]interface{})
	}

	ufbdDoc := fc.Collection("trending-frames").Doc("users-frame-by-dataid")
	// ufbdsnap, err := ufbdDoc.Get(context.Background())
	// if err != nil {
	// 	log.Println(err)
	// }
	// ufbddata := ufbdsnap.Data()
	// users = make(map[string]interface{})
	// if ufbddata["users-frame-by-dataid"] != nil {
	// 	users = ufbddata["users-frame-by-dataid"].(map[string]interface{})
	// }
	// userDataIds := make(map[string]interface{})
	// if users[UserFid] != nil {
	// 	userDataIds = users[UserFid].(map[string]interface{})
	// }

	ufbddoc := make(map[string]interface{})
	ufbddoc["users-frame-by-dataid"] = map[string]interface{}{
		UserFid: map[string]interface{}{
			DataId: tfdataVal,
		},
	}
	_, err = ufbdDoc.Set(context.Background(), ufbddoc, firestore.MergeAll)
	if err != nil {
		log.Println(err)
		return PutSuccess{Success: false}
	}
	ufbbddoc := make(map[string]interface{})
	ufbbddoc[weekStr] = map[string]interface{}{
		UserFid: map[string]interface{}{
			DataId: true,
		},
	}
	_, err = ufbbdDoc.Set(context.Background(), ufbbddoc, firestore.MergeAll)
	if err != nil {
		log.Println(err)
		return PutSuccess{Success: false}
	}
	return PutSuccess{Success: true, Msg: "Succesfully, Added !!!"}
}
