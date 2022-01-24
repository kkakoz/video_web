package logic

import "go.uber.org/fx"

var Provider = fx.Provide(NewVideoLogic, NewUserLogic, NewCategoryLogic)
