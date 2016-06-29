package data

type ApiJson struct {
	Domain	string `json:"domain"`
	Url           string            `json:"url"`
	StartSelector string            `json:"start_selector"`
	Datas         map[string]string `json:"datas"`
	NextPage      string            `json:"next_page"`
	Limit         int               `json:"limit"`
}

type ResultJson struct {
	Url   string        `json:"url"`
	Datas []interface{} `json:"datas"`
}
