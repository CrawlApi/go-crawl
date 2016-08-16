package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/routers"
	"runtime"
	"github.com/llitfkitfk/cirkol/pkg/pprof"
)

func main() {

	common.StartUp()
	initClient()
	setMaxProcs()
	engine := gin.Default()
	pprof.Wrap(engine)
	if !common.AppConfig.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	engine.LoadHTMLGlob("pkg/resources/*.templ.html")

	routers.InitRouters(&engine.RouterGroup)

	addr := common.AppConfig.Server

	common.Log.Info("Start Server...")
	common.Log.Info("Server Version: ", common.AppConfig.Version)

	engine.Run(addr)
}

func setMaxProcs() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	common.Log.Info("Running with ", nuCPU, " CPUs")
}

func initClient() {
	cl := client.New()

	data.Agent = cl
}
