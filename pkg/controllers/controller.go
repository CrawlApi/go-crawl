package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

func getRealUid(c *gin.Context, repo data.UID) {
	var uid models.UID
	if repo == nil {
		uid.FetchErr(nil)
		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
		return
	}

	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.UID)

	go fetchUID(repo, pCh)

	select {
	case uid = <-pCh:
	case <-timer:
		uid = TimeOutUID()
	}
	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}

func fetchUID(repo data.UID, ch chan models.UID) {
	defer close(ch)
	var uid models.UID

	body, err := repo.FetchUIDApi()
	if err != nil {
		uid.FetchErr(err)
		ch <- uid
		return
	}
	uid = repo.ParseRawUID(body)

	ch <- uid
}

func getProfile(c *gin.Context, repo data.Profile) {
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

func fetchProfile(repo data.Profile, ch chan models.Profile) {
	defer close(ch)
	var profile models.Profile

	body, err := repo.FetchProfileApi()
	if err != nil {
		profile.FetchErr(err)
		ch <- profile
		return
	}
	profile = repo.ParseRawProfile(body)

	ch <- profile
}

func getPosts(c *gin.Context, repo data.Posts) {
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

func fetchPosts(repo data.Posts, ch chan models.Posts) {
	defer close(ch)

	var posts models.Posts

	body, err := repo.FetchPostsApi()
	if err != nil {
		posts.FetchErr(err)
		ch <- posts
		return
	}
	posts = repo.ParseRawPosts(body)

	ch <- posts
}

func getPostInfo(c *gin.Context, repo data.Post) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.Post)

	go fetchPostInfo(repo, pCh)

	var post models.Post
	select {
	case post = <-pCh:
	case <-timer:
		post = TimeOutPost()
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func fetchPostInfo(repo data.Post, ch chan models.Post) {
	defer close(ch)

	var post models.Post

	body, err := repo.FetchPostInfo()
	if err != nil {
		post.FetchErr(err)
		ch <- post
		return
	}
	post = repo.ParsePostInfo(body)

	ch <- post
}

func TimeOutProfile() models.Profile {
	var p models.Profile
	p.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	p.Date = common.Now()
	p.Status = false

	return p
}

func TimeOutPosts() models.Posts {
	var p models.Posts
	p.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	p.Date = common.Now()
	p.Status = false
	return p
}

func TimeOutPost() models.Post {
	var p models.Post
	p.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	p.Date = common.Now()
	p.Status = false
	return p
}

func TimeOutUID() models.UID {
	var u models.UID
	u.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	u.Date = common.Now()
	u.Status = false
	return u
}
