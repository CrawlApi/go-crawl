package models

type YTBModel struct {
	Datas []struct {
		Title      string `json:"title"`
		URL        string `json:"url"`
		ViewCounts string `json:"view_counts"`
	} `json:"datas"`
	URL string `json:"url"`
}
