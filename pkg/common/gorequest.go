package common

import "github.com/parnurzeal/gorequest"

var reqAgent *gorequest.SuperAgent

func initRequestAgent() {
	reqAgent = gorequest.New()
}

func GetAgent() *gorequest.SuperAgent {
	return reqAgent
}
