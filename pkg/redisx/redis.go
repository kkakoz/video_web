package redisx

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type RedisOptions struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

func Init(viper *viper.Viper) error {
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", "1")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.pool_size", 5)
	options := &redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	}

	client = redis.NewClient(options)
	_, err := client.Ping().Result()
	return errors.Wrap(err, "redis初始化失败")
}

var client *redis.Client

func Client() *redis.Client {
	return client
}
