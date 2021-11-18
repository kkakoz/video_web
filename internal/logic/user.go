package logic

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/pkg/cryption"
	"github/kkakoz/video_web/pkg/mysqlx"
)

var _ domain.IUserLogic = (*UserLogic)(nil)

type UserLogic struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
	redis    *redis.Client
}

func NewUserLogic(userRepo domain.IUserRepo, authRepo domain.IAuthRepo, redis *redis.Client) domain.IUserLogic {
	return &UserLogic{userRepo: userRepo, authRepo: authRepo, redis: redis}
}

func (u UserLogic) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	return u.userRepo.GetUserList(ctx, ids)
}

func (u UserLogic) Register(ctx context.Context, req *dto.RegisterReq) (err error) {
	ctx, checkError := mysqlx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	oldAuth, err := u.authRepo.GetAuthByIdentify(ctx, req.IdentityType, req.Identifier)
	if oldAuth.ID != 0 {
		return errors.New("已经注册")
	}
	auth := &domain.Auth{
		IdentityType: req.IdentityType,
		Identifier:   req.Identifier,
		Credential:   cryption.Md5Str(req.Credential),
	}
	user := &domain.User{
		Name:  req.Name,
		State: 1,
		Auth:  auth,
	}
	err = u.userRepo.AddUser(ctx, user)
	return err
}

func (u UserLogic) Login(ctx context.Context, req *dto.LoginReq) (string, error) {
	auth, err := u.authRepo.GetAuthByIdentify(ctx, req.IdentityType, req.Identifier)
	if err != nil {
		return "", err
	}
	if auth.ID == 0 {
		return "", errors.New("未找到账号")
	}
	if auth.Credential != req.Credential {
		return "", errors.New("密码错误")
	}
	user, err := u.userRepo.GetUser(ctx, auth.UserId)
	if err != nil {
		return "", err
	}
	token := cryption.UUID()
	err = u.redis.Set("user:token:"+token, user, 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}
