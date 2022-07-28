package dto

type CommentAdd struct {
	TargetType uint8  `json:"target_type"`
	TargetId   int64  `json:"target_id"`
	Content    string `json:"content"`
}

type SubCommentAdd struct {
	CommentId int64  `uri:"comment_id"`
	ToId      int64  `json:"to_id"`
	ToName    string `json:"to_name"`
	ParentId  int64  `json:"parent_id"`
	Content   string `json:"content"`
}

type CommentList struct {
	TargetType uint8 `query:"target_type"`
	TargetId   int64 `query:"target_id"`
	LastId     int64 `query:"last_id"`
}

type SubCommentList struct {
	CommentId int64 `uri:"comment_id"`
	LastId    int64 `query:"last_id"`
}

type CommentDel struct {
	CommentId int64 `uri:"comment_id"`
}

type SubCommentDel struct {
	CommentId    int64 `uri:"comment_id"`
	SubCommentId int64 `uri:"sub_comment_id"`
}
