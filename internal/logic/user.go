package logic

import (
	"context"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strings"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/model"
	"video_web/internal/repo"

	"video_web/internal/pkg/keys"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
	"video_web/pkg/local"

	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
)

type UserLogic struct {
	userRepo         *repo.UserRepo
	userSecurityRepo *repo.UserSecurityRepo
	redis            *redis.Client
	followGroup      *repo.FollowGroupRepo
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

func (item *UserLogic) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return item.userRepo.GetById(ctx, id)
}

func (item *UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*model.User, error) {
	return item.userRepo.GetList(ctx, opt.In("id", ids))
}

func (item *UserLogic) Register(ctx context.Context, req *request.RegisterReq) (err error) {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		salt := cryption.UUID()
		user := &model.User{
			Name:  req.Name,
			Email: req.Email,
			State: 1,
		}
		err = item.userRepo.Add(ctx, user)
		if err != nil {
			e := &mysql.MySQLError{}
			if errors.As(err, &e) {
				// 唯一index冲突
				if e.Number == 1062 {
					return errno.New400("已经注册")
				}
			}
			return err
		}
		security := &model.UserSecurity{
			UserId:   user.ID,
			Password: cryption.Md5Str(req.Password + salt),
			Salt:     salt,
		}
		err = item.userSecurityRepo.Add(ctx, security)
		if err != nil {
			return err
		}
		return item.userInit(ctx, user)
	})
}

func (item *UserLogic) Login(ctx context.Context, req *request.LoginReq) (string, error) {
	options := opt.NewOpts()
	if strings.Contains(req.Name, "@") {
		options = options.Where("email = ?", req.Name)
	} else {
		options = options.Where("phone = ?", req.Name)
	}
	user, err := item.userRepo.Get(ctx, options...)
	if user == nil {
		return "", errno.New400("账号不存在")
	}
	if err != nil {
		return "", err
	}
	security, err := item.userSecurityRepo.Get(ctx, opt.Where("user_id = ?", user.ID))
	if err != nil {
		return "", err
	}
	if security.Password != cryption.Md5Str(req.Password+security.Salt) {
		return "", errno.New400("密码错误")
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

func NewUserLogic(userRepo *repo.UserRepo, userSecurityRepo *repo.UserSecurityRepo, redis *redis.Client, followGroup *repo.FollowGroupRepo) *UserLogic {
	return &UserLogic{userRepo: userRepo, userSecurityRepo: userSecurityRepo, redis: redis, followGroup: followGroup}
}
