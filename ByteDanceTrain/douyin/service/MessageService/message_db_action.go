package MessageService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
)

// CreateMessageEvent create a new record in the mysql database
func CreateMessageEvent(msg *entity.Message) *dyerror.DouyinError {
	if err := DBService.GetDB().Create(msg).Error; err != nil {
		return dyerror.DBCreateCommentEventError
	}
	return nil
}

// QueryMessagesByIDs query messages by fromUserID and toUserID
func QueryMessagesByIDs(fromUserID, toUserID int64) (messages []*entity.Message) {
	ids := []int64{fromUserID, toUserID}
	DBService.GetDB().Where("from_user_id IN ? and to_user_id IN ?", ids, ids).Find(&messages)
	return messages
}
