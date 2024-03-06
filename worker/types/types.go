package types

// N.B update the fields in cast as needed
type Cast struct {
	Hash      string `json:"hash"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type ChannelCasts struct {
	Casts []Cast `json:"casts"`
}

type Notification struct {
	Type string `json:"type"`
	Cast Cast   `json:"cast"`
}

type UserMentions struct {
	Notifications []Notification `json:"notifications"`
}

type Payload struct {
	Parent      string
	Channel_Id  string
	Signer_uuid string
	Text        string
	Embeds_url  string
}

type TxType int64

const (
	USERMENTION_TX TxType = iota + 1
	STAKE_TX
	UNSTAKE_TX
	DEPOSIT_TX
	WITHDRAW_TX
)

type Tx struct {
	BatchNo   int64
	BatchIdx  int64
	Timestamp int64
	Type      TxType
	CastText  string
	CastHash  string
}
