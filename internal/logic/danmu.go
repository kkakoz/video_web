package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"video_web/internal/model"
	"video_web/internal/pkg/ws"
	"video_web/internal/repo"
)

type DanmuLogic struct {
	danmuRepo *repo.DanmuRepo
	videoConn *ws.VideoConn
}

func NewDanmuLogic(danmuRepo *repo.DanmuRepo, conn *ws.VideoConn) *DanmuLogic {
	return &DanmuLogic{danmuRepo: danmuRepo, videoConn: conn}
}

func (item *DanmuLogic) Add(ctx context.Context, danmu *model.Danmu) error {
	err := item.danmuRepo.Add(ctx, danmu)
	if err != nil {
		return err
	}
	item.videoConn.Send(danmu.VideoId, ws.VideoWsRes{
		Type:    1,
		Content: danmu,
	})
	return nil
}

func (item *DanmuLogic) ListByVideoId(ctx context.Context, videoId int64) ([]*model.Danmu, error) {
	return item.danmuRepo.GetList(ctx, opt.Where("video_id = ?", videoId))
}
