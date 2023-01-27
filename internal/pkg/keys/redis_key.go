package keys

import (
	"fmt"
	"time"
	"video_web/internal/model/entity"
)

func TokenKey(token string) string {
	return "user:token:" + token
}

func LikeHashCache(targetType entity.TargetType, targetId int64) string {
	return fmt.Sprintf("like:%d:%d", targetType, targetId)
}

func LikeHashSync(targetType entity.TargetType) string {
	return fmt.Sprintf("likesync:%d", targetType)
}

func UserLikeKey() string {
	return "user:like"
}

func LikeValueKey() string {
	return fmt.Sprintf("like:value")
}

func UniqueVisitorKey(t time.Time) string {
	return fmt.Sprintf("uv:%s", t.Format("20060102"))
}

func UniqueVisitorRangeKey(start time.Time, end time.Time) string {
	return fmt.Sprintf("uv:%s:%s", start.Format("20060102"), end.Format("20060102"))
}

func DailyActiveUserKey(t time.Time) string {
	return fmt.Sprintf("dau:%s", t.Format("20060102"))
}

func DailyActiveUserRangeKey(start time.Time, end time.Time) string {
	return fmt.Sprintf("dau:%s:%s", start.Format("20060102"), end.Format("20060102"))
}

func CalculateVideoScoreKey() string {
	return "cal:video:score"
}

func UserActive(userId int64) string {
	return fmt.Sprintf("user:active:%d", userId)
}

func VideoViewIncrKey() string {
	return fmt.Sprintf("video:view")
}
