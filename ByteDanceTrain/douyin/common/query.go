package common

type UserLoginFields struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type TokenAuthFields struct {
	UserID int64  `form:"user_id" json:"user_id"`
	Token  string `form:"token" json:"token"`
}

type ActionTypeField struct {
	ActionType int `form:"action_type" json:"action_type"`
}

type VideoIDField struct {
	VideoID int64 `form:"video_id" json:"video_id"`
}

type ToUserIDField struct {
	ToUserID int64 `form:"to_user_id" json:"to_user_id"`
}

type ContentFields struct {
	CommentID   int64  `form:"comment_id" json:"comment_id"`
	CommentText string `form:"comment_text" json:"comment_text"`
}
