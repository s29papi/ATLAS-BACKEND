package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DURATION_STR     string
	CHANNEL_CAST_URL string
)

func loadEnv() {
	godotenv.Load()
}

func init() {
	loadEnv()

	DURATION_STR = os.Getenv("REQUEST_DURATION")
	CHANNEL_CAST_URL = os.Getenv("CHANNEL_CAST_URL")
}
