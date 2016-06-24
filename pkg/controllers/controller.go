package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
	"time"
)

func GetProfile(c *gin.Context, repo data.Repo) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Profile)

	go fetchProfile(repo, pCh)

	var profile models.Profile
	select {
	case profile = <-pCh:
	case <-timer:
		profile = TimeOutProfile()
	}
	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}

func fetchProfile(repo data.Repo, ch chan models.Profile) {
	var profile models.Profile

	body, err := repo.FetchApi()
	if err != nil {
		profile.FetchErr(err)
		ch <- profile
		return
	}
	profile = repo.ParseRawProfile(body)

	ch <- profile
}

func GetPosts(c *gin.Context, repo data.Repo) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Posts)

	go fetchPosts(repo, pCh)

	var posts models.Posts
	select {
	case posts = <-pCh:
	case <-timer:
		posts = TimeOutPosts()
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func fetchPosts(repo data.Repo, ch chan models.Posts) {
	var posts models.Posts

	body, err := repo.FetchApi()
	if err != nil {
		posts.FetchErr(err)
		ch <- posts
		return
	}
	posts = repo.ParseRawPosts(body)

	ch <- posts
}

func TimeOutProfile() models.Profile {
	var p models.Profile
	p.ErrCode = common.ERROR_CODE_API_TIMEOUT
	p.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	p.Date = time.Now().Unix()
	p.Status = false

	return p
}

func TimeOutPosts() models.Posts {
	var p models.Posts
	p.ErrCode = common.ERROR_CODE_API_TIMEOUT
	p.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	p.Date = time.Now().Unix()
	p.Status = false
	return p
}
