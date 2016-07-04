package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"net/http"
	"github.com/llitfkitfk/cirkol/pkg/common"
)

func GetHTMLAPI(c *gin.Context) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	query := c.Param("query")

	var api data.ApiJson
	err := c.BindJSON(&api)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "json data error",
		})
		return
	}
	rCh := make(chan data.ResultJson, 1)
	go data.ParseHTMLAPI(api, query, rCh)

	var result data.ResultJson
	select {
	case result = <-rCh:
	case <-timer:
		result = TimeOutResult()
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func TimeOutResult() data.ResultJson {
	var r data.ResultJson
	r.Date = common.Now()
	r.Status = false
	return r
}
