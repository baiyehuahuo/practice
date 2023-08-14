package CommentService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
)

// CreateCommentEvent create a new record in the mysql database
func CreateCommentEvent(comment *entity.Comment) *dyerror.DouyinError {
	if err := DBService.GetDB().Create(comment).Error; err != nil {
		return dyerror.DBCreateCommentEventError
	}
	return nil
}

// DeleteCommentEvent delete a Comment record in mysql
func DeleteCommentEvent(comment *entity.Comment) *dyerror.DouyinError {
	// 小心越权删除
	if affect := DBService.GetDB().Where("user_id = ?", comment.UserID).Delete(comment).RowsAffected; affect != 1 {
		return dyerror.DBDeleteCommentEventError
	}
	return nil
}

// QueryCommentByIDLimitByIDs query comment by id, but userID and videoID is limit
func QueryCommentByIDLimitByIDs(commentID, userID, videoID int64) (comment *entity.Comment) {
	comment = &entity.Comment{}
	DBService.GetDB().Where("id = ? and user_id = ? and video_id = ?", commentID, userID, videoID).First(comment)
	return comment
}

// QueryVideoCommentsByVideoID query comments by videoID
func QueryVideoCommentsByVideoID(videoID int64) (comments []*entity.Comment) {
	DBService.GetDB().Where("video_id = ?", videoID).Find(&comments)
	return comments
}
