package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

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

func ParseHTMLAPI(api ApiJson, query string) ResultJson {
	var result ResultJson
	url := fmt.Sprintf(api.Url, query)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return result
	}

	result.Url = fmt.Sprintf(api.Url, query)

	doc.Find(api.StartSelector).Each(func(i int, s *goquery.Selection) {

		item := make(map[string]string)
		for k, v := range api.Datas {
			item[k] = filterValue(v, s)
		}
		result.Datas = append(result.Datas, item)
	})

	if len(api.NextPage) > 0 {
		for i := 0; i < api.Limit; i++ {
			url := filterValue(api.NextPage, doc.Selection)
			if !strings.Contains(url, "http") {
				url = api.Domain + url
			}
			doc, err = goquery.NewDocument(url)
			if err != nil {
				return result
			}

			doc.Find(api.StartSelector).Each(func(i int, s *goquery.Selection) {

				item := make(map[string]string)
				for k, v := range api.Datas {
					item[k] = filterValue(v, s)
				}
				result.Datas = append(result.Datas, item)
			})
		}
	}
	return result
}

func filterValue(value string, s *goquery.Selection) string {
	if strings.Contains(value, "@") {
		i := strings.Index(value, "@")
		result, _ := s.Find(value[0:i]).Attr(value[i + 1: len(value)])
		return result
	} else {
		return s.Find(value).Text()
	}

}