package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	logCh = make(chan interface{}, 10)
	go Logging()
	apiFB := router.Group("/api/fb")
	{
		apiFB.GET("/uid", GetFBUid)
		apiFB.GET("/profile/:userId", GetFBProfile)
		apiFB.GET("/posts/:userId", GetFBPosts)
	}

	apiIG := router.Group("/api/ig")
	{
		apiIG.GET("/uid", GetIGUid)
		apiIG.GET("/profile/:userId", GetIGProfile)
		apiIG.GET("/posts/:userId", GetIGPosts)
	}

	apiWB := router.Group("/api/wb")
	{
		apiWB.GET("/uid", GetWBUid)
		apiWB.GET("/profile/:userId", GetWBProfile)
		apiWB.GET("/posts/:userId", GetWBPosts)
	}

}
