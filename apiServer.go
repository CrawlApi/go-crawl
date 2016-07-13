package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/routers"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/client"
)

func main() {

	common.StartUp()
	initClient()
	engine := gin.Default()

	if !common.AppConfig.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	routers.InitRouters(&engine.RouterGroup)

	addr := common.AppConfig.Server

	common.Info("Start Server...")
	common.Info("Server Version: ", common.AppConfig.Version)

	engine.Run(addr)
}

func initClient() {
	cl := client.New()

	data.Agent = cl
}