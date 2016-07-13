package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func InitRouters(router *gin.RouterGroup) {
	setupFBRouters(router)
	setupIGRouters(router)
	setupWBRouters(router)
	setupWXRouters(router)
	setupUIDRouters(router)
	setupAPIRouters(router)
	setupYTBRouters(router)
}

func setupAPIRouters(router *gin.RouterGroup) {
	router.POST("/api/json/:query", controllers.GetHTMLAPI)
}
