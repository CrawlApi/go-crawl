package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

func GetIGProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := &data.IGRepo{
		Agent:  common.GetAgent(),
		Url:    common.UrlString(URL_WECHAT_PROFILE, userId),
	}
	GetProfile(c, repo)
}

func GetIGPosts(c *gin.Context) {

}
