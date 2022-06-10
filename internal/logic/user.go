package logic

import (
	"context"
	"encoding/json"
<<<<<<< HEAD
	"video_web/internal/domain"
=======
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"video_web/internal/consts"
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
	"video_web/internal/dto/request"
	"video_web/internal/model"
	"video_web/internal/repo"

	"video_web/internal/pkg/keys"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
	"video_web/pkg/local"
<<<<<<< HEAD
	"video_web/pkg/mysqlx"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
)
=======
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
)

type UserLogic struct {
	userRepo    *repo.UserRepo
	authRepo    *repo.AuthRepo
	redis       *redis.Client
	followGroup *repo.FollowGroupRepo
}

func (item *UserLogic) GetCurUser(ctx context.Context, token string) (*local.User, error) {
	res, err := item.redis.WithContext(ctx).Get(keys.TokenKey(token)).Result()
	if err != nil {
		return nil, err
	}
	user := &local.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

<<<<<<< HEAD
func (u UserLogic) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.Get(ctx, u.userRepo.WithId(id))
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	return u.userRepo.GetList(ctx, u.userRepo.WithIds(ids))
=======
func (item *UserLogic) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return item.userRepo.GetById(ctx, id)
}

func (item *UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*model.User, error) {
	return item.userRepo.GetList(ctx, opts.In("id", ids))
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
}

func (item *UserLogic) Register(ctx context.Context, req *request.RegisterReq) (err error) {
	ctx, checkError := ormx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
<<<<<<< HEAD
	oldAuth, err := u.authRepo.GetAuth(ctx, u.authRepo.WithIdentifierAndType(req.Identifier, req.IdentityType))
=======
	salt := cryption.UUID()
	user := &model.User{
		Name:  req.Name,
		State: 1,
	}
	err = item.userRepo.Add(ctx, user)
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
	if err != nil {
		return err
	}
	auth := &model.Auth{
		IdentityType: req.IdentityType,
		Identifier:   req.Identifier,
		Credential:   cryption.Md5Str(req.Credential + salt),
		Salt:         salt,
		UserId:       user.ID,
	}
	err = item.authRepo.Add(ctx, auth)
	if err != nil {
		e := &mysql.MySQLError{}
		if errors.As(err, &e) {
			if e.Number == 1062 {
				return errno.New400("已经注册")
			}
		}
		return err
	}
<<<<<<< HEAD
	err = u.userRepo.Add(ctx, user)
	return err
=======
	return item.userInit(ctx, user)
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
}

func (item *UserLogic) Login(ctx context.Context, req *request.LoginReq) (string, error) {
	auth, err := item.authRepo.Get(ctx, opts.Where("identity_type = ? and identifier = ?",
		req.IdentityType, req.Identifier))
	if err != nil {
		return "", err
	}
	if auth.ID == 0 {
		return "", errno.New400("账号不存在")
	}
	if auth.Credential != cryption.Md5Str(req.Credential+auth.Salt) {
		return "", errno.New400("密码错误")
	}
<<<<<<< HEAD
	user, err := u.userRepo.Get(ctx, u.userRepo.WithId(auth.UserId))
=======
	user, err := item.userRepo.GetById(ctx, auth.UserId)
>>>>>>> eb83ab769aa25fe40b1cffa3e54381d191dee291
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
	err = item.redis.Set(keys.TokenKey(token), data, 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (item *UserLogic) userInit(ctx context.Context, user *model.User) error {
	return item.followGroup.AddList(ctx, []*model.FollowGroup{{ // 添加默认关注分组
		UserId:    user.ID,
		Type:      consts.FollowGroupTypeDefault,
		GroupName: "默认关注",
	}, { // 特别关注分组
		UserId:    user.ID,
		Type:      consts.FollowGroupTypeSpecial,
		GroupName: "特别关注"}})
}

func NewUserLogic(userRepo *repo.UserRepo, authRepo *repo.AuthRepo, redis *redis.Client, followGroup *repo.FollowGroupRepo) *UserLogic {
	return &UserLogic{userRepo: userRepo, authRepo: authRepo, redis: redis, followGroup: followGroup}
}
