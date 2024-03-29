package dto

type CommentAdd struct {
	TargetId int64  `json:"target_id"`
	Content  string `json:"content"`
}

type SubCommentAdd struct {
	CommentId int64  `json:"comment_id"`
	ToId      int64  `json:"to_id"`
	RootId    int64  `json:"root_id"`
	Content   string `json:"content"`
}

type CommentList struct {
	VideoId int64 `json:"video_id" query:"video_id"`
	LastId  int64 `json:"last_id" query:"last_id"`
}

type SubCommentList struct {
	CommentId int64 `query:"comment_id"`
	LastId    int64 `query:"last_id"`
}

type CommentDel struct {
	CommentId int64 `json:"comment_id"`
}

type SubCommentDel struct {
	SubCommentId int64 `json:"sub_comment_id"`
}
