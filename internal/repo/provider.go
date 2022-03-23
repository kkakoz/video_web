package repo

import "go.uber.org/fx"

var Provider = fx.Provide(NewAuthRepo, NewUserRepo, NewVideoRepo, NewCategoryRepo, NewEpisodeRepo, NewCommentRepo, NewSubCommentRepo, NewLikeRepoRepo)
