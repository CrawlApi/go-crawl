package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

type ApiJson struct {
	Domain        string `json:"domain"`
	Url           string            `json:"url"`
	StartSelector string            `json:"start_selector"`
	Datas         map[string]string `json:"datas"`
	NextPage      string            `json:"next_page"`
	Limit         int               `json:"limit"`
}

type ResultJson struct {
	Url    string        `json:"url"`
	Datas  []interface{} `json:"datas"`
	Status bool   `json:"status"`
	Date   int64  `json:"date"`
}

func findDocWithUrl(url, query string) *goquery.Document {
	rawUrl := fmt.Sprintf(url, query)
	doc, _ := goquery.NewDocument(rawUrl)
	return doc
}

func findDocWithSelect(selection *goquery.Selection, selector, domain string) *goquery.Document {
	url := filterValue(selection, selector)
	if !strings.Contains(url, "http") {
		url = domain + url
	}
	doc, _ := goquery.NewDocument(url)
	return doc
}

func findItems(selection *goquery.Selection, selector string, config map[string]string) []interface{} {
	var items []interface{}

	selection.Find(selector).Each(func(i int, s *goquery.Selection) {
		item := make(map[string]string)
		for k, v := range config {
			item[k] = filterValue(s, v)
		}
		items = append(items, item)
	})
	return items
}

func ParseHTMLAPI(api ApiJson, query string, ch chan ResultJson) {

	var wg sync.WaitGroup

	var result ResultJson
	result.Url = fmt.Sprintf(api.Url, query)

	doc := findDocWithUrl(api.Url, query)
	wg.Add(1)
	go func() {
		items := findItems(doc.Selection, api.StartSelector, api.Datas)

		result.Datas = append(result.Datas, items)
		wg.Done()
	}()

	if len(api.NextPage) > 0 {
		for i := 0; i < api.Limit; i++ {
			wg.Add(1)
			go func() {
				docN := findDocWithSelect(doc.Selection, api.NextPage, api.Domain)

				itemsN := findItems(docN.Selection, api.StartSelector, api.Datas)

				result.Datas = append(result.Datas, itemsN)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	ch <- result

}

func filterValue(selection *goquery.Selection, selector string) string {
	if strings.Contains(selector, "@") {
		i := strings.Index(selector, "@")
		result, _ := selection.Find(selector[0:i]).Attr(selector[i + 1: len(selector)])
		return result
	} else {
		return selection.Find(selector).Text()
	}

}