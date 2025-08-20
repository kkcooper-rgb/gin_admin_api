package core

import (
	"go_admin_api/config"
	"go_admin_api/global"

	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

func RedisInit() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: "",
		DB:       config.Config.Redis.Db,
	})
	_, err := RedisDb.Ping(global.Ctx).Result()
	if err != nil {
		return err
	}
	global.Log.Infof("redis连接成功")
	return err
}
