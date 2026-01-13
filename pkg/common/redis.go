// pkg/common/redis.go
package common

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 修改成你的 Redis 地址
		Password: "",               // 无密码
		DB:       0,
	})

	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}
	log.Println("Redis connected successfully!")
}
