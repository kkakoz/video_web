package dto

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserId struct {
	UserId int64 `query:"user_id"`
}

type UpdateAvatar struct {
	Url string `json:"url" binding:"required"`
}

type UserActive struct {
	Code   string `query:"code" binding:"required"`
	UserId int64  `query:"user_id" binding:"required"`
}
