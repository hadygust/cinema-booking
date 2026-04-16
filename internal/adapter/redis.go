package adapter

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Printf("Ping status: %v", err)
	}
	log.Printf("connected to address: %s", addr)

	return rdb
}
