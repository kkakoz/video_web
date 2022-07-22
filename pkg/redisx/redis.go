package redisx

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"sync"
	"video_web/pkg/conf"
)

type RedisOptions struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

func NewRedis(viper *viper.Viper) (*redis.Client, error) {
	viper.SetDefault("redisx.host", "127.0.0.1")
	viper.SetDefault("redisx.port", "6379")
	viper.SetDefault("redisx.db", "1")
	viper.SetDefault("redisx.password", "")
	viper.SetDefault("redisx.pool_size", 5)
	options := &redis.Options{
		Addr:     viper.GetString("redisx.host") + ":" + viper.GetString("redisx.port"),
		Password: viper.GetString("redisx.password"),
		DB:       viper.GetInt("redisx.db"),
		PoolSize: viper.GetInt("redisx.pool_size"),
	}

	client = redis.NewClient(options)
	_, err := client.Ping().Result()
	return client, errors.Wrap(err, "redis初始化失败")
}

var client *redis.Client

var once sync.Once

func Client() *redis.Client {
	once.Do(func() {
		viper := conf.Conf()
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
		if err != nil {
			panic(errors.Wrap(err, "redis初始化失败"))
		}
	})
	return client
}
