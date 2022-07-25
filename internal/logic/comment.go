package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	repo2 "video_web/internal/logic/internal/repo"
	"video_web/internal/model"
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

func (item *commentLogic) Add(ctx context.Context, req *request.CommentAddReq) error {
	comment := &model.Comment{}
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
	return repo2.Comment().Add(ctx, comment)
}

func (item *commentLogic) AddSub(ctx context.Context, req *request.SubCommentAddReq) error {
	subComment := &model.SubComment{}
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
	return repo2.SubComment().Add(ctx, subComment)
}

// 查找评论和部分子评论
func (item *commentLogic) GetList(ctx context.Context, req *request.CommentListReq) ([]*model.Comment, error) {
	list, err := repo2.Comment().GetList(ctx, opt.Where("target_id = ? and target_type = ?", req.TargetId, req.TargetType),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
	if err != nil {
		return nil, err
	}
	// 根据list返回 id list
	commentIds := lo.Map(list, func(t *model.Comment, i int) int64 { return t.ID })
	// 根据id list查找 sub comment
	subComments, err := repo2.SubComment().GetList(ctx, opt.Where("comment_id in ?", commentIds), opt.Limit(50))
	if err != nil {
		return nil, err
	}

	groupSubComment := lo.GroupBy(subComments, func(subComment *model.SubComment) int64 { return subComment.CommentId }) // 根据id分组
	lo.ForEach(list, func(comment *model.Comment, i int) {                                                               // 赋值
		comment.SubComments = []*model.SubComment{}
		if subs, ok := groupSubComment[comment.ID]; ok {
			comment.SubComments = subs
		}
	})
	return list, nil
}

func (item *commentLogic) GetSubList(ctx context.Context, req *request.SubCommentListReq) ([]*model.SubComment, error) {
	return repo2.SubComment().GetList(ctx, opt.Where("comment_id = ? ", req.CommentId),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
}

func (item *commentLogic) Delete(ctx context.Context, req *request.CommentDelReq) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo2.Comment().DeleteById(ctx, req.CommentId)
	})
}

func (item *commentLogic) DeleteSubComment(ctx context.Context, req *request.SubCommentDelReq) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo2.SubComment().DeleteById(ctx, req.SubCommentId)
	})
}
