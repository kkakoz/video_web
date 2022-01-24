package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"video_web/internal/domain"
	"video_web/pkg/gormx"
	"video_web/pkg/mysqlx"
)

var _ domain.IUserRepo = (*UserRepo)(nil)

type UserRepo struct {
}

func NewUserRepo() domain.IUserRepo {
	return &UserRepo{}
}

func (u UserRepo) WithId(id int64) gormx.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func (u UserRepo) AddUser(ctx context.Context, user *domain.User) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(user).Error
	return errors.Wrap(err, "添加失败")
}

func (u UserRepo) GetUser(ctx context.Context, opts ...gormx.DBOption) (*domain.User, error) {
	db := mysqlx.GetDB(ctx)
	user := &domain.User{}
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(user).Error
	return user, errors.Wrap(err, "查询失败")
}

func (u UserRepo) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	db := mysqlx.GetDB(ctx)
	list := make([]*domain.User, 0, len(ids))
	err := db.Where("id in ?", ids).Find(&list).Error
	return list, errors.Wrap(err, "查询失败")
}
