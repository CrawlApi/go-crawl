package routers

import "github.com/gin-gonic/gin"

func InitRouters(router *gin.Engine) {

	setupIGRouters(router)
	setupFBRouters(router)
	setupWBRouters(router)
	setupWXRouters(router)
	setupUIDRouters(router)
}
