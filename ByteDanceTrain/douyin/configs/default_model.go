package configs

import "douyin/pb"

var (
	DefaultUser    *pb.User
	DefaultVideo   *pb.Video
	DefaultComment *pb.Comment
	DefaultMessage *pb.Message
)

func init() {
	DefaultUser = &pb.User{
		Id:              &DefaultInt64,
		Name:            &DefaultString,
		FollowCount:     &DefaultInt64,
		FollowerCount:   &DefaultInt64,
		IsFollow:        &DefaultBool,
		Avatar:          &DefaultString,
		BackgroundImage: &DefaultString,
		Signature:       &DefaultString,
		TotalFavorited:  &DefaultInt64,
		WorkCount:       &DefaultInt64,
		FavoriteCount:   &DefaultInt64,
	}
	DefaultVideo = &pb.Video{
		Id:            &DefaultInt64,
		Author:        DefaultUser,
		PlayUrl:       &DefaultString,
		CoverUrl:      &DefaultString,
		FavoriteCount: &DefaultInt64,
		CommentCount:  &DefaultInt64,
		IsFavorite:    &DefaultBool,
		Title:         &DefaultString,
	}
	DefaultComment = &pb.Comment{
		Id:         &DefaultInt64,
		User:       DefaultUser,
		Content:    &DefaultString,
		CreateDate: &DefaultString,
	}
	DefaultMessage = &pb.Message{
		Id:         &DefaultInt64,
		ToUserId:   &DefaultInt64,
		FromUserId: &DefaultInt64,
		Content:    &DefaultString,
		CreateTime: &DefaultString,
	}
}
