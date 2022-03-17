package logic_test

import (
	"context"
	"log"
	"testing"
	"video_web/internal/domain"
	"video_web/internal/logic"
	"video_web/internal/repo"
	"video_web/pkg/conf"
	"video_web/pkg/redis"

	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var userLogic domain.IUserLogic
var categoryLogin domain.ICategoryLogic
var videoLogic domain.IVideoLogic

func Init() error {
	conf.InitTestConfig()
	if _, err := ormx.New(viper.GetViper()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.FlushDB()
	db := ormx.DB(context.TODO())
	db.AutoMigrate(&domain.Auth{}, &domain.Comment{}, &domain.Count{},
		&domain.Category{}, &domain.Episode{}, &domain.User{}, &domain.Video{})
	return fx.New(
		logic.Provider,
		repo.Provider,
		redis.Provider,
		fx.Supply(viper.GetViper()),
		fx.Populate(&userLogic),
		fx.Populate(&categoryLogin),
		fx.Populate(&videoLogic),
	).Err()
}

func TestMain(m *testing.M) {
	err := Init()
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}
