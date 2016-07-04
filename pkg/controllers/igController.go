package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
)

func GetIGProfile(c *gin.Context) {
	userId := c.Param("userId")
	v := c.DefaultQuery("version", "v2")

	var repo data.Profile
	switch v {
	case "v2":
		repo = data.NewIGV2RepoWithUid(userId)
	default:
		repo = data.NewIGRepoWithUid(userId)
	}

	getProfile(c, repo)
}

func GetIGPosts(c *gin.Context) {
	userId := c.Param("userId")
	v := c.DefaultQuery("version", "v2")

	var repo data.Posts
	switch v {
	case "v2":
		repo = data.NewIGV2RepoWithUid(userId)
	default:
		repo = data.NewIGRepoWithUid(userId)
	}

	getPosts(c, repo)
}
