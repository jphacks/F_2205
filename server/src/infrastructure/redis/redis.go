package redis

import (
	"github.com/jphacks/F_2205/server/src/config"

	redis "github.com/go-redis/redis/v8"
)

// redis clientの生成
var Rs = redis.NewClient(&redis.Options{
	Addr: config.GetRedisUrl(),
})
