package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
	"log"
	"github.com/llitfkitfk/cirkol/pkg/common"
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
	log.Println("select url: ", url)

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
		for _, item := range items {
			result.Datas = append(result.Datas, item)
		}
		wg.Done()
	}()

	if len(api.NextPage) > 0 {
		var selection *goquery.Selection
		selection = doc.Selection
		for i := 1; i < api.Limit; i++ {
			docN := findDocWithSelect(selection, api.NextPage, api.Domain)
			wg.Add(1)
			go func() {
				itemsN := findItems(docN.Selection, api.StartSelector, api.Datas)
				for _, item := range itemsN {
					result.Datas = append(result.Datas, item)
				}
				wg.Done()
			}()
			selection = docN.Selection
		}
	}

	wg.Wait()

	ch <- result
}

type Config struct {
	Selector      string
	Attr          string
	Limit         string
	First         bool
	Last          bool
	ChildSelector string
}

func parseSelector(src string) (string, string) {
	//if strings.Contains(src, " ") {
	//	selectorArr := strings.Split(src, " ")
	//	return selectorArr[0], selectorArr[1]
	//}
	return src, ""
}

func parseConfig(selector string) Config {
	var config Config
	if strings.Contains(selector, "@") {
		config.Attr = selector[strings.Index(selector, "@") + 1: len(selector)]
		selector = selector[0:strings.Index(selector, "@")]
	}

	if strings.Contains(selector, "|") {
		configArr := strings.Split(selector, "|")
		config.Selector, config.ChildSelector = parseSelector(configArr[0])
		config.Limit = configArr[1]
		for _, value := range configArr {
			if strings.EqualFold(value, "FIRST") {
				config.First = true
			}
			if strings.EqualFold(value, "LAST") {
				config.Last = true
			}
		}
	} else {
		config.Selector, config.ChildSelector = parseSelector(selector)
	}
	return config
}

func filterValue(selection *goquery.Selection, selector string) string {
	config := parseConfig(selector)

	s := selection.Find(config.Selector)
	var res string
	if len(config.ChildSelector) > 0 {
		log.Println(s.Find(config.ChildSelector).Nodes)
		s.Find(config.ChildSelector)
	}

	if config.First {
		s = s.First()
	}

	if config.Last {
		s = s.Last()
	}

	if len(config.Attr) > 0 {
		res, _ = s.Attr(config.Attr)
	} else {
		res = s.Text()
	}

	if len(config.Limit) > 0 {
		splitor := strings.Split(config.Limit, "-")
		if len(splitor[1]) == 0 {
			res = res[common.Str2Int(splitor[0]):]
		} else {
			res = res[common.Str2Int(splitor[0]): common.Str2Int(splitor[1])]
		}
	}

	return res
}