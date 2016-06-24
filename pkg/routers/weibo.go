package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func SetupWBRouters(router *gin.Engine) {
	router.GET("/api/wb/profile/:userId", controllers.GetWBProfile)
	router.GET("/api/wb/posts/:userId", controllers.GetWBPosts)
}
