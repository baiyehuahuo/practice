package query

type ParamsBody struct {
	// User Message
	UserID   int64  `form:"user_id" json:"user_id"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Token    string `form:"token" json:"token"`

	ActionType int `form:"action_type" json:"action_type"`

	VideoID    int64  `form:"video_id" json:"video_id"`
	Title      string `form:"title" json:"title"`
	LatestTime int    `form:"latest_time" json:"latest_time"`

	ToUserID int64 `form:"to_user_id" json:"to_user_id"`

	CommentID   int64  `form:"comment_id" json:"comment_id"`
	CommentText string `form:"comment_text" json:"comment_text"`

	Content string `form:"content" json:"content"`
}
