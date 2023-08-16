package entity

type Message struct {
	ID         int64  `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:消息 ID" json:"id"`
	ToUserID   int64  `gorm:"column:to_user_id;type:int(11) unsigned;comment:消息接受者的 ID" json:"to_user_id"`
	FromUserID int64  `gorm:"column:from_user_id;type:int(11) unsigned;comment:消息发送者的 ID" json:"from_user_id"`
	Content    string `gorm:"column:content;type:varchar(100);comment:消息内容;NOT NULL" json:"content"`
	CreateTime string `gorm:"column:create_time;type:varchar(5);comment:消息创建时间;NOT NULL" json:"create_time"`
}

func (Message) TableName() string {
	return "MessageEvents"
}
