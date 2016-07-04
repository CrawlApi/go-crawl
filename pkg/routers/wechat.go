package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupWXRouters(router *gin.RouterGroup) {
	router.GET("/api/wx/profile/:userId", controllers.GetWXProfile)
	router.GET("/api/wx/posts/:userId", controllers.GetWXPosts)
}
