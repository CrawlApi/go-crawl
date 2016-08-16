package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupHealthz(router *gin.RouterGroup) {
	profileHealthz := router.Group("/healthz/profile")
	{
		profileHealthz.GET("/fb", controllers.ProfileFBHealthz)
		profileHealthz.GET("/ig", controllers.ProfileIGHealthz)
		profileHealthz.GET("/wx", controllers.ProfileWXHealthz)
		profileHealthz.GET("/wb", controllers.ProfileWBHealthz)
		profileHealthz.GET("/ytb", controllers.ProfileYTBHealthz)
	}

	postsHealthz := router.Group("/healthz/posts")
	{
		postsHealthz.GET("/fb", controllers.PostsFBHealthz)
		postsHealthz.GET("/ig", controllers.PostsIGHealthz)
		postsHealthz.GET("/wx", controllers.PostsWXHealthz)
		postsHealthz.GET("/wb", controllers.PostsWBHealthz)
		postsHealthz.GET("/ytb", controllers.PostsYTBHealthz)
	}

	router.GET("/healthz", controllers.Healthz)
}