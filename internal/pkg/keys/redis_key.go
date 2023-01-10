package keys

import (
	"fmt"
	"time"
)

func TokenKey(token string) string {
	return "user:token:" + token
}

func UserLikeField(userId int64, targetType uint8, targetId int64) string {
	return fmt.Sprintf("%d::%d::%d", userId, targetType, targetId)
}

func UserLikeKey() string {
	return "user:like"
}

func LikeValueKey(targetType uint8, targetId int64) string {
	return fmt.Sprintf("like:value:%d:%d", targetType, targetId)
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
