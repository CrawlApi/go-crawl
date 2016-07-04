package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"net/http"
)

func GetHTMLAPI(c *gin.Context) {
	query := c.Param("query")

	var api data.ApiJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "json data error",
		})
		return
	}
	result := data.ParseHTMLAPI(api, query)

	c.JSON(http.StatusOK, gin.H{
		"api": result,
	})
}

