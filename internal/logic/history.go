package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
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
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil
	}
	exist, err := repo.History().GetExist(ctx, opt.Where("user_id = ? and video_id = ?", user.ID, req.VideoId))
	if err != nil {
		return err
	}
	if exist {
		return repo.History().Updates(ctx, map[string]any{"duration": req.Duration, "ip": req.IP, "resource_id": req.ResourceId}, opt.Where("user_id = ? and video_id = ?", user.ID, req.VideoId))
	}
	history := &entity.History{
		UserId:     user.ID,
		VideoId:    req.VideoId,
		ResourceId: req.ResourceId,
		Duration:   req.Duration,
		IP:         req.IP,
	}
	return repo.History().Add(ctx, history)
}

func (historyLogic) List(ctx context.Context, req *dto.HistoryList) ([]*entity.History, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	options := opt.NewOpts().Where("user_id = ?", user.ID).Order("updated_at desc, id desc").Preload("Video.User")

	if req.LastId != 0 {
		last, err := repo.History().GetById(ctx, req.LastId)
		if err != nil {
			return nil, err
		}
		if last == nil {
			return nil, errno.New400("未找到更多历史记录")
		}
		options = options.Where("updated_at <= ? and id < ?", last.UpdatedAt, last.ID)
	}

	return repo.History().GetList(ctx, options...)
}

func (historyLogic) Get(ctx context.Context, req *dto.VideoId) (*entity.History, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	history, err := repo.History().Get(ctx, opt.Where("user_id = ? and video_id = ?", user.ID, req.VideoId))
	if err != nil {
		return nil, err
	}
	return history, nil
}
