package main

import (
	"gopkg.in/redis.v3"
	"github.com/huandu/facebook"
	"github.com/gin-gonic/gin"
	"log"
)

func RedisClient(key string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.95:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	for i := 0; i < 10; i++ {
		err := client.RPush("LIST", `{"type": "facebook", "id": "5718732097", "access_token": "` + key + `", "ids": ["5718732097_10153714754837098", "5718732097_10153714372452098", "5718732097_10153711314922098"]}`).Err()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		RedisClient(key)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and server on 0.0.0.0:8080

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