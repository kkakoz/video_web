package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"strings"
	"sync"
	"time"
	"video_web/internal/async/producer"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/emailx"
	"video_web/internal/pkg/keys"
	"video_web/internal/pkg/local"
	"video_web/pkg/cryption"
	"video_web/pkg/errno"
)

type userLogic struct {
	//userClient userpb.UserHandlerClient
}

var userOnce sync.Once
var _user *userLogic

func User() *userLogic {
	userOnce.Do(func() {
		//userClient, err := client.NewUserClient(context.Background(), etcdx.Client())
		//if err != nil {
		//	log.Fatal(err)
		//}
		_user = &userLogic{}
	})
	return _user
}

func (userLogic) GetCurUser(ctx context.Context, token string) (*entity.User, error) {
	res, err := redisx.Client().WithContext(ctx).Get(ctx, keys.TokenKey(token)).Result()
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
	res.Followed = Follow().FollowedUser(ctx, id)
	return res, nil

}

func (userLogic) GetUsers(ctx context.Context, ids []int64) ([]*entity.User, error) {
	return repo.User().GetList(ctx, opt.In("id", ids))
}

func (item userLogic) Register(ctx context.Context, req *dto.Register) (err error) {
	return ormx.Transaction(ctx, func(ctx context.Context) error {

		salt := cryption.UUID()
		user := &entity.User{
			Name:   req.Name,
			Email:  req.Email,
			State:  1,
			Avatar: "https://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/avatar/default_avatar.gif",
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

		// 发送激活邮件
		code := uuid.NewV4().String()
		_, err := redisx.Client().Set(ctx, keys.UserActive(user.ID), code, time.Hour*24*3).Result()
		if err != nil {
			return err
		}
		err = emailx.Send(req.Email, "感谢注册", fmt.Sprintf(html, viper.GetString("app.addr")+fmt.Sprintf("/user/active?code=%s&user_id=%d", code, user.ID)))
		if err != nil {
			return err
		}

		// 发送注册事件 初始化
		return producer.SendVideoEvent(&dto.Event{
			EventType:     dto.EventTypeUserRegister,
			TargetId:      user.ID,
			TargetType:    0,
			ActorId:       user.ID,
			TargetOwnerId: user.ID,
		})
	})
}

var html = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8"/>
</head>

<body>
    <p>感谢注册,点击<a href="%s">此处</a>激活账户</p>
</body>

</html>
`

func (u *userLogic) Login(ctx context.Context, req *dto.Login) (string, error) {
	//login, err := u.userClient.Login(ctx, &userpb.LoginReq{
	//	Name:     req.Name,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return "", err
	//}
	//return login.Token, nil
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

	if user.State == entity.UserStateRegister {
		return "", errno.New400("账号未激活")
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
	err = redisx.Client().Set(ctx, keys.TokenKey(token), data, time.Hour*24*3).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (userLogic) UserInit(ctx context.Context, userId int64) error {
	return repo.FollowGroup().AddList(ctx,
		[]*entity.FollowGroup{
			{ // 特别关注分组
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

func (item userLogic) Active(ctx context.Context, req *dto.UserActive) error {
	code, err := redisx.Client().Get(ctx, keys.UserActive(req.UserId)).Result()
	if err != nil {
		return err
	}
	if code != req.Code {
		return errno.New400("激活码错误")
	}
	user, err := repo.User().GetById(ctx, req.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errno.New400("未找到用户")
	}
	if user.State == entity.UserStateActive {
		return errno.New400("用户已激活")
	}
	return repo.User().Updates(ctx, map[string]any{
		"active": entity.UserStateActive,
	}, opt.Where("id = ?", req.UserId))
}
