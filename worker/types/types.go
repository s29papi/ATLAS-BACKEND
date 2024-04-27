package types

// N.B update the fields in cast as needed
type CastData struct {
	Hash          string        `json:"hash"`
	Text          string        `json:"text"`
	Timestamp     string        `json:"timestamp"`
	Author        Author        `json:"author"`
	Parent_Author Parent_Author `json:"parent_author"`
}

type Cast struct {
	Data CastData `json:"cast"`
}

type Author struct {
	Fid      int64  `json:"fid"`
	Username string `json:"username"`
}
type Parent_Author struct {
	Fid int64 `json:"fid"`
}

// type ChannelCasts struct {
// 	Casts []Cast `json:"casts"`
// }

// type Notification struct {
// 	Type string `json:"type"`
// 	Cast Cast   `json:"cast"`
// }

// type UserMentions struct {
// 	Notifications []Notification `json:"notifications"`
// }
// type Notifications struct {
// 	Notifications []Notification `json:"notifications"`
// }

type HubbleCastAddBody struct {
	EmbedsDeprecated  []interface{} `json:"embedsDeprecated"`
	Mentions          []int         `json:"mentions"`
	Text              string        `json:"text"`
	MentionsPositions []int         `json:"mentionsPositions"`
	Embeds            []interface{} `json:"embeds"`
}

type HubbleMessageData struct {
	Type        string            `json:"type"`
	FID         int               `json:"fid"`
	Timestamp   int               `json:"timestamp"`
	Network     string            `json:"network"`
	CastAddBody HubbleCastAddBody `json:"castAddBody"`
}

type HubbleMessagesForUserMentions struct {
	Data            HubbleMessageData `json:"data"`
	Hash            string            `json:"hash"`
	HashScheme      string            `json:"hashScheme"`
	Signature       string            `json:"signature"`
	SignatureScheme string            `json:"signatureScheme"`
	Signer          string            `json:"signer"`
}
type HubbleNotificationsForUserMentions struct {
	Messages []HubbleMessagesForUserMentions `json:"messages"`
}

type Payload struct {
	Parent_Cast_Hash string
	Channel_Id       string
	Signer_uuid      string
	Text             string
	Embeds_url       string
}

type TxType int64

const (
	USERMENTION_TX TxType = iota + 1
	NOTIFICATION_TX
	STAKE_TX
	UNSTAKE_TX
	DEPOSIT_TX
	WITHDRAW_TX
)

type Tx struct {
	Timestamp      int64
	AuthorFid      int64
	Type           TxType
	CastText       string
	CastHash       string
	AuthorUsername string
}

type ToshiPayTx struct {
	Timestamp      int64
	RecipientFid   int64
	SenderUsername string
	Type           TxType
	CastText       string
	CastHash       string
}

// curl --request GET \
//      --url 'https://api.neynar.com/v2/farcaster/cast?identifier=0x25e621250906f784a4a6eec22c0bd4d898d4564a&type=hash' \
//      --header 'accept: application/json' \
//      --header 'api_key: NEYNAR_API_DOCS'

// 	 curl --request GET \
//      --url 'https://api.neynar.com/v2/farcaster/user/search?q=raise-bot&viewer_fid=3&limit=5' \
//      --header 'accept: application/json' \
//      --header 'api_key: NEYNAR_API_DOCS'
