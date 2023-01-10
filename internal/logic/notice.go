package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
)

type noticeLogic struct{}

var noticeOnce sync.Once
var _notice *noticeLogic

func Notice() *noticeLogic {
	noticeOnce.Do(func() {
		_notice = &noticeLogic{}
	})
	return _notice
}

func (noticeLogic) Add(ctx context.Context, notice *entity.Notice) error {
	return repo.Notice().Add(ctx, notice)
}

func (noticeLogic) UnReadCount(ctx context.Context) (int64, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return 0, err
	}
	return repo.Notice().Count(ctx, opt.Where("user_id = ? and is_read = ?", user.ID, false))
}

func (noticeLogic) ReadNotice(ctx context.Context, id int64) error {
	return repo.Notice().Updates(ctx, map[string]any{
		"is_read": true,
	}, opt.Where("id = ?", id))
}

func (noticeLogic) GetList(ctx context.Context, lastId int64) ([]*entity.Notice, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	opts := opt.NewOpts().Where("user_id = ? ", user.ID).Limit(consts.DefaultLimit).Preload("FromUser").Order("created_at desc, id desc")
	if lastId != 0 {
		lastNotice, err := repo.Notice().GetById(ctx, lastId)
		if err != nil {
			return nil, err
		}
		if lastNotice == nil {
			return []*entity.Notice{}, nil
		}
		opts = opts.Where("created_at <= ? and id < ?", lastNotice.CreatedAt, lastNotice.ID)
	}

	return repo.Notice().GetList(ctx, opts...)
}

func (noticeLogic) ReadAll(ctx context.Context) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return repo.Notice().Updates(ctx, map[string]any{
		"is_read": true,
	}, opt.Where("user_id = ? and is_read = ?", user.ID, false))
}
