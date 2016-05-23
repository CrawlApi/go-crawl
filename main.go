package main

import (
	"fmt"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/cmd"
	"os"
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
	cmd.Client = cl
	cmd.TokenCh = make(chan string)
	cmd.CommCh = make(chan string)
}

func main() {

	initClient()

	if cmd.Client != nil {
		defer cmd.Client.Close()
	}

	startService()
}
