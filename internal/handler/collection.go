package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/echox"
)

type collectionHandler struct {
}

var collectionOnce sync.Once
var _collection *collectionHandler

func Collection() *collectionHandler {
	collectionOnce.Do(func() {
		_collection = &collectionHandler{}
	})
	return _collection
}

func (item *collectionHandler) Add(ctx echo.Context) error {
	req := &dto.CollectionAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Collection().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *collectionHandler) BackList(ctx echo.Context) error {
	req := &dto.BackCollectionList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	collections, count, err := logic.Video().GetPageCollections(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"count": count,
		"data":  collections,
	})
}

func (item *collectionHandler) Get(ctx echo.Context) error {
	req := &dto.CollectionId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	collections, err := logic.Video().GetCollection(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, collections)
}
