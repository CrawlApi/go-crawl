package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

const (
	URL_WEIBO_PROFILE = "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=%s"
	URL_WEIBO_POSTS = "http://m.weibo.cn/%s"
)

func GetWBProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WBRepo{
		Agent:  common.GetAgent(),
		Url:   common.UrlString(URL_WEIBO_PROFILE, userId),
	}
	GetProfile(c, repo)
}

func GetWBPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WBRepo{
		Agent: common.GetAgent(),
		Url:    common.UrlString(URL_WEIBO_POSTS, userId),
	}
	GetPosts(c, repo)
}
