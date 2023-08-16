package entity

import (
	"time"
)

// Video in mysql
// todo CommentCount, IsFavorite should be count by other table
type Video struct {
	ID           int64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:视频唯一标识" json:"id"`
	AuthorID     int64     `gorm:"column:author_id;type:int(11) unsigned;comment:视频作者id;NOT NULL" json:"author_id"`
	PlayURL      string    `gorm:"column:play_url;type:varchar(64);comment:视频播放地址;NOT NULL" json:"play_url"`
	CoverURL     string    `gorm:"column:cover_url;type:varchar(64);comment:视频封面地址" json:"cover_url"`
	CommentCount int64     `gorm:"column:comment_count;type:int(11) unsigned;default:0;comment:视频的评论总数" json:"comment_count"`
	IsFavorite   bool      `gorm:"column:is_favorite;type:tinyint(1);default:0;comment:true-已点赞，false-未点赞" json:"is_favorite"`
	Title        string    `gorm:"column:title;type:varchar(32);comment:视频标题" json:"title"`
	PublishTime  time.Time `gorm:"column:publish_time;type:timestamp;default:0;comment:发表时间" json:"publish_time"`
}

func (Video) TableName() string {
	return "Videos"
}
