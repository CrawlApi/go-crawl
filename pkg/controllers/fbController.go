package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetFBProfile(c *gin.Context) {
	userId := c.Param("userId")

	repo := &data.FBRepo{
		Agent: common.GetAgent(),
		Url:   "https://graph.facebook.com/v2.6/" + userId + "?fields=" + common.PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + common.GetFBToken(),
	}
	GetProfile(c, repo)
}

func GetFBPosts(c *gin.Context) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	repo := &data.FBRepo{
		Agent: common.GetAgent(),
		Url:   "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + common.PAGE_FEED_FIELDS_ENABLE + "," + common.PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + common.GetFBToken(),
	}
	GetPosts(c, repo)
}
