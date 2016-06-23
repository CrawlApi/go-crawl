package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"net/http"
)

func GetFBProfile(c *gin.Context) {
	timer := common.Timer(c)
	pCh := make(chan Profile)

	go fetchFBProfile(c, pCh)

	select {
	case profile := <-pCh:
		c.JSON(http.StatusOK, gin.H{
			"profile": profile,
		})
	case <-timer:
		c.JSON(http.StatusOK, gin.H{
			"profile": TimeOutProfile(),
		})
	}
}

func GetFBPosts(c *gin.Context) {
	timer := common.Timer(c)
	pCh := make(chan Posts)

	go fetchFBPosts(c, pCh)

	select {
	case posts := <-pCh:
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	case <-timer:
		c.JSON(http.StatusOK, gin.H{
			"posts": TimeOutPosts(),
		})
	}
}

func fetchFBProfile(c *gin.Context, ch chan Profile) {
	var profile Profile

	repo := data.FBRepo{
		Agent:common.GetAgent(),
	}
	rawData, err := repo.GetRawProfile(c)
	if err != nil {
		profile.FetchErr(err)
		ch <- profile
		return
	}
	profile.ParseRawProfile(rawData)
	ch <- profile
}

func fetchFBPosts(c *gin.Context, ch chan Posts) {
	var posts Posts
	repo := data.FBRepo{
		Agent:common.GetAgent(),
	}
	rawData, err := repo.GetRawPosts(c)
	if err != nil {
		posts.FetchErr(err)
		ch <- posts
		return
	}
	posts.ParseRawPosts(rawData)
	ch <- posts
}