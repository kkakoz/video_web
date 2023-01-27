package dto

import "video_web/internal/model/entity"

type Event struct {
	EventType     EventType         `json:"event_type"`
	TargetId      int64             `json:"target_id"`
	TargetType    entity.TargetType `json:"target_type"`
	ActorId       int64             `json:"actor_id"`
	TargetOwnerId int64             `json:"target_owner_id"`
}

type EventType int8

const (
	EventTypeUserRegister EventType = iota + 1
	EventTypeMail
	EventTypeLike
	EventTypeAddVideo
	EventTypeDisLike
	EventTypeFollow
	EventTypeUnFollow
	EventTypeComment
	EventTypeCalculateHot
)

const EventKey = "event:video"
