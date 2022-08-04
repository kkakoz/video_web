package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
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

		return repo.Collect().Add(ctx, &entity.Collect{
			UserId:     user.ID,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
			GroupId:    req.GroupId,
		})
	})

}
