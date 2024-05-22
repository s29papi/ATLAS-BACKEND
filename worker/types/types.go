package types

type CastData struct {
	Hash          string        `json:"hash"`
	Text          string        `json:"text"`
	Timestamp     string        `json:"timestamp"`
	Author        Author        `json:"author"`
	Parent_Author Parent_Author `json:"parent_author"`
	Frames        []Frames      `json:"frames"`
}

type Casts struct {
	Data []CastData `json:"casts"`
}

type Author struct {
	Fid        int64  `json:"fid"`
	Username   string `json:"username"`
	PowerBadge bool   `json:"power_badge"`
}

type Parent_Author struct {
	Fid int64 `json:"fid"`
}

type Frames struct {
	Title     string `json:"title"`
	Image     string `json:"image"`
	FramesUrl string `json:"frames_url"`
}

type FrameData struct {
	ImageUrl       string `firestore:"image_url,omitempty"`
	FramesTitle    string `firestore:"frames_title,omitempty"`
	AuthorUserName string `firestore:"author_username,omitempty"`
	Text           string `firestore:"text,omitempty"`
	Week           string `firestore:"week,omitempty"`
	DataId         string `firestore:"dataid,omitempty"`
	Timestamp      int64  `firestore:"timestamp,omitempty"`
	AuthorFid      int64  `firestore:"author_fid,omitempty"`
}

type CurrentReqInfo struct {
	TimeStamp int64
	Hash      string
	DataId    int64
}

type DomainData struct {
	NFTDomain []string `json:"nftdomain"`
}

type BulkFollowingRequestBody struct {
	ViewerFid string `json:"viewer_fid"`
}

type ViewerCtx struct {
	Following bool `json:"following"`
}

type UserData struct {
	ViewerCtx ViewerCtx `json:"viewer_context"`
}

type UsersData struct {
	UserData []UserData `json:"users"`
}
