package sdk

import (
	"Taichi/log"
	"Taichi/redis"
)

var Log = log.NewLog(nil, nil)

var Redis = redis.GetRedis()
