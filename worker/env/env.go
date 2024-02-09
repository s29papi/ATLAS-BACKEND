package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DURATION_STR      string
	CHANNEL_CAST_URL  string
	USER_MENTIONS_URL string
)

func loadEnv() {
	godotenv.Load()
}

func init() {
	loadEnv()

	// DURATION_STR is the interval in seconds duration of bot request for user mentions.
	DURATION_STR = os.Getenv("REQUEST_DURATION")
	// CHANNEL_CAST_URL is the endpoint to fetch casts in  a channel.
	CHANNEL_CAST_URL = os.Getenv("CHANNEL_CAST_URL")
	// USER_MENTIONS_URL is the endpoint to fetch user mentions.
	USER_MENTIONS_URL = os.Getenv("USER_MENTIONS_URL")
}
