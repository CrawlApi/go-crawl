package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func ProfileFBHealthz(c *gin.Context) {

	r := data.GR().GetFBProfileResult("JustinTimberlake")
	if r.HasError() {
		c.AbortWithError(http.StatusBadRequest, r.GetError())
	}
	var rawData models.FBRawProfile
	err := common.ParseJson(r.Body, &rawData)
	if err != nil {
		c.AbortWithError(http.StatusConflict, err)
	}
	if rawData.ID == "" {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.String(http.StatusOK, "ok")
}

func ProfileIGHealthz(c *gin.Context) {

}

func ProfileWBHealthz(c *gin.Context) {

}

func ProfileWXHealthz(c *gin.Context) {

}

func ProfileYTBHealthz(c *gin.Context) {

}

func PostsFBHealthz(c *gin.Context) {

	r := data.GR().GetFBPostsResult("JustinTimberlake", "5")
	if r.HasError() {
		c.AbortWithError(http.StatusBadRequest, r.GetError())
	}
	var rawData models.FBRawPosts
	err := common.ParseJson(r.Body, &rawData)
	if err != nil {
		c.AbortWithError(http.StatusConflict, err)

	}
	if len(rawData.Data) > 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.String(http.StatusOK, "ok")
}

func PostsIGHealthz(c *gin.Context) {

}

func PostsWXHealthz(c *gin.Context) {

}
func PostsWBHealthz(c *gin.Context) {

}

func PostsYTBHealthz(c *gin.Context) {

}
