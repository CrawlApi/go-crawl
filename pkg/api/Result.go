package api

type UID struct {
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
	Date    int64  `json:"date"`
	Url     string `json:"url"`
	Status  bool   `json:"status"`
	Media   string `json:"media"`
	Message string `json:"message"`
}

type Profile struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	PostNum   int64  `json:"postNum"`
	FollowNum int64  `json:"followNum"`
	Fans      int64  `json:"fans"`
	Birthday  string `json:"birthday"`
	Website   string `json:"website"`
	About     string `json:"about"`
	Url       string `json:"url"`
}
