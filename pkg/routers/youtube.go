package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupYTBRouters(router *gin.RouterGroup) {
	router.GET("/api/ytb/profile/:userId", controllers.GetYTBProfile)
	router.GET("/api/ytb/posts/:userId", controllers.GetYTBPosts)
}