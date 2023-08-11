package entity

type Favorite struct {
	UserID   int64 `gorm:"column:user_id;type:int(11) unsigned;primary_key;comment:用户 ID" json:"user_id"`
	VideoID  int64 `gorm:"column:video_id;type:int(11) unsigned;primary_key;comment:用户昵称" json:"video_id"`
	AuthorID int64 `gorm:"column:author_id;type:int(11) unsigned;comment:视频作者 ID" json:"author_id"`
}

func (Favorite) TableName() string {
	return "FavoriteEvents"
}
