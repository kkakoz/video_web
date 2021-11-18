package repo

import (
	"context"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/pkg/mysqlx"
)

var _ domain.IAuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {

}

func NewAuthRepo() domain.IAuthRepo {
	return &AuthRepo{}
}

func (a AuthRepo) GetAuth(ctx context.Context, id int64) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	auth := &domain.Auth{}
	err := db.Find(auth, id).Error
	return auth, errors.Wrap(err, "查询失败")
}

func (a AuthRepo) DeleteAuth(ctx context.Context, id int64) error {
	panic("implement me")
}

func (a AuthRepo) GetAuthByIdentify(ctx context.Context, identityType int32, identifier string) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	auth := &domain.Auth{}
	err := db.Where("identity_type = ? and identifier = ?", identityType, identifier).Find(auth).Error
	return auth, errors.Wrap(err, "查询失败")
}

