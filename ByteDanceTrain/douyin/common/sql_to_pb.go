package common

import (
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/CommentService"
	"douyin/service/FavoriteService"
	"douyin/service/RelationService"
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
	*pbUser.FollowCount = RelationService.QueryFollowCountByUserID(user.ID)
	*pbUser.FollowerCount = RelationService.QueryFollowerCountByUserID(user.ID)
	*pbUser.IsFollow = user.IsFollow // todo handle isFollow (remove from entity.User)
	*pbUser.Avatar = user.Avatar
	*pbUser.BackgroundImage = user.BackgroundImage
	*pbUser.Signature = user.Signature
	*pbUser.TotalFavorited = FavoriteService.QueryTotalFavoritedByAuthorID(user.ID)
	*pbUser.WorkCount = VideoService.QueryWorkCountByAuthorID(user.ID)
	*pbUser.FavoriteCount = FavoriteService.QueryFavoriteCountByUserID(user.ID)
	return pbUser
}

func ConvertToPBVideo(video *entity.Video, author *pb.User) *pb.Video {
	pbVideo := &pb.Video{
		Id:            new(int64),
		Author:        author,
		PlayUrl:       new(string),
		CoverUrl:      new(string),
		FavoriteCount: new(int64),
		CommentCount:  new(int64),
		IsFavorite:    new(bool),
		Title:         new(string),
	}
	*pbVideo.Id = video.ID
	*pbVideo.PlayUrl = video.PlayURL
	*pbVideo.CoverUrl = video.CoverURL
	*pbVideo.FavoriteCount = FavoriteService.QueryFavoriteCountByVideoID(video.ID)
	*pbVideo.CommentCount = CommentService.QueryCommentCountByVideoID(video.ID)
	*pbVideo.Title = video.Title
	return pbVideo
}

func ConvertToPBComment(comment *entity.Comment, user *pb.User) *pb.Comment {
	pbComment := &pb.Comment{
		Id:         new(int64),
		User:       user,
		Content:    new(string),
		CreateDate: new(string),
	}
	*pbComment.Id = comment.ID
	*pbComment.Content = comment.Content
	*pbComment.CreateDate = comment.CreateDate
	return pbComment
}

func ConvertToPBMessage(msg *entity.Message) *pb.Message {
	pbMessage := &pb.Message{
		Id:         new(int64),
		ToUserId:   new(int64),
		FromUserId: new(int64),
		Content:    new(string),
		CreateTime: new(string),
	}
	*pbMessage.Id = msg.ID
	*pbMessage.ToUserId = msg.ToUserID
	*pbMessage.FromUserId = msg.FromUserID
	*pbMessage.Content = msg.Content
	*pbMessage.CreateTime = msg.CreateTime
	return pbMessage
}
