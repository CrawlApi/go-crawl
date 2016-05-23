package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v3"
	"time"
)

type Client struct {
	db    *sql.DB
	redis *redis.Client
}

func New() (*Client, error) {

	db, err := sql.Open("mysql", "root:summer@tcp(127.0.0.1:3306)/cirkol")
	if err != nil {
		return nil, err
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.95:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client := Client{
		db:    db,
		redis: redis,
	}
	return &client, nil
}

func (client *Client) Close() error {

	if client.db != nil {
		return client.db.Close()
	}
	return nil
}

func (client *Client) StartRedis() ([]string, error) {
	return client.redis.BLPop(1*time.Second, "LIST").Result()
}

func (client *Client) PushData(key string, data string) {
	client.redis.RPush(key, data).Err()
}
