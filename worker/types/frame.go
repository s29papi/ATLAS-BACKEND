package types

// update to have creator
type Game struct {
	Id         int
	Name       string
	Setup      string
	Date       int64
	Token      string
	Amount     float64
	Url        string
	CreatorFid string
}

type Info struct {
	Id         int
	Name       string
	Setup      string
	Date       int64
	Token      string
	Amount     float64
	Url        string
	CreatorFid string
}
type ToshiPayInfo struct {
	Amount   int64
	Username string
}
