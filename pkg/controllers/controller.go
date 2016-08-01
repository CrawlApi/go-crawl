package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

func getRealUid(c *gin.Context, repo data.UID) {
	if repo == nil {
		c.JSON(http.StatusOK, gin.H{
			"uid": data.FetchUIDErr(common.EmptyRepoError()),
		})
		return
	}

	timeout := c.DefaultQuery("timeout", "10")
	timer := common.Timeout(timeout)
	pCh := make(chan models.UID)

	go func() {
		defer close(pCh)
		pCh <- repo.ParseRawUID(repo.FetchUIDApi())
	}()

	var uid models.UID
	select {
	case uid = <-pCh:
	case <-timer:
		uid = data.FetchUIDErr(common.TimeOutError())
	}
	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}

func getProfile(c *gin.Context, repo data.Profile) {
	timeout := c.DefaultQuery("timeout", "10")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Profile)

	go func() {
		defer close(pCh)
		pCh <- repo.ParseRawProfile(repo.FetchProfileApi())
	}()

	var profile models.Profile
	select {
	case profile = <-pCh:
	case <-timer:
		profile = data.FetchProfileErr(common.TimeOutError())
	}
	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}

func getPosts(c *gin.Context, repo data.Posts) {
	timeout := c.DefaultQuery("timeout", "10")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Posts)

	go func() {
		defer close(pCh)
		pCh <- repo.ParseRawPosts(repo.FetchPostsApi())
	}()

	var posts models.Posts
	select {
	case posts = <-pCh:
	case <-timer:
		posts = data.FetchPostsErr(common.TimeOutError())
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func getPostInfo(c *gin.Context, repo data.Post) {
	timeout := c.DefaultQuery("timeout", "10")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Post)

	go func() {
		defer close(pCh)
		pCh <- repo.ParsePostInfo(repo.FetchPostInfo())
	}()

	var post models.Post
	select {
	case post = <-pCh:
	case <-timer:
		post = data.FetchPostErr(common.TimeOutError())
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func getUrlFromJson(c *gin.Context) (string, error) {
	var api models.APIJson
	err := c.BindJSON(&api)
	if err != nil {
		return "", err
	}
	return api.Url, nil
}
