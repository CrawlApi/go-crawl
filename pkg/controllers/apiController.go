package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)

func GetHTMLAPI(c *gin.Context) {
	query := c.Param("query")

	var api data.ApiJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "json data error",
		})
		return
	}

	url := fmt.Sprintf(api.Url, query)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "url fetch error",
		})
		return
	}

	var result data.ResultJson

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
				c.JSON(http.StatusOK, gin.H{
					"message": "url fetch error",
					"api":result,
				})
				return
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

	c.JSON(http.StatusOK, gin.H{
		"api": result,
	})
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
