package client

import (
	"database/sql"
	"gopkg.in/redis.v3"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"errors"
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
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})

	_, err = redis.Ping().Result()

	if err != nil {
		return nil, err
	}

	client := Client{
		db: db,
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

func (client *Client) StartService() ([]string, error){

	if client.redis != nil {
		data, err := client.redis.BLPop(1 * time.Second, "LIST").Result()
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("Redis Client is nil")
}