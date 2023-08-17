package query

type ParamsBody struct {
	// User Message
	UserID   int64  `form:"user_id" json:"user_id" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`

	ActionType int `form:"action_type" json:"action_type" binding:"required"`

	VideoID    int64  `form:"video_id" json:"video_id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	LatestTime int    `form:"latest_time" json:"latest_time"`
	Content    string `form:"content" json:"content" binding:"required"`

	ToUserID int64 `form:"to_user_id" json:"to_user_id" binding:"required"`

	CommentID   int64  `form:"comment_id" json:"comment_id"`
	CommentText string `form:"comment_text" json:"comment_text"`
}
