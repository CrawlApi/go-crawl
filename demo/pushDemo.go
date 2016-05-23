package main

import (
	"github.com/huandu/facebook"
	"gopkg.in/redis.v3"
	"log"
	"time"
)

func RedisClient(key string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.95:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for true {
		err := client.RPush("LIST", `{"type": "facebook", "url": "https://www.facebook.com/justintimberlake/"}`).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	//r := gin.Default()
	//r.GET("/key/:key", func(c *gin.Context) {
	//	key := c.Param("key")
	//	RedisClient(key)
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":8080") // listen and server on 0.0.0.0:8080
	RedisClient("EAACEdEose0cBADZABlVMw8E5E9uxbjeZALGs3dMRUkZAWi3ZBxkW5semxefQWttu56a225duvYC0a5LZACEISgHc6fgL29tMCk7bTjemdHBhUuQkBDnBZAFwpZB1ZAZBcX7nApRgK1xLEHJkmBwL1tIpDBsj3wifnMfA3JdbTFOZBa5McnlGKbH4l0")

}

func FaceBookTest() {
	facebook.Version = "v2.6"

	res, _ := facebook.Get("/5718732097_10153714372452098", facebook.Params{
		"fields":       "message,picture,full_picture,shares,updated_time,created_time,name",
		"access_token": "token",
	})

	likes, _ := facebook.Get("/5718732097_10153714372452098/likes?summary=true", facebook.Params{
		"fields":       "message,picture,full_picture,shares,updated_time,created_time,name",
		"access_token": "token",
	})

	log.Printf("%+v", res)
	log.Printf("%+v", likes)
}
