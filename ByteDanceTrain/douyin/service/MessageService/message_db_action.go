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
