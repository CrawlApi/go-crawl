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
	var api models.APIJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"message": "json data error",
		})
		return
	}
	//getPostInfo()
}


func getReactions(c *gin.Context, repo *data.FBRepo) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.FBReactions)

	go fetchReactions(repo, pCh)

	var reactions models.FBReactions
	select {
	case reactions = <-pCh:
	case <-timer:
		reactions = TimeOutReactions()
	}
	c.JSON(http.StatusOK, gin.H{
		"like_counts": reactions,
	})
}

func fetchReactions(repo *data.FBRepo, ch chan models.FBReactions) {
	defer close(ch)
	var reactions models.FBReactions

	body, err := repo.FetchReactionsApi()
	if err != nil {
		reactions.FetchErr(err)
		ch <- reactions
		return
	}
	reactions = repo.ParseRawReactions(body)

	ch <- reactions
}

func TimeOutReactions() models.FBReactions {
	var r models.FBReactions
	r.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	r.Date = common.Now()
	r.Status = false
	return r
}
