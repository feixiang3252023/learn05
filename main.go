package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建一个新的 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 没有密码时设置为空
		DB:       0,                // 默认 Database 选择
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")

	// 存储数据
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// 获取数据
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Printf("key: %s\n", val)

	// 关闭 Redis 客户端
	defer rdb.Close()
}
