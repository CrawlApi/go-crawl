package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupUIDRouters(router *gin.RouterGroup) {
	router.POST("/api/uid", controllers.GetUid)
}
