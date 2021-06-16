package cache

import (
	"server/pkg/config"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	con := config.GetRedisConfig()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     con.Addr,
		Password: con.Password, // no password set
		DB:       0,            // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
