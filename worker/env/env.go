package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AUTH_TOKEN string

	CASTWITHFRAMES_URL string

	SERVICE_ACCOUNT_JSON_HEX string

	API_KEY string

	BULKFOLLOWING_URL string
)

func loadEnv() {
	godotenv.Load()
}

func init() {
	loadEnv()

	AUTH_TOKEN = os.Getenv("AUTH_TOKEN")

	CASTWITHFRAMES_URL = os.Getenv("CASTWITHFRAMES_URL")

	SERVICE_ACCOUNT_JSON_HEX = os.Getenv("SERVICE_ACCOUNT_JSON_HEX")

	API_KEY = os.Getenv("API_KEY")

	BULKFOLLOWING_URL = os.Getenv("BULKFOLLOWING_URL")
}
