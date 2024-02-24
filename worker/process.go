package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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

	var currentTime int64 // move to db
	var wg sync.WaitGroup
	for idx, notifs := range userMentions.Notifications {
		t := timestamp2secs(notifs.Cast.Timestamp)

		if idx == 0 && t > *w.lastProcReqTime {
			currentTime = t
		}

		if t <= *w.lastProcReqTime {
			break
		}

		if notifs.Type != "mention" {
			continue
		}

		// parse text
		var info = &types.Game{}
		errNo := process(notifs.Cast.Text, info)
		switch errNo {
		case ERR_MAX_NUMBER_LINES_NO:
		}

		payload := &types.Payload{
			Parent:      notifs.Cast.Hash,
			Channel_Id:  env.CHANNEL_ID,
			Signer_uuid: env.SIGNER_UUID,
		}

		if errNo == 0 {
			payload.Text = "Open /stadium Challenge Accepted: \r\n"
			payload.Embeds_url = "https://wag3r-bot-gamma.vercel.app/"
		}

		if errNo != 0 {
			payload.Text = "Wrong Message Format."
		}

		wg.Add(1)
		go func(p *types.Payload) {
			defer wg.Done()
			data := buildCastReply(p)
			fmt.Println(data)
			castNewReply(w.s, data)
		}(payload)

	}
	wg.Wait()
	if currentTime > *w.lastProcReqTime {
		w.lastProcReqTime = &currentTime
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
					wagerAmountStr := strings.TrimSpace(value)
					if !strings.HasPrefix(wagerAmountStr, "$") {
						return ERR_MISSING_CURRENCY_SYMBOL_NO
					}

					wagerAmountParts := strings.Fields(value)
					if len(wagerAmountParts) != 3 {
						return ERR_INVALID_LENGTH_WAGERAMOUNT_NO
					}

					amountStr := wagerAmountParts[1]
					token := wagerAmountParts[2]
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
	payloadString := fmt.Sprintf(`
						{
						"parent": "%s",
						"channel_id": "%s",
						"signer_uuid": "%s",
						"text": "%s",
						"embeds": [{"url": "%s"}]
						}`, payload.Parent, payload.Channel_Id, payload.Signer_uuid, payload.Text, payload.Embeds_url)

	return strings.NewReader(payloadString)
}

func castNewReply(service *client.HTTPService, data *strings.Reader) {
	req := client.MentionReplyRequest(data)
	service.SendRequest(http.MethodPost, req)
}
