package api

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
)

var (
	DBClient *client.Client
)

func StartService() {

	for true {
		DBClient.StartService()
	}
}

