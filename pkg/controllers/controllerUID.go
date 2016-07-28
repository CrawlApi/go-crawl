package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
	"net/http"
)

func GetUid(c *gin.Context) {

	var api models.APIJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "post json data error",
		})
		return
	}

	var repo data.UID
	switch parser.CheckUrl(api.Url) {
	case "facebook":
		repo = data.NewFBRepoWithUrl(api.Url)
	case "instagram":
		repo = data.NewIGV2RepoWithUrl(api.Url)
	case "weixin":
		repo = data.NewWXRepoWithUrl(api.Url)
	case "weibo":
		repo = data.NewWBRepoWithUrl(api.Url)
	case "youtube":
		repo = data.NewYTBRepoWithUrl(api.Url)
	}

	getRealUid(c, repo)
}
