package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromID  uint   // sender
	ToID    uint   // receiver
	Type    int    // message type (group or one-to-one)
	Media   int    // message type (text or picture)
	Content string // message content
	Picture string
	Desc    string
	Amount  int // other
}

func (msg *Message) TableName() string {
	return "message"
}

func AutoMigrateMessage() error {
	return utils.GetDB().AutoMigrate(&Message{})
}
