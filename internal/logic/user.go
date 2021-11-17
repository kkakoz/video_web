package logic

import (
	"context"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
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

func (u UserLogic) Register(ctx context.Context, req *dto.RegisterReq) (err error) {
	ctx, checkError := mysqlx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	auth := &domain.Auth{
		IdentityType: req.IdentityType,
		Identifier:   req.Identifier,
		Credential:   req.Credential,
	}
	user := &domain.User{
		Name:        req.Name,
		Avatar:      "",
		Brief:       "",
		FollowCount: 0,
		FansCount:   0,
		LikeCount:   0,
		State:       1,
		Auth:        auth,
	}
	err = u.userRepo.AddUser(ctx, user)
	return err
}

func (u UserLogic) Login(ctx context.Context, user *domain.Auth) (string, error) {
	panic("implement me")
}
