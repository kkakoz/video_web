package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
)

type historyLogic struct {
}

var historyOnce sync.Once
var _history *historyLogic

func History() *historyLogic {
	historyOnce.Do(func() {
		_history = &historyLogic{}
	})
	return _history
}

func (historyLogic) Add(ctx context.Context, req *dto.HistoryAdd) error {
	user, ok := local.GetUserLocal(ctx)
	if !ok {
		return nil
	}
	exist, err := repo.History().GetExist(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.HistoryTypeVideo, req.VideoId))
	if err != nil {
		return err
	}
	if exist {
		return repo.History().Updates(ctx, map[string]any{"duration": req.Duration, "ip": req.IP}, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.HistoryTypeVideo, req.VideoId))
	}
	history := &entity.History{
		UserId:     user.ID,
		TargetType: entity.HistoryTypeVideo,
		TargetId:   req.VideoId,
		Duration:   req.Duration,
		IP:         req.IP,
	}
	return repo.History().Add(ctx, history)
}

func (historyLogic) List(ctx context.Context) ([]*entity.History, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return repo.History().GetList(ctx, opt.NewOpts().Where("user_id = ?", user.ID).Order("updated_at").Preload("User").Preload("Resource")...)
}
