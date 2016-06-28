package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func InitRouters(router *gin.Engine) {

	setupIGRouters(router)
	setupFBRouters(router)
	setupWBRouters(router)
	setupWXRouters(router)
	setupUIDRouters(router)
	//setupAPIRouters(router)
}

func setupAPIRouters(router *gin.Engine) {
	router.GET("/api/help", controllers.GetAPIHelper)
}
