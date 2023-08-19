package MessageService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
	"time"
)

// CreateMessageEvent create a new record in the mysql database
func CreateMessageEvent(msg *entity.Message) error {
	if err := DBService.GetDB().Create(msg).Error; err != nil {
		return dyerror.DBCreateCommentEventError
	}
	return nil
}

// QueryMessagesByIDsAndTime query messages by fromUserID and toUserID
func QueryMessagesByIDsAndTime(fromUserID, toUserID int64, preMsgTime time.Time) (messages []*entity.Message) {
	DBService.GetDB().Where("(from_user_id = ? and to_user_id = ? or from_user_id = ? and to_user_id = ?) and create_time > ?", fromUserID, toUserID, toUserID, fromUserID, preMsgTime.Format("2006-01-02 15:04:05")).Find(&messages)
	return messages
}
