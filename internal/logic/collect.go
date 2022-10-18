package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
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
		exist, err := repo.Collect().GetExist(ctx, opt.Where("user_id = ? and target_id = ? ", user.ID, req.TargetId), opt.IsWhere(req.GroupId != 0, "group_id = ?", req.GroupId))
		if err != nil {
			return err
		}
		if exist {
			if req.Collect {
				return errno.New400("视频已经收藏")
			}
			return nil
		} else {
			if req.Collect {
				return repo.Collect().Add(ctx, &entity.Collect{
					UserId:   user.ID,
					TargetId: req.TargetId,
					GroupId:  req.GroupId,
				})
			}
			return nil
		}
	})

}
