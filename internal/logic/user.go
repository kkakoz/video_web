package logic

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/internal/pkg/keys"
	"github/kkakoz/video_web/pkg/cryption"
	"github/kkakoz/video_web/pkg/errno"
	"github/kkakoz/video_web/pkg/mysqlx"
)

var _ domain.IUserLogic = (*UserLogic)(nil)

type UserLogic struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
	redis    *redis.Client
}

func (u UserLogic) GetCurUser(ctx context.Context, token string) (*domain.User, error) {
	res, err := u.redis.Get(keys.TokenKey(token), ).Result()
	if err != nil {
		return nil, err
	}
	user := &domain.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return nil, err
	}
	return user, nil
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
	oldAuth, err := u.authRepo.GetAuthByIdentify(ctx, int32(req.IdentityType), req.Identifier)
	if oldAuth.ID != 0 {
		return errors.New("已经注册")
	}
	auth := &domain.Auth{
		IdentityType: int32(req.IdentityType),
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
		return "", errno.New400("账号不存在")
	}
	if auth.Credential != cryption.Md5Str(req.Credential) {
		return "", errno.New400("密码错误")
	}
	user, err := u.userRepo.GetUser(ctx, auth.UserId)
	if err != nil {
		return "", err
	}
	token := cryption.UUID()
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = u.redis.Set(keys.TokenKey(token), data, 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewUserLogic(userRepo domain.IUserRepo, authRepo domain.IAuthRepo, redis *redis.Client) domain.IUserLogic {
	return &UserLogic{userRepo: userRepo, authRepo: authRepo, redis: redis}
}
