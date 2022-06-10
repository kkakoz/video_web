package logic_test

import (
	"context"
	"log"
	"testing"
	"video_web/internal/logic"
	"video_web/internal/model"
	"video_web/internal/repo"
	"video_web/pkg/conf"
	"video_web/pkg/redis"

	"github.com/kkakoz/ormx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var userLogic *logic.UserLogic
var categoryLogin *logic.CategoryLogic
var videoLogic *logic.VideoLogic

func Init() error {
	conf.InitTestConfig()
	if _, err := ormx.New(viper.GetViper()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.FlushDB()
	db := ormx.DB(context.TODO())
	db.AutoMigrate(&model.UserSecurity{}, &model.Comment{}, &model.Count{},
		&model.Category{}, &model.Episode{}, &model.User{}, &model.Video{}, &model.Follow{}, &model.FollowGroup{})
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
