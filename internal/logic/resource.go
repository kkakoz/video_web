package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
)

type resourceLogic struct {
}

var resourceOnce sync.Once
var _resource *resourceLogic

func Resource() *resourceLogic {
	resourceOnce.Do(func() {
		_resource = &resourceLogic{}
	})
	return _resource
}

func (resourceLogic) Add(ctx context.Context, req *dto.ResourceAdd) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode := &entity.Resource{
		Name:     req.Name,
		Brief:    req.Brief,
		Url:      req.Url,
		UserId:   user.ID,
		Duration: req.Duration,
	}
	err = repo.Resource().Add(ctx, episode)
	return err
}

func (resourceLogic) DelVideo(ctx context.Context, req *dto.ResourceId) (err error) {
	return repo.Resource().DeleteById(ctx, req.ResourceId)
}

// 获取视频详情
func (resourceLogic) Get(ctx context.Context, resourceId int64) (*entity.Resource, error) {
	video, err := repo.Resource().Get(ctx, opt.NewOpts().Where("id = ? ", resourceId).Preload("User")...)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (resourceLogic) GetPageList(ctx context.Context, req *dto.BackResourceList) ([]*entity.Resource, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	options = options.IsWhere(req.CategoryId > 0, "category_id = ?", req.CategoryId)
	count, err := repo.Resource().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset())

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Resource().GetList(ctx, options...)
	return list, count, err
}

func (resourceLogic) AddList(ctx context.Context, req *dto.ResourceAddList) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	resources := make([]*entity.Resource, len(req.Resources))
	for i, v := range req.Resources {
		resources[i] = &entity.Resource{
			Name:     v.Name,
			Brief:    v.Brief,
			Url:      v.Url,
			UserId:   user.ID,
			Duration: v.Duration,
			VideoId:  req.VideoId,
			Sort:     int64(i),
		}
	}

	return repo.Resource().AddList(ctx, resources)
}
