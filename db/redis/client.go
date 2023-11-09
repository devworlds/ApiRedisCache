package redis

import (
	"github.com/go-redis/redis"
)

func ClientInstance() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	println(rdb.Ping().String())

	return rdb
}
