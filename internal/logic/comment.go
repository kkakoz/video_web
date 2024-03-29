package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/async/producer"
	"video_web/internal/consts"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
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

func (commentLogic) Add(ctx context.Context, req *dto.CommentAdd) (*entity.Comment, error) {

	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	comment := &entity.Comment{
		UserId:       user.ID,
		Username:     user.Name,
		Avatar:       user.Avatar,
		Content:      req.Content,
		Top:          false,
		CommentCount: 0,
		LikeCount:    0,
		TargetType:   entity.CommentTargetTypeVideo,
		TargetId:     req.TargetId,
		SubComments:  nil,
	}

	err = ormx.Transaction(ctx, func(txctx context.Context) error {
		err = repo.Comment().Add(txctx, comment)
		if err != nil {
			return err
		}
		err = repo.Video().Updates(txctx, map[string]any{
			"comment": gorm.Expr("comment + 1"),
		}, opt.Where("id = ?", req.TargetId))

		_ = Video().AddHots(txctx, req.TargetId)

		return producer.SendVideoEvent(&dto.Event{
			EventType:  dto.EventTypeComment,
			TargetId:   comment.ID,
			TargetType: entity.TargetTypeComment,
			ActorId:    user.ID,
		})
	})

	return comment, err
}

func (commentLogic) AddSub(ctx context.Context, req *dto.SubCommentAdd) (*entity.SubComment, error) {

	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	toUser, err := repo.User().GetById(ctx, req.ToId)
	if err != nil {
		return nil, err
	}
	toUserName := ""
	if toUser != nil {
		toUserName = toUser.Name
	}

	subComment := &entity.SubComment{
		CommentId:        req.CommentId,
		RootSubCommentId: req.RootId,
		UserId:           user.ID,
		Username:         user.Name,
		UserAvatar:       user.Avatar,
		ToId:             req.ToId,
		ToName:           toUserName,
		Content:          req.Content,
	}

	err = ormx.Transaction(ctx, func(txctx context.Context) error {

		err = repo.SubComment().Add(txctx, subComment)
		if err != nil {
			return err
		}
		comment, err := repo.Comment().GetById(txctx, subComment.CommentId)
		if err != nil {
			return err
		}

		if comment.TargetType == entity.CommentTargetTypeVideo {
			err = repo.Video().Updates(txctx, map[string]any{
				"comment": gorm.Expr("comment + 1"),
			}, opt.Where("id = ?", comment.TargetId))
			if err != nil {
				return err
			}

			_ = Video().AddHots(ctx, comment.TargetId)
		}

		return nil
	})

	return subComment, err
}

// GetList 查找评论和部分子评论
func (commentLogic) GetList(ctx context.Context, req *dto.CommentList) ([]*vo.Comment, error) {

	video, err := repo.Video().GetById(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, errno.NewErr(404, 404, "对应视频信息未找到")
	}

	opts := opt.NewOpts().Where("target_id = ? and target_type = ?", video.ID, entity.CommentTargetTypeVideo).Limit(consts.DefaultLimit).Order("created_at desc, id desc")
	if req.LastId != 0 {
		lastComment, err := repo.Comment().GetById(ctx, req.LastId)
		if err != nil {
			return nil, err
		}
		if lastComment == nil {
			return []*vo.Comment{}, nil
		}
		opts = opts.Where("created_at <= ? and id < ?", lastComment.CreatedAt, lastComment.ID)
	}
	var list []*entity.Comment
	list, err = repo.Comment().GetList(ctx, opts...)
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
			if len(subs) > 4 {
				comment.SubComments = subs[0:4]
			} else {
				comment.SubComments = subs
			}
		}
	})

	//IdLike := map[int64]*entity.Like{}
	//user, _ := local.GetUser(ctx)
	//if user != nil {
	//	likes, err := Like().Likes(ctx, &dto.UserLikeList{
	//		UserId:     user.ID,
	//		TargetIds:  commentIds,
	//		TargetType: entity.TargetTypeComment,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	IdLike = lox.Group(likes, func(value *entity.Like) int64 {
	//		return value.TargetId
	//	})
	//}

	res := make([]*vo.Comment, len(list))
	for i, v := range list {
		cur := vo.ConvertToComment(v)
		//if user != nil {
		//	like, ok := IdLike[v.ID]
		//	if ok {
		//		cur.Like = like.Like
		//	}
		//}
		res[i] = cur
	}

	return res, nil
}

func (commentLogic) GetSubList(ctx context.Context, req *dto.SubCommentList) ([]*vo.SubComment, error) {
	subComments, err := repo.SubComment().GetList(ctx, opt.Where("comment_id = ? ", req.CommentId),
		opt.IsWhere(req.LastId != 0, "id > ?", req.LastId), opt.Limit(consts.DefaultLimit))
	if err != nil {
		return nil, err
	}

	//subCommentIds := lo.Map(subComments, func(t *entity.SubComment, i int) int64 { return t.ID })

	//IdLike := map[int64]*entity.Like{}
	//user, _ := local.GetUser(ctx)
	//if user != nil {
	//	likes, err := Like().Likes(ctx, &dto.UserLikeList{
	//		UserId:     user.ID,
	//		TargetIds:  subCommentIds,
	//		TargetType: entity.TargetTypeSubComment,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	IdLike = lox.Group(likes, func(value *entity.Like) int64 {
	//		return value.TargetId
	//	})
	//}

	subs := make([]*vo.SubComment, len(subComments))
	for i, sub := range subComments {
		cur := vo.ConvertToSubComment(sub)
		//if user != nil {
		//	like, ok := IdLike[sub.ID]
		//	if ok {
		//		cur.Like = like.Like
		//	}
		//}
		subs[i] = cur
	}

	return subs, nil
}

func (commentLogic) Delete(ctx context.Context, req *dto.CommentDel) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	if user.Type == entity.UserTypeAdmin {
		return repo.SubComment().DeleteById(ctx, req.CommentId)
	}
	return repo.Comment().Delete(ctx, opt.Where("id = ? and user_id = ?", req.CommentId, user.ID))
}

func (commentLogic) DeleteSubComment(ctx context.Context, req *dto.SubCommentDel) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	if user.Type == entity.UserTypeAdmin {
		return repo.SubComment().DeleteById(ctx, req.SubCommentId)
	}
	return repo.SubComment().Delete(ctx, opt.Where("id = ? and user_id = ?", req.SubCommentId, user.ID))
}

func (l commentLogic) Get(ctx context.Context, id int64) (*entity.Comment, error) {
	return repo.Comment().GetById(ctx, id)

}

func (l commentLogic) GetSubComment(ctx context.Context, id int64) (*entity.SubComment, error) {
	return repo.SubComment().GetById(ctx, id)
}
