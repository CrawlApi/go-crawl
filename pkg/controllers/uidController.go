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

	var repo data.Repo
	switch checkUrl(rawurl) {
	case "facebook":
		repo = &data.FBRepo{
			Agent:  common.GetAgent(),
			Url:   rawurl,
		}
	case "instagram":
		repo = &data.IGV2Repo{
			Agent:  common.GetAgent(),
			Url:   rawurl,
		}
	case "weixin":
		repo = &data.WXRepo{
			Agent:  common.GetAgent(),
			Url:   rawurl,
		}
	case "weibo":
		repo = &data.WBRepo{
			Agent:  common.GetAgent(),
			Url:   rawurl,
		}
	}

	GetRealUid(c, repo)
}

func checkUrl(url string) string {
	matcher := common.Matcher(REGEX_URL_TYPE, url)
	if len(matcher) > 0 {
		return matcher[0]
	}
	return ""
}
