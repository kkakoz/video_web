package entity

type Message struct {
	ID        int64  `json:"id"`
	FromId    int64  `json:"from_id"`
	ToUserId  int64  `json:"to_user_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
}

type Notice struct {
	ID int64 `json:"id"`
}
