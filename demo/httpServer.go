package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
)

type ResultOne struct {
	//Enroll_data_array []string `json:"enroll_data_array"`
	User_id string `json:"user_id"`
}
type EnrollData struct {
	Fk_name string         `json:"fk_name"`
	Fk_time string         `json:"fk_time"`
	Fk_info EnrollDataInfo `json:"fk_info"`
}

type EnrollDataInfo struct {
	Face_data_ver         int      `json:"face_data_ver"`
	Firmware              string   `json:"firmware"`
	Firmware_filename     string   `json:"firmware_filename"`
	Fk_bin_data_lib       string   `json:"fk_bin_data_lib"`
	Fp_data_ver           int      `json:"fp_data_ver"`
	Supported_enroll_data []string `json:"supported_enroll_data"`
}

const (
	RECEIVE_CMD          = "receive_cmd"
	REALTIME_GLOG        = "realtime_glog"
	REALTIME_ENROLL_DATA = "realtime_enroll_data"
)

func LogHeader(header map[string][]string) {
	log.Println("Header User-Agent: ", header["User-Agent"])

	log.Println("Header Content-Type: ", header["Content-Type"])
	log.Println("Header Content-Length: ", header["Content-Length"])

	log.Println("Header Accept: ", header["Accept"])
	log.Println("Header Accept-Language: ", header["Accept-Language"])
	log.Println("Header Accept-Encoding: ", header["Accept-Encoding"])

	log.Println("Header Connection: ", header["Connection"])

	log.Println("Header trans_id: ", header["trans_id"])
	log.Println("Header blk_no: ", header["blk_no"])
	log.Println("Header blk_len: ", header["blk_len"])
	log.Println("Header dev_id: ", header["dev_id"])

	log.Println("Header cmd_return_code: ", header["cmd_return_code"])
	log.Println("Header request_code: ", header["request_code"])

	log.Println("Header: ", header)
}

func ParseEnrollData(header map[string][]string, body []byte) {
	//log.Println(string(body[4:282]))
	var enrollData EnrollData
	err := json.Unmarshal(body[4:282], &enrollData)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", enrollData)
}

func ParseRealTimeGlog(header map[string][]string, body []byte) {
	log.Println(string(body))
}

func ParseRealTimeEnrollData(header map[string][]string, body []byte) {
	log.Println(string(body))
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		gorequest.New().Post("http://192.168.30.64:8001").Set("request_code", "[send_cmd_result]").Set("cmd_code", "GET_ENROLL_DATA").Set("trans_id", "ReceiveCommandAction").Send(`{"user_id":"1", "backup_number":"10"}`).End()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/", func(c *gin.Context) {

		header := c.Request.Header
		//LogHeader(header)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		request_code := header["request_code"][0]

		switch request_code {
		case RECEIVE_CMD:
		//ParseEnrollData(header, body)
		case REALTIME_GLOG:
			ParseRealTimeGlog(header, body)
		case REALTIME_ENROLL_DATA:
			//ParseRealTimeEnrollData(header, body)
		}

		//realBody := make([]byte,  500)
		//realBody := body[4:190]
		//
		//var res ResultOne
		//err  = json.Unmarshal(realBody, &res)
		//if err != nil {
		//	panic(err)
		//}
		//
		//log.Printf("Result: %+v", res)
		//
		//log.Println(string(realBody))

		//err = ioutil.WriteFile("output.txt", realBody, 0644)
		//if err != nil {
		//	panic(err)
		//}

		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run(":8001")
}
