package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

func GetFBProfile(c *gin.Context) {

	userId := c.Param("userId")
	repo := data.NewFBRepoWithUid(userId)
	getProfile(c, repo)
}

func GetFBPosts(c *gin.Context) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	repo := data.NewFBRepoWithLimit(userId, limit)
	getPosts(c, repo)
}

func GetFBPostReactions(c *gin.Context) {
	postId := c.Param("postId")
	repo := data.NewFBRepoWithPid(postId)
	getReactions(c, repo)
}

func GetFBPostInfo(c *gin.Context) {
	url, err := getUrlFromJson(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	repo := data.NewFBRepoWithUrl(url)
	getPostInfo(c, repo)
}

func getReactions(c *gin.Context, repo *data.FBRepo) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.FBReactions)

	go func() {
		defer close(pCh)
		pCh <- repo.ParseRawReactions(repo.FetchReactionsApi())
	}()

	var reactions models.FBReactions
	select {
	case reactions = <-pCh:
	case <-timer:
		reactions = data.FetchReactionsError(common.TimeOutError())
	}
	c.JSON(http.StatusOK, gin.H{
		"like_counts": reactions,
	})
}
