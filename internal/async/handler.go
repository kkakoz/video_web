package async

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
)

type handler func(event *dto.Event) error

var handlerMap = map[dto.EventType]handler{}

func init() {
	handlerMap[dto.EventTypeLike] = handlerLike
	handlerMap[dto.EventTypeUserRegister] = handlerUserRegister
}

// 添加通知
func handlerLike(event *dto.Event) error {
	return ormx.Transaction(context.Background(), func(txctx context.Context) error {
		user, err := logic.User().GetUser(txctx, event.ActorId)
		if err != nil {
			return err
		}
		target := ""
		var userId int64 = 0
		switch event.TargetType {
		case entity.TargetTypeVideo:
			target = "视频"
			video, err := logic.Video().GetById(txctx, event.TargetId)
			if err != nil {
				return err
			}
			userId = video.UserId
		case entity.TargetTypeComment:
			target = "评论"
			comment, err := logic.Comment().Get(txctx, event.TargetId)
			if err != nil {
				return err
			}
			userId = comment.UserId
		case entity.TargetTypeSubComment:
			target = "评论"
			subComment, err := logic.Comment().GetSubComment(txctx, event.TargetId)
			if err != nil {
				return err
			}
			userId = subComment.UserId
		}
		notice := &entity.Notice{
			Content:    fmt.Sprintf("%s点赞了你的%s", user.Name, target),
			FromUserId: event.ActorId,
			IsRead:     false,
			UserId:     userId,
			TargetType: event.TargetType,
			TargetId:   event.TargetId,
		}
		return logic.Notice().Add(txctx, notice)
	})
}

// 初始化关注分组,收藏分组等
func handlerUserRegister(event *dto.Event) error {
	return ormx.Transaction(context.Background(), func(txctx context.Context) error {
		return logic.User().UserInit(txctx, event.ActorId)
	})
}
