package handler

import (
	"go.uber.org/fx"
	"video_web/internal/logic"
	"video_web/internal/repo"
)

var Provider = fx.Options(handlerProvider, logic.Provider, repo.Provider)

var handlerProvider = fx.Provide(NewUserHandler, NewVideoHandler, NewCategoryHandler, NewCommentHandler, NewLikeHandler)
