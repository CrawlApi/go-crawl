package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/routers"
	"log"
)

func main() {

	common.StartUp()

	engine := gin.Default()
	routers.InitRouters(engine)

	addr := common.AppConfig.Server
	log.Println("Start Server...")
	engine.Run(addr)
}
