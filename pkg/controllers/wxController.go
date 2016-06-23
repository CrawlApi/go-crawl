package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"log"
)

func GetWXProfile(c *gin.Context) {
	timer := common.Timer(c)
	pCh := make(chan Profile)

	go fetchWXProfile(c, pCh)

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

func GetWXPosts(c *gin.Context) {
	timer := common.Timer(c)
	pCh := make(chan Posts)

	go fetchWXPosts(c, pCh)

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

func fetchWXProfile(c *gin.Context, ch chan Profile) {
	var profile Profile

	repo := data.WXRepo{
		Agent:common.GetAgent(),
	}
	rawData, err := repo.GetRawProfile(c)
	if err != nil {
		profile.FetchErr(err)
		ch <- profile
		return
	}
	//profile.ParseRawProfile(rawData)
	log.Println(rawData)
	ch <- profile
}

func fetchWXPosts(c *gin.Context, ch chan Posts) {
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