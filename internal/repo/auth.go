package repo

import (
<<<<<<< HEAD
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"video_web/internal/domain"
	"video_web/pkg/gormx"
	"video_web/pkg/mysqlx"
)
=======
	"video_web/internal/model"
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291

	"github.com/kkakoz/ormx"
)

type AuthRepo struct {
	ormx.IRepo[model.Auth]
}

<<<<<<< HEAD
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
=======
func NewAuthRepo() *AuthRepo {
	return &AuthRepo{
		ormx.NewRepo[model.Auth](),
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
	}
}
