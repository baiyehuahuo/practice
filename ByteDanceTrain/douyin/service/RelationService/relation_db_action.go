package RelationService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
)

// CreateRelationEvent create a new record in the mysql database
func CreateRelationEvent(relation *entity.Relation) *dyerror.DouyinError {
	if err := DBService.GetDB().Create(relation).Error; err != nil {
		return dyerror.DBCreateRelationEventError
	}
	return nil
}

// DeleteRelationEvent delete a relation record in mysql
func DeleteRelationEvent(relation *entity.Relation) *dyerror.DouyinError {
	if affect := DBService.GetDB().Where("user_id = ? and to_user_id = ?", relation.UserID, relation.ToUserID).Delete(relation).RowsAffected; affect != 1 {
		return dyerror.DBDeleteRelationEventError
	}
	return nil
}
