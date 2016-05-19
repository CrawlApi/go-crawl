package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

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


func RedisClient() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	err = client.RPush("LIST", "value").Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	RedisClient()
	//const APP_TOKEN = "EAACEdEose0cBAN2vcb22ODdD8i0oyNRthLvVrvv7iv5aKSIxbCknnMhz4p5gLaEEQvOQSdRShEDXYFpjJOf4wi7XYelB3WDFX2No5Gi8JMs1jpB1vF9exgYh7gSZBcZCbZC7PcsGf7ZBQenRseZBWQizk6VbjBWhWIlPI9VhnX1JZAP2b9yqUp"
	//facebook.Version = "v2.6"
	//
	//r := gin.Default()
	//
	//r.GET("/fb/:userId/info", func(c *gin.Context) {
	//	userId := "/" + c.Param("userId")
	//	res, err := facebook.Get(userId, facebook.Params{
	//		"fields":       "name,is_verified,fan_count,posts{created_time,shares,message,full_picture,picture}",
	//		"access_token": APP_TOKEN,
	//	})
	//
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": err.Error(),
	//			"status":  false,
	//			"date":    time.Now().Unix(),
	//		})
	//	} else {
	//		var result FaceBook
	//
	//		result.Id = res["id"]
	//		result.Fan_count = res["fan_count"]
	//		result.Name = res["name"]
	//		result.Is_verified = res["is_verified"]
	//		result.Posts = res["posts"]
	//		//result.Id = "teste"
	//		//result.Fan_count= "teste"
	//		//result.Name= "teste"
	//		//result.Is_verified= "teste"
	//		c.JSON(http.StatusOK, gin.H{
	//			"data":    result,
	//			"message": "OK",
	//			"status":  true,
	//			"date":    time.Now().Unix(),
	//		})
	//	}
	//})
	//
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//
	//r.Run(":8080")
}
