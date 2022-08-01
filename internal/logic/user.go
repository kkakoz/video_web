package logic

import (
	"context"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strings"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/keys"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
	"video_web/pkg/redisx"

	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
)

type userLogic struct {
}

var userOnce sync.Once
var _user *userLogic

func User() *userLogic {
	userOnce.Do(func() {
		_user = &userLogic{}
	})
	return _user
}

func (item *userLogic) GetCurUser(ctx context.Context, token string) (*entity.User, error) {
	res, err := redisx.Client().WithContext(ctx).Get(keys.TokenKey(token)).Result()
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (item *userLogic) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return repo.User().GetById(ctx, id)
}

func (item *userLogic) GetUsers(ctx context.Context, ids []int64) ([]*entity.User, error) {
	return repo.User().GetList(ctx, opt.In("id", ids))
}

func (item *userLogic) Register(ctx context.Context, req *dto.Register) (err error) {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		salt := cryption.UUID()
		user := &entity.User{
			Name:  req.Name,
			Email: req.Email,
			State: 1,
		}
		err = repo.User().Add(ctx, user)
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
		security := &entity.UserSecurity{
			UserId:   user.ID,
			Password: cryption.Md5Str(req.Password + salt),
			Salt:     salt,
		}
		err = repo.UserSecurity().Add(ctx, security)
		if err != nil {
			return err
		}
		return item.userInit(ctx, user)
	})
}

func (item *userLogic) Login(ctx context.Context, req *dto.Login) (string, error) {
	options := opt.NewOpts()
	if strings.Contains(req.Name, "@") {
		options = options.Where("email = ?", req.Name)
	} else {
		options = options.Where("phone = ?", req.Name)
	}
	user, err := repo.User().Get(ctx, options...)
	if user == nil {
		return "", errno.New400("账号不存在")
	}
	if err != nil {
		return "", err
	}
	security, err := repo.UserSecurity().Get(ctx, opt.Where("user_id = ?", user.ID))
	if err != nil {
		return "", err
	}
	if security == nil {
		return "", errno.New400("账号不存在")
	}
	if security.Password != cryption.Md5Str(req.Password+security.Salt) {
		return "", errno.New400("密码错误")
	}
	token := cryption.UUID()
	target := &entity.User{}
	err = copier.Copy(target, user)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(target)
	if err != nil {
		return "", err
	}
	err = redisx.Client().Set(keys.TokenKey(token), data, 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (item *userLogic) userInit(ctx context.Context, user *entity.User) error {
	return repo.FollowGroup().AddList(ctx, []*entity.FollowGroup{{ // 添加默认关注分组
		UserId:    user.ID,
		Type:      entity.FollowGroupTypeNormal,
		GroupName: "默认关注",
	}, { // 特别关注分组
		UserId:    user.ID,
		Type:      entity.FollowGroupTypeSpecial,
		GroupName: "特别关注"}})
}
