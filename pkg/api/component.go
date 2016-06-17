package api

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"github.com/gin-gonic/gin"
)

var (
	logCh chan interface{}
	reqClient  *gorequest.SuperAgent
)

func SetupComponent(router *gin.Engine) {
	setupRequestClient()
	setupLogger()
}

func setupRequestClient() {

	reqClient = gorequest.New()
}

func setupLogger() {
	logCh = make(chan interface{}, 10)
	go logging()
}

func logging() {
	for {
		log.Println(<-logCh)
	}
}