package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

const (
	URL_INSTAGRAM_PROFILE = "https://www.instagram.com/%s/"
	URL_INSTAGRAM_PROFILE_V2 = "https://i.instagram.com/api/v1/users/%s/info/"
	URL_INSTAGRAM_POSTS = "https://www.instagram.com/%s/media/"
	URL_INSTAGRAM_POSTS_V2 = "https://i.instagram.com/api/v1/users/%s/info/"

)

func GetIGProfile(c *gin.Context) {
	userId := c.Param("userId")
	v := c.DefaultQuery("version", "v2")

	var repo data.Profile
	switch v {
	case "v2":
		repo = &data.IGV2Repo{
			Agent:  common.GetAgent(),
			Url:    common.UrlString(URL_INSTAGRAM_PROFILE_V2, userId),

		}
	default:
		repo = &data.IGRepo{
			Agent:  common.GetAgent(),
			Url:    common.UrlString(URL_INSTAGRAM_PROFILE, userId),

		}
	}

	getProfile(c, repo)
}

func GetIGPosts(c *gin.Context) {
	userId := c.Param("userId")
	v := c.DefaultQuery("version", "v2")

	var repo data.Posts
	switch v {
	case "v2":
		repo = &data.IGV2Repo{
			Agent:  common.GetAgent(),
			Url:    common.UrlString(URL_INSTAGRAM_POSTS_V2, userId),

		}
	default:
		repo = &data.IGRepo{
			Agent:  common.GetAgent(),
			Url:    common.UrlString(URL_INSTAGRAM_POSTS, userId),

		}
	}

	getPosts(c, repo)
}
