package redis

import (
	"Taichi/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetRedis(redisConfig config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Addr, redisConfig.Port),
		Password: redisConfig.Pwd, // 没有密码，默认值
		DB:       redisConfig.Db,  // 默认DB 0
	})

	return rdb
}
