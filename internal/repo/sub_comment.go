package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type subCommentRepo struct {
	ormx.IRepo[model.SubComment]
}

var subCommentOnce sync.Once
var _subComment *subCommentRepo

func SubComment() *subCommentRepo {
	subCommentOnce.Do(func() {
		_subComment = &subCommentRepo{IRepo: ormx.NewRepo[model.SubComment]()}
	})
	return _subComment
}
