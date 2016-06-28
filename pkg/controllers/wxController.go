package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

const (
	URL_WECHAT_PROFILE = "http://weixin.sogou.com/weixin?type=1&query=%s&ie=utf8&_sug_=n&_sug_type_="
	URL_WECHAT_POSTS =   "http://weixin.sogou.com/weixin?type=1&query=%s&ie=utf8&_sug_=n&_sug_type_="
)

func GetWXProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WXRepo{
		Agent:  common.GetAgent(),
		Url:    common.UrlString(URL_WECHAT_PROFILE, userId),
	}
	getProfile(c, repo)
}

func GetWXPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WXRepo{
		Agent: common.GetAgent(),
		Url:     common.UrlString(URL_WECHAT_POSTS, userId),
	}
	getPosts(c, repo)
}
