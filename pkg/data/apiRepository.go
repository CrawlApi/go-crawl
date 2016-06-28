package data

type ApiJson struct {
	Url string `json:"url"`
	StartSelector string `json:"start_selector"`
	Datas map[string]string `json:"datas"`
}


type ResultJson struct {
	Url string `json:"url"`
	Datas []interface{} `json:"datas"`
}