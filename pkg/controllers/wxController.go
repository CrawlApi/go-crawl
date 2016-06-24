package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetWXProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WXRepo{
		Agent:  common.GetAgent(),
		Url:    "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_=",
		UserId: userId,
	}
	GetProfile(c, repo)
}

func GetWXPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WXRepo{
		Agent: common.GetAgent(),
		Url:   "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_=",
	}
	GetPosts(c, repo)
}
