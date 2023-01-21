package bootstrap

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/conf"
	"github.com/pkg/errors"
	"video_web/internal/model/entity"
	"video_web/pkg/logs"
)

func migrate() error {
	if _, err := ormx.New(conf.Conf(), ormx.WithLogger(logs.NewGLogger())); err != nil {
		return errors.WithMessage(err, "init orm failed")
	}

	db := ormx.DB(context.TODO())
	return db.AutoMigrate(&entity.User{}, &entity.Video{}, &entity.Resource{}, &entity.FollowGroup{}, &entity.Follow{}, &entity.Collect{}, &entity.CollectGroup{},
		&entity.Category{}, &entity.Comment{}, &entity.SubComment{}, &entity.UserSecurity{}, &entity.Like{}, &entity.Newsfeed{}, &entity.History{}, &entity.Notice{})

}
