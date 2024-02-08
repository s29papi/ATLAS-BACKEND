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

// 	defer res.Body.Close()
// body, _ := io.ReadAll(res.Body)

// 	// fmt.Println(string(body))

// 	var stadiumCasts StadiumCasts
// 	err := json.Unmarshal(body, &stadiumCasts)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON:", err)
// 		return
// 	}
// 	fmt.Println(stadiumCasts.Casts[0].Text)
// 	fmt.Println(stadiumCasts.Casts[0].Timestamp.UnixMicro())
// 	fmt.Println(time.Now().UnixMicro())

// }
