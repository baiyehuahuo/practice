package entity

type Relation struct {
	UserID   int64 `gorm:"column:user_id;type:int(11) unsigned;primary_key;comment:用户 ID" json:"user_id"`
	ToUserID int64 `gorm:"column:to_user_id;type:int(11) unsigned;comment:被关注的用户 ID" json:"to_user_id"`
}

func (Relation) TableName() string {
	return "RelationEvents"
}
