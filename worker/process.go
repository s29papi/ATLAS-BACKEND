package worker

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/s29papi/wag3r-bot/worker/client"
	"github.com/s29papi/wag3r-bot/worker/db"
	"github.com/s29papi/wag3r-bot/worker/env"
	"github.com/s29papi/wag3r-bot/worker/types"
)

// func fetchVersusNotificationsInStadium(service *client.HTTPService) types.Notifications {
// 	req := client.NotificationsInStadiumRequest()
// 	resp := service.SendRequest(http.MethodGet, req)
// 	var notifs types.Notifications
// 	if err := json.Unmarshal(resp, &notifs); err != nil {
// 		log.Println(err)
// 		return types.Notifications{}
// 	}
// 	return notifs
// }

// func filterNotificationsByLastUpdate(m types.Notifications, lastUserMentionUpdateTime int64) types.Notifications {
// 	var notifications []types.Notification
// 	for _, notif := range m.Notifications {
// 		notifTimeInt64 := notificationsTimestamp2secs(notif.Cast.Timestamp)
// 		if notifTimeInt64 <= lastUserMentionUpdateTime {
// 			continue
// 		}
// 		notifications = append(notifications, notif)
// 	}
// 	return types.Notifications{Notifications: notifications}
// }

// func filterNotificationsByMentions(m types.Notifications) types.Notifications {
// 	var notifications []types.Notification
// 	for _, notif := range m.Notifications {
// 		if notif.Type != "mention" {
// 			continue
// 		}
// 		notifications = append(notifications, notif)
// 	}
// 	return types.Notifications{Notifications: notifications}
// }

// func (w *Worker) buildUserMentionToTx() []types.Tx {
// 	var txs []types.Tx
// 	notifications := fetchVersusNotificationsInStadium(w.s)
// 	// lastusermentionsupdatetime := w.db.LastUserMentionsRequestTime() // gets last row of a request table
// 	lastusermentionsupdatetime := int64(0)
// 	notifsByLastUpdate := filterNotificationsByLastUpdate(notifications, lastusermentionsupdatetime)
// 	notifsByMention := filterNotificationsByMentions(notifsByLastUpdate)
// 	for _, notif := range notifsByMention.Notifications {
// 		tx := types.Tx{
// 			Timestamp: notificationsTimestamp2secs(notif.Cast.Timestamp),
// 			Type:      types.NOTIFICATION_TX,
// 			CastText:  notif.Cast.Text,
// 			CastHash:  notif.Cast.Hash,
// 			AuthorFid: notif.Cast.Author.Fid,
// 		}
// 		txs = append(txs, tx)
// 	}
// 	return txs
// }

func (w *Worker) processTx(fn func() []types.Tx) {
	txArr := fn()
	// var wg sync.WaitGroup
	for _, tx := range txArr {
		if tx.Type == types.NOTIFICATION_TX {
			fmt.Println("Hello World")
			// 		replyData := processUserMentionTx(tx, w.db)
			// 		fmt.Println("We dey here ooo 4444")
			// 		if replyData != nil {
			// 			wg.Add(1)
			// 			go func(r *strings.Reader, t int64) {
			// 				defer wg.Done()
			// 				if t == 1710859059 {
			// 					// respData := castNewReply(w.s, r)
			// 					// fmt.Println(string(respData))
			// 					// respond that thge
			// 				}
			// 			}(replyData, tx.Timestamp)
			// 		}
		}
	}
	// wg.Wait()
	fmt.Println(len(txArr))
	fmt.Println("We dey here ooo 3333")
}

func processUserMentionTx(tx types.Tx, db *db.DB) *strings.Reader {
	var game = &types.Game{}
	errNo := process(tx.CastText, game)
	switch errNo {
	case ERR_MAX_NUMBER_LINES_NO:
		log.Println(ERR_MAX_NUMBER_LINES)
		return nil
	case ERR_UNEXPECTED_FIELD_NO:
		log.Println(ERR_UNEXPECTED_FIELD)
		return nil
	case ERR_MISSING_CURRENCY_SYMBOL_NO:
		log.Println(ERR_MISSING_CURRENCY_SYMBOL)
		return nil
	case ERR_INVALID_LENGTH_WAGERAMOUNT_NO:
		log.Println(ERR_INVALID_LENGTH_WAGERAMOUNT)
		return nil
	case ERR_INVALID_AMOUNT_NO:
		log.Println(ERR_INVALID_AMOUNT)
		return nil
	case ERR_MISSING_REQ_FIELD_NO:
		log.Println(ERR_MISSING_REQ_FIELD)
		return nil
	case 0:
		latestgameId := db.LatestGameID()
		latestgameId += 1
		fmt.Println(game.Date)

		// payload := &types.Payload{
		// 	Parent:      tx.CastHash,
		// 	Channel_Id:  env.CHANNEL_ID,
		// 	Signer_uuid: env.SIGNER_UUID,
		// 	Text:        "Open /stadium Challenge Accepted: ",
		// 	// let queryParams = `gameId=${gameId}&&gameName=${gameName}&&gameSetup=${gameSetup}&&stakeAmount=${stakeAmount}&&creatorFid=${creatorFid}`
		// 	Embeds_url: buildEmbedUrl(latestgameId, game.Name, game.Setup, game.Token, game.Amount, tx.AuthorFid),
		// }

		log.Printf("[processUserMentionTx] New payload created")

		// update db

		// fmt.Println(tx.Timestamp)

		// return buildCastReply(payload)
	}

	return nil
}

