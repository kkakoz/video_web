package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type RedisOptions struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

func NewRedis(viper *viper.Viper) (*redis.Client, error) {
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

	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	return client, errors.Wrap(err, "redis初始化失败")
}

var Provider = fx.Provide(NewRedis)
