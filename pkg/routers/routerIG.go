package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupIGRouters(router *gin.RouterGroup) {
	router.GET("/api/ig/profile/:userId", controllers.GetIGProfile)
	router.GET("/api/ig/posts/:userId", controllers.GetIGPosts)

	router.POST("/api/ig/post", controllers.GetIGPostInfo)

}
