package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type subCommentRepo struct {
	ormx.IRepo[entity.SubComment]
}

var subCommentOnce sync.Once
var _subComment *subCommentRepo

func SubComment() *subCommentRepo {
	subCommentOnce.Do(func() {
		_subComment = &subCommentRepo{IRepo: ormx.NewRepo[entity.SubComment]()}
	})
	return _subComment
}
