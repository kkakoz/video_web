package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
)

type collectLogic struct {
}

var collectOnce sync.Once
var _collect *collectLogic

func Collect() *collectLogic {
	collectOnce.Do(func() {
		_collect = &collectLogic{}
	})
	return _collect
}

func (collectLogic) Add(ctx context.Context, req *dto.CollectAdd) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	return ormx.Transaction(ctx, func(ctx context.Context) error {
		err = repo.Video().Updates(ctx, map[string]any{
			"collect": gorm.Expr("collect + ?", lo.Ternary(req.Collect, 1, -1)),
		}, opt.Where("id = ?", req.TargetId))
		if err != nil {
			return err
		}
		if req.Collect {
			exist, err := repo.Collect().GetExist(ctx, opt.Where("user_id = ? and target_id = ? ", user.ID, req.TargetId), opt.IsWhere(req.GroupId != 0, "group_id = ?", req.GroupId))
			if err != nil {
				return err
			}
			if exist {
				if req.Collect {
					return errno.New400("视频已经收藏")
				}
				return nil
			}
			return repo.Collect().Add(ctx, &entity.Collect{
				UserId:  user.ID,
				VideoId: req.TargetId,
				GroupId: req.GroupId,
			})
		} else {
			return repo.Collect().Delete(ctx, opt.Where("user_id = ? and target_id = ?", user.ID, req.TargetId))
		}

	})

}

func (collectLogic) List(ctx context.Context, req *dto.CollectList) ([]*entity.Collect, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	options := opt.NewOpts().Where("user_id = ?", user.ID).Order("created_at desc, id desc").Preload("Video.User")

	if req.LastId != 0 {
		last, err := repo.Collect().GetById(ctx, req.LastId)
		if err != nil {
			return nil, err
		}
		if last == nil {
			return nil, errno.New400("未找到更多")
		}
		options = options.Where("created_at <= ? and id < ?", last.CreatedAt, last.ID)
	}

	return repo.Collect().GetList(ctx, options...)
}
