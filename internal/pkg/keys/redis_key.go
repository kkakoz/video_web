package keys

import "fmt"

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
