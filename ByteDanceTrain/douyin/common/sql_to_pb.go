package common

import (
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/FavoriteService"
	"douyin/service/VideoService"
)

func ConvertToPBUser(user *entity.User) *pb.User {
	pbUser := &pb.User{
		Id:              new(int64),
		Name:            new(string),
		FollowCount:     new(int64),
		FollowerCount:   new(int64),
		IsFollow:        new(bool),
		Avatar:          new(string),
		BackgroundImage: new(string),
		Signature:       new(string),
		TotalFavorited:  new(int64),
		WorkCount:       new(int64),
		FavoriteCount:   new(int64),
	}
	*pbUser.Id = user.ID
	*pbUser.Name = user.Name
	*pbUser.FollowCount = user.FollowCount
	*pbUser.FollowerCount = user.FollowerCount
	*pbUser.IsFollow = user.IsFollow
	*pbUser.Avatar = user.Avatar
	*pbUser.BackgroundImage = user.BackgroundImage
	*pbUser.Signature = user.Signature
	*pbUser.TotalFavorited = FavoriteService.QueryTotalFavoritedByAuthorID(user.ID)
	*pbUser.WorkCount = VideoService.QueryWorkCountByAuthorID(user.ID)
	*pbUser.FavoriteCount = FavoriteService.QueryFavoriteCountByUserID(user.ID)
	return pbUser
}
