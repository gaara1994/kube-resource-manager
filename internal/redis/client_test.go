package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	InitRedis()
	err := SetKeyValue( "name", "tom")
	if err != nil {
		return
	}
	value, err := GetKeyValue("name")
	if err != nil {
		return 
	}
	fmt.Println("name =",value)
}
// SetKeyValue 设置键值对
func SetKeyValue( key string, value string) error {
	err := GetClient().Set(key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set key-value pair: %v", err)
	}
	return nil
}

// GetKeyValue 获取键对应的值
func GetKeyValue( key string) (string, error) {
	val, err := GetClient().Get( key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist: %s", key)
	} else if err != nil {
		return "", fmt.Errorf("failed to get value: %v", err)
	}
	return val, nil
}
