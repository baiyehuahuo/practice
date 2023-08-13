package entity

type Comment struct {
	ID         int64  `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:评论 ID" json:"id"` // 视频评论id
	UserID     int64  `gorm:"column:user_id;type:int(11) unsigned;comment:用户 ID" json:"user_id"`                  // 评论用户信息
	VideoID    int64  `gorm:"column:video_id;type:int(11) unsigned;comment:视频 ID" json:"video_id"`
	Content    string `gorm:"column:content;type:varchar(100);comment:评论内容" json:"content"`                  // 评论内容
	CreateDate string `gorm:"column:create_date;type:varchar(5);comment:评论发布日期，格式 mm-dd" json:"create_date"` // 评论发布日期，格式 mm-dd
}

func (Comment) TableName() string {
	return "CommentEvents"
}
