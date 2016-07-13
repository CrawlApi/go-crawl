package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupFBRouters(router *gin.RouterGroup) {
	router.GET("/api/fb/profile/:userId", controllers.GetFBProfile)
	router.GET("/api/fb/posts/:userId", controllers.GetFBPosts)
	router.GET("/api/fb/post/:postId/reactions", controllers.GetFBPostReactions)

	router.POST("/api/fb/post", controllers.GetFBPostInfo)

}
