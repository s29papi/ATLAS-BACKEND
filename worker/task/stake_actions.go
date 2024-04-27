package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/s29papi/wag3r-bot/worker/client"
	"github.com/s29papi/wag3r-bot/worker/env"
	"github.com/s29papi/wag3r-bot/worker/types"
)

func fetchToshiPayBotNotifications(service *client.HTTPService) types.HubbleNotificationsForUserMentions {
	req := client.NotificationsForToshiPayBot()
	resp := service.SendRequest(http.MethodGet, req)
	var notifs types.HubbleNotificationsForUserMentions // types notifications
	if err := json.Unmarshal(resp, &notifs); err != nil {
		log.Println(err)
		return types.HubbleNotificationsForUserMentions{}
	}
	return notifs
}

func filterToshiPayBotNotificationsByLastUpdate(m types.HubbleNotificationsForUserMentions, lastNotificationsUpdateTime int64) types.HubbleNotificationsForUserMentions {
	var filteredNotifications types.HubbleNotificationsForUserMentions

	for _, hubbleMessage := range m.Messages {
		if int64(hubbleMessage.Data.Timestamp) <= lastNotificationsUpdateTime {
			continue
		}
		filteredNotifications.Messages = append(filteredNotifications.Messages, hubbleMessage)
	}

	return filteredNotifications
}

func filterNotificationsByNetwork(m types.HubbleNotificationsForUserMentions) types.HubbleNotificationsForUserMentions {
	var filteredNotifications types.HubbleNotificationsForUserMentions
	for _, hubbleMessage := range m.Messages {
		if hubbleMessage.Data.Network != "FARCASTER_NETWORK_MAINNET" {
			continue
		}
		filteredNotifications.Messages = append(filteredNotifications.Messages, hubbleMessage)
	}
	return filteredNotifications
}

func notifs2Tx(service *client.HTTPService, messages []types.HubbleMessagesForUserMentions) []types.ToshiPayTx {
	var txs []types.ToshiPayTx
	for _, message := range messages {
		if message.Data.CastAddBody.Mentions[0] != 502736 {
			continue
		}
		if len(message.Data.CastAddBody.Mentions) > 2 {
			continue
		}
		castData := getCastData(service, message.Hash)
		tx := types.ToshiPayTx{
			Timestamp:      int64(message.Data.Timestamp),
			Type:           types.NOTIFICATION_TX,
			CastText:       message.Data.CastAddBody.Text,
			CastHash:       message.Hash,
			SenderUsername: castData.Author.Username,
		}

		if len(message.Data.CastAddBody.Mentions) == 2 {
			tx.RecipientFid = int64(message.Data.CastAddBody.Mentions[1])
		} else {
			if castData.Parent_Author.Fid != 0 {
				tx.RecipientFid = castData.Parent_Author.Fid
			} else {
				tx.RecipientFid = int64(message.Data.FID)
			}
		}
		txs = append(txs, tx)
	}
	return txs
}

func getCastData(service *client.HTTPService, hash string) types.CastData {
	req := client.DataForCastHash(hash)
	resp := service.SendRequest(http.MethodGet, req)
	var cast types.Cast // types notifications
	if err := json.Unmarshal(resp, &cast); err != nil {
		log.Println(err)
		return types.CastData{}
	}
	return cast.Data
}

func getValidToshiPayPayload(t []types.ToshiPayTx) []types.Payload {
	var payloads []types.Payload
	for _, tx := range t {
		valid, info := checkValidToshiPayTxText(tx.CastText)
		if valid {
			text := fmt.Sprintf("@%v %d Toshi Tip Frame Created, meow 🎉🎉🎉", tx.SenderUsername, info.Amount)
			payload := types.Payload{
				Parent_Cast_Hash: tx.CastHash,
				Signer_uuid:      env.SIGNER_UUID,
				Text:             text,
				Embeds_url:       buildToshiPayEmbedUrl(info.Amount, strconv.Itoa(int(tx.RecipientFid))),
			}
			payloads = append(payloads, payload)
		}
	}
	return payloads
}
func checkValidToshiPayTxText(s string) (valid bool, info types.ToshiPayInfo) {
	mentionReg := regexp.MustCompile(`(?:^|@(\w+)\s+)?(\d+)\s*TOSHI(?:\s+@(\w+))?`)
	matches := mentionReg.FindStringSubmatch(s)

	if len(matches) >= 3 {
		amount, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			return false, types.ToshiPayInfo{}
		}

		toshiPayInfo := types.ToshiPayInfo{
			Amount: amount,
		}
		return true, toshiPayInfo
	}

	return false, types.ToshiPayInfo{}
}

func buildToshiPayEmbedUrl(amount int64, recipientFid string) string {
	baseUrl := env.FRAMES_URL
	queryParams := url.Values{
		"amount": {strconv.Itoa(int(amount))},
		"fid":    {recipientFid},
		"tip":    {"true"},
	}.Encode()
	urlStr := baseUrl + "/?" + queryParams
	return urlStr
}

func buildToshiPayCastReply(p types.Payload) *strings.Reader {
	var payloadString string
	if len(p.Embeds_url) == 0 {
		payloadString = fmt.Sprintf(`
		{
		"parent": "%s",
		"signer_uuid": "%s",
		"text": "%s"
		}`, p.Parent_Cast_Hash, p.Signer_uuid, p.Text)
	} else {
		payloadString = fmt.Sprintf(`
		{
		"parent": "%s",
		"signer_uuid": "%s",
		"text": "%s",
		"embeds": [{"url": "%s"}]
		}`, p.Parent_Cast_Hash, p.Signer_uuid, p.Text, p.Embeds_url)
	}

	return strings.NewReader(payloadString)
}

func castToshiPayReplies(s *client.HTTPService, p []types.Payload) []byte {
	for _, payload := range p {
		data := buildToshiPayCastReply(payload)
		req := client.MentionReplyRequest(data)
		s.SendRequest(http.MethodPost, req)
	}
	return nil
}

func notificationsTimestamp2secs(t string) int64 {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println("Error parsing timestamp:", err)
		return 0
	}
	return parsedTime.Unix()
}

func sortHighestTime(txArr []types.ToshiPayTx) int64 {
	for _, tx := range txArr {
		if txArr[0].Timestamp < tx.Timestamp {
			txArr[0].Timestamp = tx.Timestamp
		}
	}
	return txArr[0].Timestamp
}

func updatelastnotificationsupdatetime(ctx context.Context, client *firestore.Client, t int64) error {
	_, err := client.Collection("Toshi-TIp-Farc-Request").Doc("Toshi-TIp-Farc-Request").Update(ctx, []firestore.Update{
		{
			Path:  "Req-TimeStamp",
			Value: t,
		},
	})

	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	return err
}

func getCurrentReqTimeStamp(client *firestore.Client) int64 {
	dsnap, err := client.Collection("Toshi-TIp-Farc-Request").Doc("Toshi-TIp-Farc-Request").Get(context.Background())
	if err != nil {
		log.Println(err)
	}
	m := dsnap.Data()
	if intValue, ok := m["Req-TimeStamp"].(int64); ok {
		return intValue
	}
	return 0
}
