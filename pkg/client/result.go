package client

type FaceBook struct {
	Content      interface{} `json:"content"`
	Picture      interface{} `json:"picture"`
	Full_picture interface{} `json:"full_picture"`
	Shares       interface{} `json:"shares"`
	Name         interface{} `json:"name"`
	UpdatedTime  interface{} `json:"updated_time"`
	CreatedTime  interface{} `json:"created_time"`
	Likes        interface{} `json:"likes"`
	Comments     interface{} `json:"comments"`
	Id           interface{} `json:"id"`
	UserId       interface{} `json:"user_id"`
	Source       interface{} `json:"source"`
	Type         interface{} `json:"type"`
	Message      interface{} `json:"message"`
	Status       interface{} `json:"status"`
	Date         interface{} `json:"date"`
}

type FaceBookPost struct {
	Post    *FaceBook   `json:"post"`
	Message interface{} `json:"message"`
	Status  interface{} `json:"status"`
	Date    interface{} `json:"date"`
}

type TokenStr struct {
	Url         string   `json:"url"`
	Id          string   `json:"id"`
	Type        string   `json:"type"`
	AccessToken string   `json:"access_token"`
	Ids         []string `json:"ids"`
}

func (t *TokenStr) FetchData() {

	switch t.Type {
	case "facebook":
	case "weibo":
	}
}
