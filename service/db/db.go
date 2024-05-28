package db

import (
	"context"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/s29papi/atlas-backend/worker/env"
	"google.golang.org/api/option"
)

type DB struct {
	fc *firestore.Client
}

func NewDB() *DB {
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

	return &DB{
		fc: fc,
	}
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
	FramesUrl      string `json:"frames_url"`
	FramesTitle    string `json:"frames_title"`
	AuthorUserName string `json:"author_username"`
	AuthorFid      int64  `json:"author_fid"`
	Text           string `json:"text"`
	Week           string `json:"week"`
	DataId         string `json:"dataId"`
	Timestamp      int64  `json:"timestamp"`
}
