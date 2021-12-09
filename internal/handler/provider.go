package handler

import (
	"github/kkakoz/video_web/internal/logic"
	"github/kkakoz/video_web/internal/repo"
	"go.uber.org/fx"
)

var Provider = fx.Options(userProvider, videoProvider, categoryProvider)

var userProvider = fx.Provide(NewUserHandler, logic.NewUserLogic, repo.NewUserRepo, repo.NewAuthRepo)

var videoProvider = fx.Provide(NewVideoHandler, logic.NewVideoLogic, repo.NewVideoRepo)

var categoryProvider = fx.Provide(NewCategoryHandler, logic.NewCategoryLogic, repo.NewCategoryRepo)