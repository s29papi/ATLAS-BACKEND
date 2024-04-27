package env

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

var (
	DURATION_STR                    string
	CHANNEL_CAST_URL                string
	USER_MENTIONS_URL               string
	MENTIONS_REPLY_URL              string
	API_KEY                         string
	API_KEY_TOSHI                   string
	CHANNEL_ID                      string
	SIGNER_UUID                     string
	RENDER_POSTGRES_URL             string
	FRAMES_URL                      string
	AUTH_TOKEN                      string
	NotificationsForToshiPayBot_URL string
	DataForCastHash_URL             string
	SERVICE_ACCOUNT_JSON_HEX        string
	STARTING_BLOCK_HASH             common.Hash
	PRIZE_POOL_ADDRESS              common.Address
	EVENT_ETH_DEP_SIG               common.Hash
)

func loadEnv() {
	godotenv.Load()
}

func init() {
	loadEnv()

	// DURATION_STR is the interval in seconds duration of bot request for user mentions.
	DURATION_STR = os.Getenv("REQUEST_DURATION")
	// // CHANNEL_CAST_URL is the endpoint to fetch casts in  a channel.
	CHANNEL_CAST_URL = os.Getenv("CHANNEL_CAST_URL")
	// // USER_MENTIONS_URL is the endpoint to fetch user mentions.
	USER_MENTIONS_URL = os.Getenv("USER_MENTIONS_URL")

	MENTIONS_REPLY_URL = os.Getenv("MENTIONS_REPLY_URL")

	API_KEY = os.Getenv("API_KEY")

	CHANNEL_ID = os.Getenv("CHANNEL_ID")

	SIGNER_UUID = os.Getenv("SIGNER_UUID")

	RENDER_POSTGRES_URL = os.Getenv("RENDER_POSTGRES_URL")

	FRAMES_URL = os.Getenv("FRAMES_URL")

	AUTH_TOKEN = os.Getenv("AUTH_TOKEN")

	DataForCastHash_URL = os.Getenv("DataForCastHash_URL")

	API_KEY_TOSHI = os.Getenv("API_KEY_TOSHI")

	SERVICE_ACCOUNT_JSON_HEX = os.Getenv("SERVICE_ACCOUNT_JSON_HEX")

	NotificationsForToshiPayBot_URL = os.Getenv("NotificationsForToshiPayBot_URL")

	STARTING_BLOCK_HASH = common.HexToHash(os.Getenv("STARTING_BLOCK_HASH"))

	PRIZE_POOL_ADDRESS = common.HexToAddress(os.Getenv("PRIZE_POOL_ADDRESS"))

	EVENT_ETH_DEP_SIG = common.HexToHash(os.Getenv("EVENT_ETH_DEP_SIG"))
}
