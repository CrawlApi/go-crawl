package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Server  string `json:"server"`
	FBToken string `json:"fb_token"`
}

var AppConfig Configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {

	data, err := ioutil.ReadFile("pkg/common/config.json")
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}

	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

func GetFBToken() string {
	return AppConfig.FBToken
}
