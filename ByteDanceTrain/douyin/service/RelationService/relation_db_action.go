package RelationService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
	"gorm.io/gorm"
)

// CreateRelationEvent create a new record in the mysql database
func CreateRelationEvent(relation *entity.Relation) error {
	if err := DBService.GetDB().Create(relation).Error; err != nil {
		return dyerror.DBCreateRelationEventError
	}
	return nil
}

// DeleteRelationEvent delete a relation record in mysql
func DeleteRelationEvent(relation *entity.Relation) error {
	if affect := DBService.GetDB().Where("user_id = ? and to_user_id = ?", relation.UserID, relation.ToUserID).Delete(relation).RowsAffected; affect != 1 {
		return dyerror.DBDeleteRelationEventError
	}
	return nil
}

// QueryRelationEventByUserID query relations by user_id
func QueryRelationEventByUserID(userID int64) (relations []*entity.Relation) {
	DBService.GetDB().Where("user_id = ?", userID).Find(&relations)
	return relations
}

// QueryRelationEventByToUserID query relations by to_user_id
func QueryRelationEventByToUserID(toUserID int64) (relations []*entity.Relation) {
	DBService.GetDB().Where("to_user_id = ?", toUserID).Find(&relations)
	return relations
}

// QueryFollowCountByUserID query follow count by user_id
func QueryFollowCountByUserID(userID int64) (followCount int64) {
	DBService.GetDB().Model(&entity.Relation{}).Where("user_id = ?", userID).Count(&followCount)
	return followCount
}

// QueryFollowerCountByUserID query follow count by user_id
func QueryFollowerCountByUserID(userID int64) (followCount int64) {
	DBService.GetDB().Model(&entity.Relation{}).Where("to_user_id = ?", userID).Count(&followCount)
	return followCount
}

// QueryFollowByIDs query follow exist between user_id and to_user_id
func QueryFollowByIDs(userID, toUserID int64) bool {
	return DBService.GetDB().Where("user_id = ? and to_user_id = ?", userID, toUserID).First(&entity.Relation{}).Error != gorm.ErrRecordNotFound
}
