package dto

type CommentAdd struct {
	VideoId int64  `json:"video_id"`
	Content string `json:"content"`
}

type SubCommentAdd struct {
	CommentId int64  `uri:"comment_id"`
	ToId      int64  `json:"to_id"`
	ToName    string `json:"to_name"`
	RootId    int64  `json:"root_id"`
	Content   string `json:"content"`
}

type CommentList struct {
	VideoId int64 `json:"video_id"`
	LastId  int64 `query:"last_id"`
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
