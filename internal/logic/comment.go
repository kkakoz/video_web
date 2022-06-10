package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx/opts"
	"github.com/samber/lo"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/model"
	"video_web/internal/repo"
	"video_web/pkg/local"
)

type CommentLogic struct {
	commentRepo    *repo.CommentRepo
	subCommentRepo *repo.SubCommentRepo
}

func (item *CommentLogic) Add(ctx context.Context, req *request.CommentAddReq) error {
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
	return item.commentRepo.Add(ctx, comment)
}

func (item *CommentLogic) AddSub(ctx context.Context, req *request.SubCommentAddReq) error {
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
	return item.subCommentRepo.Add(ctx, subComment)
}

// 查找评论和部分子评论
func (item *CommentLogic) GetList(ctx context.Context, req *request.CommentListReq) ([]*model.Comment, error) {
	list, err := item.commentRepo.GetList(ctx, opts.Where("target_id = ? and target_type = ?", req.TargetId, req.TargetType),
		opts.IsWhere(req.LastId != 0, "id > ?", req.LastId), opts.Limit(consts.DefaultLimit))
	if err != nil {
		return nil, err
	}
	// 根据list返回 id list
	commentIds := lo.Map(list, func(t *model.Comment, i int) int64 { return t.ID })
	// 根据id list查找 sub comment
	subComments, err := item.subCommentRepo.GetList(ctx, opts.Where("comment_id in ?", commentIds), opts.Limit(50))
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

func (item *CommentLogic) GetSubList(ctx context.Context, req *request.SubCommentListReq) ([]*model.SubComment, error) {
	return item.subCommentRepo.GetList(ctx, opts.Where("comment_id = ? ", req.CommentId),
		opts.IsWhere(req.LastId != 0, "id > ?", req.LastId), opts.Limit(consts.DefaultLimit))
}

func NewCommentLogic(commentRepo *repo.CommentRepo, subCommentRepo *repo.SubCommentRepo) *CommentLogic {
	return &CommentLogic{commentRepo: commentRepo, subCommentRepo: subCommentRepo}
}
