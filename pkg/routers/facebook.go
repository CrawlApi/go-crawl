package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func SetupFBRouters(router *gin.Engine) {
	router.GET("/api/fb/profile/:userId", controllers.GetFBProfile)
	router.GET("/api/fb/posts/:userId", controllers.GetFBPosts)
}