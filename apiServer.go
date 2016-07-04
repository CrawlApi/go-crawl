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
	routers.InitRouters(&engine.RouterGroup)

	addr := common.AppConfig.Server

	log.Println("Start Server...")
	log.Println("Server Version: ", common.AppConfig.Version)

	engine.Run(addr)
}
