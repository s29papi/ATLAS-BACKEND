package types

// N.B update the fields in cast as needed
type Cast struct {
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
