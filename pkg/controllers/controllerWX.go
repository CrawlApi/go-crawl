package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetWXProfile(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewWXRepoWithUid(userId)
	getProfile(c, repo)
}

func GetWXPosts(c *gin.Context) {
	userId := c.Param("userId")
	repo := data.NewWXRepoWithUid(userId)
	getPosts(c, repo)
}
