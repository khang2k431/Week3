package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic("Redis not connected")
	}
	fmt.Println("Redis connected")
}

// Presence tracking
func SetUserOnline(user string) {
	Rdb.Set(ctx, "presence:"+user, "online", 0)
}

func SetUserOffline(user string) {
	Rdb.Del(ctx, "presence:" + user)
}
