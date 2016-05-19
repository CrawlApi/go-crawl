package main

import (
	"fmt"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"os"
	"github.com/llitfkitfk/cirkol/pkg/cmd"
)

func startService() {
	fmt.Println("Starting Service...")
	cmd.StartService()
}

func exitWithMessage(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func initClient() {
	cl, err := client.New()
	if err != nil {
		exitWithMessage(err.Error())
	}
	cmd.DBClient = cl
}

func main() {

	initClient()

	if cmd.DBClient != nil {
		defer cmd.DBClient.Close()
	}

	startService()
}
