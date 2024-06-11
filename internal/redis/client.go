package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitRedis() {
	// go test 使用固定的。使用配置文件需要全局启动，否则配置为空
	//url := fmt.Sprintf("redis://:%s@%s",config.Config.Redis.Password,config.Config.Redis.Address)
	opt, err := redis.ParseURL("redis://:123456@localhost:6379")
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
