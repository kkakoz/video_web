package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
	"video_web/pkg/timex"
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

func (commentLogic) Add(ctx context.Context, req *dto.CommentAdd) error {
	video, err := repo.Video().GetById(ctx, req.VideoId)
	if err != nil {
		return err
	}
	if video == nil {
		return errno.NewErr(404, 404, "未找到对应视频信息")
	}

	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	comment := &entity.Comment{
		UserId:       user.ID,
		Username:     user.Name,
		Avatar:       user.Avatar,
		Content:      req.Content,
		Top:          false,
		CommentCount: 0,
		LikeCount:    0,
		CreatedAt:    timex.Time{},
		UpdatedAt:    timex.Time{},
		SubComments:  nil,
	}
	if video.CollectionId == 0 {
		comment.TargetType = entity.CommentTargetTypeVideo
		comment.TargetId = video.ID
	} else {
		comment.TargetType = entity.CommentTargetTypeCollection
		comment.TargetId = video.CollectionId
	}

	return repo.Comment().Add(ctx, comment)
}

func (commentLogic) AddSub(ctx context.Context, req *dto.SubCommentAdd) error {

	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	subComment := &entity.SubComment{
		CommentId:        req.CommentId,
		RootSubCommentId: req.RootId,
		FromId:           user.ID,
		FromName:         user.Name,
		FromAvatar:       user.Avatar,
		ToId:             req.ToId,
		ToName:           req.ToName,
		Content:          req.Content,
		CreatedAt:        timex.Time{},
		UpdatedAt:        timex.Time{},
	}

	return repo.SubComment().Add(ctx, subComment)
}

// 查找评论和部分子评论
func (commentLogic) GetList(ctx context.Context, req *dto.CommentList) ([]*entity.Comment, error) {

	video, err := repo.Video().GetById(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, errno.NewErr(404, 404, "对应视频信息未找到")
	}

	var list []*entity.Comment
	if video.CollectionId == 0 { // 不是合集
		list, err = repo.Comment().GetList(ctx, opt.Where("target_id = ? and target_type = ?", video.ID, entity.CommentTargetTypeVideo),
			opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
	} else { // 是合集
		list, err = repo.Comment().GetList(ctx, opt.Where("target_id = ? and target_type = ?", video.CollectionId, entity.CommentTargetTypeCollection),
			opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
	}
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

func (commentLogic) GetSubList(ctx context.Context, req *dto.SubCommentList) ([]*entity.SubComment, error) {
	return repo.SubComment().GetList(ctx, opt.Where("comment_id = ? ", req.CommentId),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
}

func (commentLogic) Delete(ctx context.Context, req *dto.CommentDel) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo.Comment().DeleteById(ctx, req.CommentId)
	})
}

func (commentLogic) DeleteSubComment(ctx context.Context, req *dto.SubCommentDel) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo.SubComment().DeleteById(ctx, req.SubCommentId)
	})
}
