package bootstrap

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"video_web/internal/model/entity"
)

func migrate() error {
	logger.InitLog(viper.GetViper())
	if _, err := ormx.New(viper.GetViper()); err != nil {
		return errors.WithMessage(err, "init orm failed")
	}

	db := ormx.DB(context.TODO())
	return db.AutoMigrate(&entity.User{}, &entity.Video{}, &entity.Resource{}, &entity.FollowGroup{}, &entity.Follow{}, &entity.Collect{}, &entity.CollectGroup{},
		&entity.Category{}, &entity.Comment{}, &entity.SubComment{}, &entity.UserSecurity{}, &entity.Like{}, &entity.Newsfeed{}, &entity.History{}, &entity.Notice{})

}
