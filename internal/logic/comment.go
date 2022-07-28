package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
)

type commentLogic struct {
}

var commentOnce sync.Once
var _comment *commentLogic

func Comment() *commentLogic {
	commentOnce.Do(func() {
		_comment = &commentLogic{}
	})
	return _comment
}

func (item *commentLogic) Add(ctx context.Context, req *dto.CommentAdd) error {
	comment := &entity.Comment{}
	err := copier.Copy(comment, req)
	if err != nil {
		return err
	}
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	comment.UserId = user.ID
	comment.Username = user.Name
	comment.Avatar = user.Avatar
	return repo.Comment().Add(ctx, comment)
}

func (item *commentLogic) AddSub(ctx context.Context, req *dto.SubCommentAdd) error {
	subComment := &entity.SubComment{}
	err := copier.Copy(subComment, req)
	if err != nil {
		return err
	}
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	subComment.FromId = user.ID
	subComment.FromName = user.Name
	subComment.FromAvatar = user.Avatar
	return repo.SubComment().Add(ctx, subComment)
}

// 查找评论和部分子评论
func (item *commentLogic) GetList(ctx context.Context, req *dto.CommentList) ([]*entity.Comment, error) {
	list, err := repo.Comment().GetList(ctx, opt.Where("target_id = ? and target_type = ?", req.TargetId, req.TargetType),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
	if err != nil {
		return nil, err
	}
	// 根据list返回 id list
	commentIds := lo.Map(list, func(t *entity.Comment, i int) int64 { return t.ID })
	// 根据id list查找 sub comment
	subComments, err := repo.SubComment().GetList(ctx, opt.Where("comment_id in ?", commentIds), opt.Limit(50))
	if err != nil {
		return nil, err
	}

	groupSubComment := lo.GroupBy(subComments, func(subComment *entity.SubComment) int64 { return subComment.CommentId }) // 根据id分组
	lo.ForEach(list, func(comment *entity.Comment, i int) {                                                               // 赋值
		comment.SubComments = []*entity.SubComment{}
		if subs, ok := groupSubComment[comment.ID]; ok {
			comment.SubComments = subs
		}
	})
	return list, nil
}

func (item *commentLogic) GetSubList(ctx context.Context, req *dto.SubCommentList) ([]*entity.SubComment, error) {
	return repo.SubComment().GetList(ctx, opt.Where("comment_id = ? ", req.CommentId),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
}

func (item *commentLogic) Delete(ctx context.Context, req *dto.CommentDel) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo.Comment().DeleteById(ctx, req.CommentId)
	})
}

func (item *commentLogic) DeleteSubComment(ctx context.Context, req *dto.SubCommentDel) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo.SubComment().DeleteById(ctx, req.SubCommentId)
	})
}
