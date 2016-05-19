package main

import (
	"encoding/json"
	"fmt"
	"github.com/huandu/facebook"
	"github.com/llitfkitfk/cirkol/pkg/api"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"os"
	"os/signal"
	"time"
)

var APP_TOKEN = "EAACEdEose0cBAJ2OWfFcsVQ08safn6OFJr7uHBd3L20qcELq8oLV89ZC4LiqBpdRDafZAP4IyZBNDZBMtPSH3CDIdETrLzAvJBa5fZBMRcf3KafHYFO6Bj0Vvon33R8DyWmU9TSnZBPbJTCBSu2rOT1DR2HnSeP7rBUq0S7hqLeNBsHgd12ZCKz"

type FaceBook struct {
	Name        interface{} `json:"name"`
	Posts       interface{} `json:"posts"`
	Fan_count   interface{} `json:"fan_count"`
	Is_verified interface{} `json:"is_verified"`
	Id          interface{} `json:"id"`
	Message     interface{} `json:"message"`
	Status      interface{} `json:"status"`
	Date        interface{} `json:"date"`
}

func init() {
	// config
	facebook.Version = "v2.6"
}

func GetInfoByUserId(userId string) {
	res, err := facebook.Get(userId, facebook.Params{
		"fields":       "name,is_verified,fan_count,posts{created_time,shares,message,full_picture,picture}",
		"access_token": APP_TOKEN,
	})
	var result FaceBook

	if err != nil {
		result.Message = err.Error()
		result.Status = false
		result.Date = time.Now().Unix()
	} else {
		result.Message = "OK"
		result.Status = true
		result.Date = time.Now().Unix()
		result.Id = res["id"]
		result.Fan_count = res["fan_count"]
		result.Name = res["name"]
		result.Is_verified = res["is_verified"]
		result.Posts = res["posts"]
	}

	out, _ := json.Marshal(result)

	fmt.Println(string(out))
}
func task() {
	fmt.Println("I am runnning task.")

}

func startService() {
	fmt.Println("Starting Service...")
	api.StartService()
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
	api.DBClient = cl
}

func main() {

	initClient()

	if api.DBClient != nil {
		defer api.DBClient.Close()
	}

	startService()

	//handleSignals()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
