package types

import "time"

// N.B update the fields in cast as needed
type Cast struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type ChannelCasts struct {
	Casts []Cast `json:"casts"`
}
