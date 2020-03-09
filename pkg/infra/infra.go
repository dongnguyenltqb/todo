package infra

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rd *redis.Client

func Setup() {
	rd = redis.NewClient(&redis.Options{
		Addr:     viper.GetViper().GetString("redis.addr"),
		Password: viper.GetViper().GetString("redis.password"),
		DB:       viper.GetViper().GetInt("redis.db"),
	})
	if _, err := rd.Ping().Result(); err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return rd
}
