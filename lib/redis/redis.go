package redis

import (
    "media-matrix/config"
    "github.com/go-redis/redis"
	"sync"
)

var redisDatabases sync.Map



func init() {

	client := redis.NewClient(&redis.Options{
				Addr:config.RedisAddr,
				Password:config.RedisPassword,
				DB:config.RedisDb,
				PoolSize:config.RedisPoolSize,
		})
	redisDatabases.Store("default", client)
}

func GetRedis() *redis.Client{
	value, ok := redisDatabases.Load("default")
	if ok {
		return value.(*redis.Client)
	}
	panic("failed to connect redis database:" +"default")
}

func GetRedisClient(name string) *redis.Client{
	value, ok := redisDatabases.Load(name)
	if ok {
	   return value.(*redis.Client)
    }
	panic("failed to connect redis database:" +name)
}