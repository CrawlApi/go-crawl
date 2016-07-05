package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetYTBProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewYTBRepoWithUid(userId)
	getProfile(c, repo)
}

func GetYTBPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewYTBRepoWithUid(userId)
	getPosts(c, repo)
}
