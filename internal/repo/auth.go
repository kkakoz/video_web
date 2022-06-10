package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"video_web/internal/domain"
	"video_web/pkg/gormx"
	"video_web/pkg/mysqlx"
)

var _ domain.IAuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {
}

func NewAuthRepo() domain.IAuthRepo {
	return &AuthRepo{}
}

func (a AuthRepo) WithIdentifierAndType(identifier string, identityType int32) gormx.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("identity_type = ? and identifier = ?",
			identityType, identifier)
	}
}

func (a AuthRepo) GetAuth(ctx context.Context, opts ...gormx.DBOption) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	auth := &domain.Auth{}
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(auth).Error
	return auth, errors.Wrap(err, "查询失败")
}

func (a AuthRepo) DeleteAuth(ctx context.Context, id int64) error {
	return nil
}

func (a AuthRepo) GetAuthByIdentify(ctx context.Context, identityType int32, identifier string) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	auth := &domain.Auth{}
	err := db.Where("identity_type = ? and identifier = ?", identityType, identifier).Find(auth).Error
	return auth, errors.Wrap(err, "查询失败")
}
