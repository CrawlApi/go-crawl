package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func WEB(c *gin.Context)  {
	a := make([]string, 5)

	c.HTML(http.StatusOK, "api.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
		"array":a,
	})
}

var array_data = ``