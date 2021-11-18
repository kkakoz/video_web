package repo

import (
	"context"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/pkg/mysqlx"
)

var _ domain.IUserRepo = (*UserRepo)(nil)

type UserRepo struct {

}

func NewUserRepo() domain.IUserRepo {
	return &UserRepo{}
}

func (u UserRepo) AddUser(ctx context.Context, user *domain.User) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(user).Error
	return errors.Wrap(err, "添加失败")
}

func (u UserRepo) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	db := mysqlx.GetDB(ctx)
	user := &domain.User{}
	err := db.Find(user, id).Error
	return user, errors.Wrap(err, "查询失败")
}

func (u UserRepo) GetUserList(ctx context.Context, ids []int64) ([]*domain.User, error) {
	db  := mysqlx.GetDB(ctx)
	list := make([]*domain.User, 0, len(ids))
	err := db.Where("id in ?", ids).Find(&list).Error
	return list, errors.Wrap(err, "查询失败")
}

