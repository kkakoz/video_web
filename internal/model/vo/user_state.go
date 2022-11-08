package vo

type UserState struct {
	FollowedCreator bool `json:"followed_creator"`
	UserLike        bool `json:"user_like"`
	UserDisLike     bool `json:"user_dis_like"`
	UserCollect     bool `json:"user_collect"`
	UserShared      bool `json:"user_shared"`
}
