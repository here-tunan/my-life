package infrastructure

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-my-life/env"
)

var Redis *redis.Client

// Redis 获取Redis 客户端
func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.Prop.Redis.HOST, env.Prop.Redis.Port),
		Password: env.Prop.Redis.Password, // no password set
		DB:       env.Prop.Redis.Db,       // use default DB
	})
}
