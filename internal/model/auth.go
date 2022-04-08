package model

type Auth struct {
	ID           int64  `json:"id"`
	IdentityType int32  `json:"identity_type" gorm:"uniqueIndex:identifier_index"`                // 登录类型
	Identifier   string `json:"identifier" gorm:"type:varchar(255);uniqueIndex:identifier_index"` // 标识
	Credential   string `json:"credential" gorm:"type:varchar(255)"`
	Salt         string `json:"salt" gorm:"type:varchar(255)"`

	UserId    int64 `json:"user_id"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}
