package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"net/http"
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

func GetWBPostInfo(c *gin.Context) {
	url, err := getUrlFromJson(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}
	repo := data.NewWBRepoWithUrl(url)
	getPostInfo(c, repo)
}
