package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/s29papi/wag3r-bot/worker/client"
	"github.com/s29papi/wag3r-bot/worker/env"
	"github.com/s29papi/wag3r-bot/worker/types"
)

// check if a new cast has been added
// sunday
// query new game Id
// query previous time
func (w *Worker) process(d []byte) {
	var userMentions types.UserMentions
	if err := json.Unmarshal(d, &userMentions); err != nil {
		log.Println(err)
		return
	}
	lastProcReqTime := w.db.GetLastProcReqTime()
	lastProcReqGameId := w.db.GetLastProcReqGameId()
	fmt.Println(lastProcReqGameId)

	var currentTime int64 // move to db
	var wg sync.WaitGroup
	for idx, notifs := range userMentions.Notifications {
		t := timestamp2secs(notifs.Cast.Timestamp)

		if idx == 0 && t > lastProcReqTime {
			currentTime = t
		}

		// if t <= lastProcReqTime {
		// 	break
		// }

		if notifs.Type != "mention" {
			continue
		}

		// parse text
		var game = &types.Game{}
		errNo := process(notifs.Cast.Text, game)

		// handle errors properly
		switch errNo {
		case ERR_MAX_NUMBER_LINES_NO:
		}

		payload := &types.Payload{
			Parent:      notifs.Cast.Hash,
			Channel_Id:  env.CHANNEL_ID,
			Signer_uuid: env.SIGNER_UUID,
		}

		if errNo == 0 {
			game.Id = lastProcReqGameId + 1
			game.Url = "https://wag3r-bot-gamma.vercel.app/?gameId=" + strconv.Itoa(game.Id)
			payload.Text = "Open /stadium Challenge Accepted: "
			payload.Embeds_url = game.Url
			lastProcReqGameId += game.Id
		}

		if errNo != 0 {
			payload.Text = "Wrong Message Format."
		}

		wg.Add(1)
		go func(p *types.Payload) {
			defer wg.Done()
			data := buildCastReply(p)
			fmt.Println(data)
			// castNewReply(w.s, data)
			// after success save game
			// add gameid save here
		}(payload)

	}
	wg.Wait()

	if currentTime > lastProcReqTime {
		w.db.UpdateLastProcReqTime(currentTime, lastProcReqTime)
		w.db.UpdateLastProcReqGameId(lastProcReqGameId)
	}
}

// process takes a cast's text and returns either nil if successful and error
// if it fails.
func process(text string, info *types.Game) int64 {
	reg := regexp.MustCompile(`^\s*\*\s*(.+?):\s*(.*)$`)
	mentionReg := regexp.MustCompile(`@\w+`)
	textLines := strings.Split(text, "\n")

	if len(textLines) >= MAXNUMBERLINES {
		return int64(ERR_MAX_NUMBER_LINES_NO)
	}

	var challengeStarted bool

	for _, textLine := range textLines {
		textLine = strings.TrimSpace(textLine)

		if textLine == "" {
			continue
		}

		if !challengeStarted {
			if textLine == "Open /stadium Challenge:" {
				challengeStarted = true
			}
			continue
		}

		textLine = mentionReg.ReplaceAllString(textLine, "")

		if strings.HasPrefix(textLine, "*") {
			match := reg.FindStringSubmatch(textLine)
			if len(match) == 3 {
				key := strings.Title(match[1])
				value := match[2]

				switch key {
				case "Game":
					info.Name = value
				case "Match Set Up":
					info.Setup = value
				case "Match Date":
					info.Date = value
				case "Wager Amount":
					var wagerAmount string
					if wagerAmountStr := strings.TrimSpace(value); strings.HasPrefix(wagerAmountStr, "$") {
						wagerAmount = strings.TrimPrefix(wagerAmountStr, "$")
					}
					wagerAmountParts := strings.Fields(wagerAmount)
					if len(wagerAmountParts) == 0 {
						return ERR_INVALID_LENGTH_WAGERAMOUNT_NO
					}
					amountStr := wagerAmountParts[0]
					token := wagerAmountParts[1]
					amount, err := strconv.ParseFloat(amountStr, 64)
					if err != nil {
						return ERR_INVALID_AMOUNT_NO
					}
					info.Amount = amount
					info.Token = token
				default:
					return ERR_UNEXPECTED_FIELD_NO

				}
			}
		}
	}

	if info.Name == "" || info.Setup == "" || info.Date == "" || info.Amount == 0 || info.Token == "" {
		return ERR_MISSING_REQ_FIELD_NO
	}
	return 0 // sucess
}

