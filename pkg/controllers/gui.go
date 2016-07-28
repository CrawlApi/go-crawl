package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func WEB(c *gin.Context)  {
	c.HTML(http.StatusOK, "api.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
	})
}
