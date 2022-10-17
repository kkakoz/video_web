package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx/opt"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
)

type danmuLogic struct {
}

var danmuOnce sync.Once
var _danmu *danmuLogic

func Danmu() *danmuLogic {
	danmuOnce.Do(func() {
		_danmu = &danmuLogic{}
	})
	return _danmu
}

func (danmuLogic) Add(ctx context.Context, req *dto.DanmuAdd) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	danmu := &entity.Danmu{}
	err = copier.Copy(danmu, req)
	if err != nil {
		return err
	}
	danmu.UserId = user.ID

	err = repo.Danmu().Add(ctx, danmu)
	if err != nil {
		return err
	}
	//item.videoConn.Send(danmu.ResourceId, ws.VideoWsRes{
	//	Type:    1,
	//	Content: danmu,
	//})
	return nil
}

func (item *danmuLogic) ListByVideoId(ctx context.Context, videoId int64) ([]*entity.Danmu, error) {
	return repo.Danmu().GetList(ctx, opt.Where("video_id = ?", videoId))
}
