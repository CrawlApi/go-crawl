package routers

import "github.com/gin-gonic/gin"

func InitRouters(router *gin.Engine) {

	SetupIGRouters(router)
	SetupFBRouters(router)
	SetupWBRouters(router)
	SetupWXRouters(router)
}
