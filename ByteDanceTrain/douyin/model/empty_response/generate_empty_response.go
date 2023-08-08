package empty_response

import (
	"douyin/constants"
	"douyin/pb"
)

func Feed() *pb.DouyinFeedResponse {
	res := &pb.DouyinFeedResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		VideoList:  []*pb.Video{constants.DefaultVideo},
		NextTime:   new(int64),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	*res.NextTime = constants.DefaultInt64
	return res
}

func PublishAction() *pb.DouyinPublishActionResponse {
	res := &pb.DouyinPublishActionResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func PublishList() *pb.DouyinPublishListResponse {
	res := &pb.DouyinPublishListResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		VideoList:  []*pb.Video{constants.DefaultVideo},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func UserLogin() *pb.DouyinUserLoginResponse {
	res := &pb.DouyinUserLoginResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserId:     new(int64),
		Token:      new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	*res.UserId = constants.DefaultInt64
	*res.Token = constants.DefaultString
	return res
}

func User() *pb.DouyinUserResponse {
	res := &pb.DouyinUserResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		User:       constants.DefaultUser,
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func UserRegister() *pb.DouyinUserRegisterResponse {
	res := &pb.DouyinUserRegisterResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserId:     new(int64),
		Token:      new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	*res.UserId = constants.DefaultInt64
	*res.Token = constants.DefaultString
	return res
}

func CommentAction() *pb.DouyinCommentActionResponse {
	res := &pb.DouyinCommentActionResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		Comment:    constants.DefaultComment,
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func CommentList() *pb.DouyinCommentListResponse {
	res := &pb.DouyinCommentListResponse{
		StatusCode:  new(int32),
		StatusMsg:   new(string),
		CommentList: []*pb.Comment{constants.DefaultComment},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func FavoriteAction() *pb.DouyinFavoriteActionResponse {
	res := &pb.DouyinFavoriteActionResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func FavoriteList() *pb.DouyinFavoriteListResponse {
	res := &pb.DouyinFavoriteListResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		VideoList:  []*pb.Video{constants.DefaultVideo},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func MessageAction() *pb.DouyinMessageActionResponse {
	res := &pb.DouyinMessageActionResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func MessageChat() *pb.DouyinMessageChatResponse {
	res := &pb.DouyinMessageChatResponse{
		StatusCode:  new(int32),
		StatusMsg:   new(string),
		MessageList: []*pb.Message{constants.DefaultMessage},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func RelationAction() *pb.DouyinRelationActionResponse {
	res := &pb.DouyinRelationActionResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func RelationFollowList() *pb.DouyinRelationFollowListResponse {
	res := &pb.DouyinRelationFollowListResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserList:   []*pb.User{constants.DefaultUser},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func RelationFollowerList() *pb.DouyinRelationFollowerListResponse {
	res := &pb.DouyinRelationFollowerListResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserList:   []*pb.User{constants.DefaultUser},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}

func RelationFriendList() *pb.DouyinRelationFriendListResponse {
	res := &pb.DouyinRelationFriendListResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserList:   []*pb.User{constants.DefaultUser},
	}
	*res.StatusCode = constants.DefaultInt32
	*res.StatusMsg = constants.DefaultString
	return res
}
