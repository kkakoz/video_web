package cache

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type RedisOptions struct {
	Host string
	Port string
	Password string
	DB int
	PoolSize int
}

func NewRedis(viper *viper.Viper) (*redis.Client, error) {
	o := &RedisOptions{}
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", "1")
	viper.SetDefault("redis.password", "")
	err := viper.UnmarshalKey("redis", o)
	if err != nil {
		return nil, errors.Wrap(err, "viper unmarshal失败")
	}
	options := &redis.Options{
		Addr:               o.Host + ":" + o.Port,
		Password:           o.Password,
		DB:                 o.DB,
		PoolSize:           o.PoolSize,
	}

	client := redis.NewClient(options)
	_, err = client.Ping().Result()
	return client, errors.Wrap(err, "redis初始化失败")
}

var Provider = fx.Provide(NewRedis)