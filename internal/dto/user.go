package dto

type RegisterReq struct {
	Name         string `json:"name"`
	IdentityType int32  `json:"identity_type"` // 登录类型
	Identifier   string `json:"identifier"`    // 标识
	Credential   string `json:"credential"`
}

type LoginReq struct {
	IdentityType int32  `json:"identity_type"` // 登录类型
	Identifier   string `json:"identifier"`    // 标识
	Credential   string `json:"credential"`
}

type UserReq struct {
	UserId int64 `uri:"user_id"`
}
