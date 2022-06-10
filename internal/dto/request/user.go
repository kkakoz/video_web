package request

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserReq struct {
	UserId int64 `uri:"user_id"`
}
