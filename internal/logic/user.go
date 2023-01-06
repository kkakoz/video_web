package logic

import (
	"context"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"strings"
	"sync"
	"time"
	"video_web/internal/async/event_send"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/keys"
	"video_web/internal/pkg/local"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
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

func (userLogic) GetCurUser(ctx context.Context, token string) (*entity.User, error) {
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

func (userLogic) GetUser(ctx context.Context, id int64) (*vo.User, error) {
	user, err := repo.User().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &vo.User{}
	err = copier.Copy(res, user)
	if err != nil {
		return nil, err
	}
	current, exist := local.GetUserExist(ctx)
	if exist {
		followed, err := repo.Follow().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", current.ID, id))
		if err != nil {
			return nil, err
		}
		res.Followed = followed
	}
	return res, nil

}

func (userLogic) GetUsers(ctx context.Context, ids []int64) ([]*entity.User, error) {
	return repo.User().GetList(ctx, opt.In("id", ids))
}

func (item userLogic) Register(ctx context.Context, req *dto.Register) (err error) {
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
		return event_send.SendEvent(&dto.Event{
			EventType:     dto.EventTypeUserRegister,
			TargetId:      user.ID,
			TargetType:    0,
			ActorId:       user.ID,
			TargetOwnerId: user.ID,
		})
	})
}

func (userLogic) Login(ctx context.Context, req *dto.Login) (string, error) {
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
	err = redisx.Client().Set(keys.TokenKey(token), data, time.Hour*24*3).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (userLogic) UserInit(ctx context.Context, userId int64) error {
	return repo.FollowGroup().AddList(ctx,
		[]*entity.FollowGroup{
			{ // 添加默认关注分组
				UserId:    userId,
				Type:      entity.FollowGroupTypeNormal,
				GroupName: "默认关注",
			}, { // 特别关注分组
				UserId:    userId,
				Type:      entity.FollowGroupTypeSpecial,
				GroupName: "特别关注",
			},
		},
	)
}

func (item userLogic) UpdateAvatar(ctx context.Context, req *dto.UpdateAvatar) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return repo.User().Updates(ctx, map[string]any{
		"avatar": req.Url,
	}, opt.Where("id = ?", user.ID))
}
