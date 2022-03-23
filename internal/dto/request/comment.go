package request

type CommentAddReq struct {
	TargetType uint8  `json:"target_type"`
	TargetId   int64  `json:"target_id"`
	Content    string `json:"content"`
}

type SubCommentAddReq struct {
	CommentId int64  `uri:"comment_id"`
	ToId      int64  `json:"to_id"`
	ToName    string `json:"to_name"`
	ParentId  int64  `json:"parent_id"`
	Content   string `json:"content"`
}

type CommentListReq struct {
	TargetType uint8 `query:"target_type"`
	TargetId   int64 `query:"target_id"`
	LastId     int64 `query:"last_id"`
}

type SubCommentListReq struct {
	CommentId int64 `uri:"comment_id"`
	LastId    int64 `query:"last_id"`
}
