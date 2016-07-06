package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetWBProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewWBRepoWithUid(userId)
	getProfile(c, repo)
}

func GetWBPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewWBRepoWithUid(userId)
	getPosts(c, repo)
}
