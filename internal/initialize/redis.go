package initialize

import (
	"context"
	"fmt"
	"starter/mod/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis connect failed, err:", zap.Error(err))
	}

	fmt.Println("redis connect success")
	global.Redis = rdb
	redisExample()
}

func redisExample() {
	err := global.Redis.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("redis set failed, err:", err)
		return
	}

	value, err := global.Redis.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("redis get failed, err:", err)
		return
	}
	fmt.Println("key", value)
	global.Logger.Info("value from redis", zap.String("key", value))
}
