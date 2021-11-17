package logic

import (
	"context"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/pkg/mysqlx"
)

type UserLogic struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
}

func (u UserLogic) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	panic("implement me")
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	panic("implement me")
}

func (u UserLogic) Register(ctx context.Context, auth *domain.Auth) (err error) {
	ctx, checkError := mysqlx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	user := &domain.User{
		Name:        "",
		Avatar:      "",
		Brief:       "",
		FollowCount: 0,
		FansCount:   0,
		LikeCount:   0,
		State:       1,
		Auth:        auth,
	}
	u.userRepo.AddUser(ctx, user)
	return nil
}

func (u UserLogic) Login(ctx context.Context, user *domain.Auth) (string, error) {
	panic("implement me")
}
