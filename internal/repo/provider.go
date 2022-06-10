package repo

import "go.uber.org/fx"

var Provider = fx.Provide(NewUserSecurityRepo, NewUserRepo, NewVideoRepo, NewCategoryRepo,
	NewEpisodeRepo, NewCommentRepo, NewSubCommentRepo, NewLikeRepoRepo,
	NewFollowRepo, NewFollowGroupRepo, NewDanmuRepo)
