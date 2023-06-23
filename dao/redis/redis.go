package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func InitDB() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.poolSize"),    // 连接池大小
	})
	if _, err := rdb.Ping().Result(); err != nil {
		return err
	}
	return
}

func Close() {
	rdb.Close()
}