func checkNewCast(service *client.HTTPService) {
	req := client.ChannelCastRequest()
	go service.SendRequest(http.MethodGet, req)
}

func GetMentions(service *client.HTTPService) []byte {
	req := client.UserMentionsRequest()
	return service.SendRequest(http.MethodGet, req)
}

func timestamp2secs(t string) int64 {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	return parsedTime.Unix()
}

func buildCastReply(payload *types.Payload) *strings.Reader {
	var payloadString string
	if len(payload.Embeds_url) == 0 {
		payloadString = fmt.Sprintf(`
		{
		"parent": "%s",
		"channel_id": "%s",
		"signer_uuid": "%s",
		"text": "%s"
		}`, payload.Parent, payload.Channel_Id, payload.Signer_uuid, payload.Text)
	} else {
		payloadString = fmt.Sprintf(`
		{
		"parent": "%s",
		"channel_id": "%s",
		"signer_uuid": "%s",
		"text": "%s",
		"embeds": [{"url": "%s"}]
		}`, payload.Parent, payload.Channel_Id, payload.Signer_uuid, payload.Text, payload.Embeds_url)
	}

	return strings.NewReader(payloadString)
}

func castNewReply(service *client.HTTPService, data *strings.Reader) {
	req := client.MentionReplyRequest(data)
	service.SendRequest(http.MethodPost, req)
}

func fetchUserMentions(service *client.HTTPService) types.UserMentions {
	req := client.UserMentionsRequest()
	resp := service.SendRequest(http.MethodGet, req)
	var userMentions types.UserMentions
	if err := json.Unmarshal(resp, &userMentions); err != nil {
		log.Println(err)
		return types.UserMentions{}
	}
	return userMentions
}

// we create a table that we update for last user mention, we fetch the time from here
func filterUserMentionsByLastUpdate(m types.UserMentions, lastUserMentionUpdateTime int64) types.UserMentions {
	var notifications []types.Notification
	for _, notif := range m.Notifications {
		notifTimeInt64 := timestamp2secs(notif.Cast.Timestamp)
		if notifTimeInt64 <= lastUserMentionUpdateTime {
			continue
		}
		notifications = append(notifications, notif)
	}
	return types.UserMentions{Notifications: notifications}
}

func (w *Worker) buildUserMentionToTx() []types.Tx {
	var txs []types.Tx
	userMentions := fetchUserMentions(w.s)
	lastusermentionsupdatetime := w.db.GetLastUserMentionUpdateTime()
	userMentionsByLastUpdate := filterUserMentionsByLastUpdate(userMentions, lastusermentionsupdatetime)
	for _, notif := range userMentionsByLastUpdate.Notifications {
		tx := types.Tx{
			Timestamp: timestamp2secs(notif.Cast.Timestamp),
			Type:      types.USERMENTION_TX,
			CastText:  notif.Cast.Text,
			CastHash:  notif.Cast.Hash,
		}
		txs = append(txs, tx)
	}
	return txs
}

// func

func fetchEthDepositsFromLastUpdate(c *ethclient.Client, ctx context.Context, lastEthDepositBlock *big.Int, lastEthDepositTime int64) {
	latestBlockNo, err := c.BlockNumber(ctx)
	if err != nil {
		log.Printf("Error Getting the Current Block Time Stamp: %v", err)
	}
	latestBlockNoBig := big.NewInt(0)
	latestBlockNoBig.SetUint64(latestBlockNo)
	latestBlock, err := c.BlockByNumber(ctx, latestBlockNoBig)
	if err != nil {
		log.Printf("Error Getting the Latest Block: %v", err)
	}
	latestBlockHash := latestBlock.Hash()
	// starting
	if lastEthDepositBlock == nil && lastEthDepositTime == 0 {
		filterQuery := ethereum.FilterQuery{
			BlockHash: &latestBlockHash,
			Addresses: []common.Address{env.PRIZE_POOL_ADDRESS},
			Topics:    [][]common.Hash{[]common.Hash{env.EVENT_ETH_DEP_SIG}},
		}

		logs, err := c.FilterLogs(ctx, filterQuery)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(logs)
	}
}

// func build usermentions to tx
