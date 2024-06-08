package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitRedis() {
	// 使用环境变量或配置文件中的 Redis 连接信息
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}

	rdb = redis.NewClient(opt)
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	log.Println("Connected to Redis!")
}

func GetClient() *redis.Client {
	return rdb
}
