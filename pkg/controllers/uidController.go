package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

const (
	REGEX_URL_TYPE = `(facebook|instagram|weixin|weibo)`
)

func GetUid(c *gin.Context) {

	var api models.APIJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "post json data error",
		})
		return
	}

	var repo data.UID
	switch checkUrl(api.Url) {
	case "facebook":
		repo = data.NewFBRepoWithUrl(api.Url)
	case "instagram":
		repo = data.NewIGV2RepoWithUrl(api.Url)
	case "weixin":
		repo = data.NewWXRepoWithUrl(api.Url)
	case "weibo":
		repo = data.NewWBRepoWithUrl(api.Url)
	}

	getRealUid(c, repo)
}

func checkUrl(url string) string {
	matcher := common.Matcher(REGEX_URL_TYPE, url)
	if len(matcher) > 0 {
		return matcher[0]
	}
	return ""
}
