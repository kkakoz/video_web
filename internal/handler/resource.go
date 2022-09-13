package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/pkg/ws"
	"video_web/pkg/echox"
)

type resourceHandler struct {
}

var resourceOnce sync.Once
var resource *resourceHandler

func Resource() *resourceHandler {
	resourceOnce.Do(func() {
		resource = &resourceHandler{}
	})
	return resource
}

func (item *resourceHandler) Add(ctx echo.Context) error {
	req := &dto.ResourceAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Resource().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *resourceHandler) AddList(ctx echo.Context) error {
	req := &dto.ResourceAddList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Resource().AddList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
func (item *resourceHandler) Get(ctx echo.Context) error {
	req := &dto.ResourceId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := logic.Resource().Get(ctx.Request().Context(), req.ResourceId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

//func (item *resourceHandler) GetList(ctx echo.Context) error {
//	req := &dto.Videos{}
//	err := ctx.Bind(req)
//	if err != nil {
//		return err
//	}
//	videos, err := logic.Resource().GetVideos(ctx.Request().Context(), req.CategoryId, req.LastValue, req.OrderType)
//	if err != nil {
//		return err
//	}
//	return ctx.JSON(200, map[string]any{
//		"data": videos,
//	})
//}

func (item *resourceHandler) GetBackList(ctx echo.Context) error {
	req := &dto.BackResourceList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, count, err := logic.Resource().GetPageList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"count": count,
		"data":  videos,
	})
}

func (item *resourceHandler) Del(ctx echo.Context) error {
	req := &dto.ResourceId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Resource().DelVideo(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *resourceHandler) Ws(ctx echo.Context) error {
	req := &dto.ResourceId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return ws.VideoConn().Add(ctx.Response(), ctx.Request(), req.ResourceId)
}

func (item *resourceHandler) List(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	data, err := logic.Resource().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": data,
	})
}
