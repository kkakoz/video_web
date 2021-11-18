package handler

import (
	"github/kkakoz/video_web/internal/logic"
	"github/kkakoz/video_web/internal/repo"
	"go.uber.org/fx"
)

var Provider = fx.Provide(NewUserHandler, logic.NewUserLogic, repo.NewUserRepo, repo.NewAuthRepo)

