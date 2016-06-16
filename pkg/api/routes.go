package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	serverApi := router.Group("/api")
	{
		serverApi.POST("/uid", GetUid)
		serverApi.GET("/:type/profile/:userId", GetProfile)
		serverApi.GET("/:type/posts/:userId", GetPosts)
	}
}
