package logic

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/internal/pkg/keys"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
	"video_web/pkg/gormx"
	"video_web/pkg/local"
	"video_web/pkg/mysqlx"
)

var _ domain.IUserLogic = (*UserLogic)(nil)

type UserLogic struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
	redis    *redis.Client
}

func (u UserLogic) GetCurUser(ctx context.Context, token string) (*domain.User, error) {
	res, err := u.redis.WithContext(ctx).Get(keys.TokenKey(token)).Result()
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
	return u.userRepo.GetUser(ctx, u.userRepo.WithId(id))
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	return u.userRepo.GetUsers(ctx, ids)
}

func (u UserLogic) Register(ctx context.Context, req *request.RegisterReq) (err error) {
	ctx, checkError := mysqlx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	oldAuth, err := u.authRepo.GetAuth(ctx, gormx.WithWhere("identity_type = ? and identifier = ?",
		req.IdentityType, req.Identifier))
	if err != nil {
		return err
	}
	if oldAuth.ID != 0 {
		return errno.New400("已经注册")
	}
	salt := cryption.UUID()
	auth := &domain.Auth{
		IdentityType: req.IdentityType,
		Identifier:   req.Identifier,
		Credential:   cryption.Md5Str(req.Credential + salt),
		Salt:         salt,
	}
	user := &domain.User{
		Name:  req.Name,
		State: 1,
		Auth:  auth,
	}
	err = u.userRepo.AddUser(ctx, user)
	return err
}

func (u UserLogic) Login(ctx context.Context, req *request.LoginReq) (string, error) {
	auth, err := u.authRepo.GetAuth(ctx, gormx.WithWhere("identity_type = ? and identifier = ?",
		req.IdentityType, req.Identifier))
	if err != nil {
		return "", err
	}
	if auth.ID == 0 {
		return "", errno.New400("账号不存在")
	}
	if auth.Credential != cryption.Md5Str(req.Credential) {
		return "", errno.New400("密码错误")
	}
	user, err := u.userRepo.GetUser(ctx, gormx.WithWhere("id = ?", auth.UserId))
	if err != nil {
		return "", err
	}
	token := cryption.UUID()
	target := &local.User{}
	err = copier.Copy(target, user)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(target)
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