func sortHighestTime(txArr []types.Tx) int64 {
	for _, tx := range txArr {
		if txArr[0].Timestamp < tx.Timestamp {
			txArr[0].Timestamp = tx.Timestamp
		}
	}
	return txArr[0].Timestamp
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
					date := gameTimestamp2secs(value)
					if date == 0 {
						break
					}
					info.Date = date
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

	if info.Name == "" || info.Setup == "" || info.Date == 0 || info.Amount == 0 || info.Token == "" {
		return ERR_MISSING_REQ_FIELD_NO
	}
	return 0 // sucess
}

func GetMentions(service *client.HTTPService) []byte {
	req := client.UserMentionsRequest()
	return service.SendRequest(http.MethodGet, req)
}

func notificationsTimestamp2secs(t string) int64 {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	return parsedTime.Unix()
}

func gameTimestamp2secs(t string) int64 {
	fmt.Println(t)
	if !strings.Contains(t, ",") || !strings.Contains(t, " ") {
		t += ", " + time.Now().Format("2006")
	}
	layout := "Jan 02 at 3pm, 2006"
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	return parsedTime.Unix()
}

func buildEmbedUrl(gameId int64, gameName, gameSetup, token string, amount float64, creatorFid int64) string {
	baseUrl := env.FRAMES_URL
	queryParams := url.Values{
		"gameId":      {fmt.Sprintf("%d", gameId)},
		"gameName":    {gameName},
		"gameSetup":   {gameSetup},
		"stakeAmount": {fmt.Sprintf("%g %s", amount, token)},
		"creatorFid":  {fmt.Sprintf("%d", creatorFid)},
	}.Encode()
	urlStr := baseUrl + "/?" + queryParams
	return urlStr
}

// func buildCastReply(payload *types.Payload) *strings.Reader {
// 	var payloadString string
// 	if len(payload.Embeds_url) == 0 {
// 		payloadString = fmt.Sprintf(`
// 		{
// 		"parent": "%s",
// 		"channel_id": "%s",
// 		"signer_uuid": "%s",
// 		"text": "%s"
// 		}`, payload.Parent, payload.Channel_Id, payload.Signer_uuid, payload.Text)
// 	} else {
// 		payloadString = fmt.Sprintf(`
// 		{
// 		"parent": "%s",
// 		"channel_id": "%s",
// 		"signer_uuid": "%s",
// 		"text": "%s",
// 		"embeds": [{"url": "%s"}]
// 		}`, payload.Parent, payload.Channel_Id, payload.Signer_uuid, payload.Text, payload.Embeds_url)
// 	}

// 	return strings.NewReader(payloadString)
// }

func castNewReply(service *client.HTTPService, data *strings.Reader) []byte {
	req := client.MentionReplyRequest(data)
	return service.SendRequest(http.MethodPost, req)
}

// func build usermentions to tx
// check if a new cast has been added
// sunday
// query new game Id
// query previous time
// func (w *Worker) process(d []byte) {
// 	var userMentions types.UserMentions
// 	if err := json.Unmarshal(d, &userMentions); err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	lastProcReqTime := w.db.GetLastProcReqTime()
// 	lastProcReqGameId := w.db.GetLastProcReqGameId()
// 	fmt.Println(lastProcReqGameId)

// 	var currentTime int64 // move to db
// 	var wg sync.WaitGroup
// 	for idx, notifs := range userMentions.Notifications {
// 		t := timestamp2secs(notifs.Cast.Timestamp)

// 		if idx == 0 && t > lastProcReqTime {
// 			currentTime = t
// 		}

// 		// if t <= lastProcReqTime {
// 		// 	break
// 		// }

// 		if notifs.Type != "mention" {
// 			continue
// 		}

// 		// parse text
// 		var game = &types.Game{}
// 		errNo := process(notifs.Cast.Text, game)

// 		// handle errors properly
// 		switch errNo {
// 		case ERR_MAX_NUMBER_LINES_NO:
// 		}

// 		payload := &types.Payload{
// 			Parent:      notifs.Cast.Hash,
// 			Channel_Id:  env.CHANNEL_ID,
// 			Signer_uuid: env.SIGNER_UUID,
// 		}

// 		if errNo == 0 {
// 			game.Id = lastProcReqGameId + 1
// 			game.Url = "https://wag3r-bot-gamma.vercel.app/?gameId=" + strconv.Itoa(game.Id)
// 			payload.Text = "Open /stadium Challenge Accepted: "
// 			payload.Embeds_url = game.Url
// 			lastProcReqGameId += game.Id
// 		}

// 		if errNo != 0 {
// 			payload.Text = "Wrong Message Format."
// 		}

// 		wg.Add(1)
// 		go func(p *types.Payload) {
// 			defer wg.Done()
// 			data := buildCastReply(p)
// 			fmt.Println(data)
// 			// castNewReply(w.s, data)
// 			// after success save game
// 			// add gameid save here
// 		}(payload)

// 	}
// 	wg.Wait()

// 	if currentTime > lastProcReqTime {
// 		w.db.UpdateLastProcReqTime(currentTime, lastProcReqTime)
// 		w.db.UpdateLastProcReqGameId(lastProcReqGameId)
// 	}
// }
// func checkNewCast(service *client.HTTPService) {
// 	req := client.ChannelCastRequest()
// 	go service.SendRequest(http.MethodGet, req)
// }

// var game = &types.Game{}
// errNo := process(tx.CastText, game)
// fmt.Println(errNo)
// timestampHighest := sortHighestTime(txArr) // store it lastely
// fmt.Println(timestampHighest)
// fmt.Println(tx.Timestamp)
