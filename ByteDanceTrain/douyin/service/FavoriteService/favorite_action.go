package FavoriteService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
)

// CreateFavoriteEvent create a new record in the mysql database
func CreateFavoriteEvent(favorite *entity.Favorite) *dyerror.DouyinError {
	if err := DBService.GetDB().Create(favorite).Error; err != nil {
		return dyerror.DBCreateFavoriteEventError
	}
	return nil
}

// DeleteFavoriteEvent delete a favorite record in mysql
func DeleteFavoriteEvent(favorite *entity.Favorite) *dyerror.DouyinError {
	if err := DBService.GetDB().Delete(favorite).Error; err != nil {
		return dyerror.DBDeleteFavoriteEventError // todo maybe not succeed forever (delete returns ok whether the data exist or not)
	}
	return nil
}

// QueryFavoritesByUserID query publish videos where user_id == userID
func QueryFavoritesByUserID(userID int64) (favorites []*entity.Favorite) {
	DBService.GetDB().Where("user_id = ?", userID).Find(&favorites)
	return favorites
}

// QueryTotalFavoritedByAuthorID query favorited count by author id
func QueryTotalFavoritedByAuthorID(authorID int64) (totalFavorited int64) {
	DBService.GetDB().Model(&entity.Favorite{}).Where("author_id = ?", authorID).Count(&totalFavorited)
	return totalFavorited
}

// QueryFavoriteCountByUserID query favorited count by user id
func QueryFavoriteCountByUserID(userID int64) (favoriteCount int64) {
	DBService.GetDB().Model(&entity.Favorite{}).Where("user_id = ?", userID).Count(&favoriteCount)
	return favoriteCount
}

// QueryFavoriteCountByVideoID query favorited count by video id
func QueryFavoriteCountByVideoID(videoID int64) (favoriteCount int64) {
	DBService.GetDB().Model(&entity.Favorite{}).Where("video_id = ?", videoID).Count(&favoriteCount)
	return favoriteCount
}
