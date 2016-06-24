package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

func GetWBProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WBRepo{
		Agent:  common.GetAgent(),
		Url:   "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId,
	}
	GetProfile(c, repo)
}

func GetWBPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.WBRepo{
		Agent: common.GetAgent(),
		Url:   "http://m.weibo.cn/d/" + userId,
	}
	GetPosts(c, repo)
}
