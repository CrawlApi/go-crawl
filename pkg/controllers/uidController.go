package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

const (
	REGEX_URL_TYPE = `(facebook|instagram|weixin|weibo)`
)

func GetUid(c *gin.Context) {
	rawurl := c.PostForm("url")

	var repo data.UID
	switch checkUrl(rawurl) {
	case "facebook":
		repo = data.NewFBRepoWithUrl(rawurl)
	case "instagram":
		repo = data.NewIGV2RepoWithUrl(rawurl)
	case "weixin":
		repo = data.NewWXRepoWithUrl(rawurl)
	case "weibo":
		repo = data.NewWBRepoWithUrl(rawurl)
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
