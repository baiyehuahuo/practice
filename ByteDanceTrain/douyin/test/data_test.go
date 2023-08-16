package test

import "douyin/pb"

const (
	RootUserID              = int64(1) // 1 是 root
	RootUsername            = "root"
	RootUserPassword        = "rootpwd"
	RootUserFollowCount     = 1
	RootUserFollowerCount   = 0
	RootUserIsFollow        = false
	RootUserAvatar          = ""
	RootUserBackgroundImage = ""
	RootUserSignature       = ""
	RootUserTotalFavorited  = 0
	RootUserWorkCount       = 1
	RootUserFavoriteCount   = 2

	TestUserID1              = int64(2) // 1 是 root
	TestUsername1            = "fwf"
	TestUserPassword1        = "fwf233"
	TestUserFollowCount1     = 1
	TestUserFollowerCount1   = 2
	TestUserIsFollow1        = false
	TestUserAvatar1          = "/uploadfiles/fwf/avatar.png"
	TestUserBackgroundImage1 = "/uploadfiles/fwf/background.png"
	TestUserSignature1       = "hello world"
	TestUserTotalFavorited1  = 3
	TestUserWorkCount1       = 2
	TestUserFavoriteCount1   = 1

	TestUserID2              = int64(3) // 1 是 root
	TestUsername2            = "范伟锋"
	TestUserPassword2        = "baiyehuahuo"
	TestUserFollowCount2     = 1
	TestUserFollowerCount2   = 1
	TestUserIsFollow2        = true
	TestUserAvatar2          = "/uploadfiles/范伟锋/avatar.png"
	TestUserBackgroundImage2 = "/uploadfiles/范伟锋/background.png"
	TestUserSignature2       = "fail world"
	TestUserTotalFavorited2  = 0
	TestUserWorkCount2       = 0
	TestUserFavoriteCount2   = 0
)

const (
	TestVideo1ID               = int64(2)
	TestVideo1PlayURL          = "uploadfiles/fwf/抉择之战.war3"
	TestVideo1CoverURL         = "uploadfiles/fwf/抉择之战.png"
	TestVideo1FavoriteCount    = 1
	TestVideo1CommentCount     = 2
	TestVideo1IsFavorite       = false
	TestVideo1Title            = "抉择之战 记录视频"
	TestVideo1PublishTimestamp = 1691686968

	TestVideo2ID               = int64(3)
	TestVideo2PlayURL          = "uploadfiles/fwf/抉择之战.mp4"
	TestVideo2CoverURL         = "uploadfiles/fwf/抉择之战.jpg"
	TestVideo2FavoriteCount    = 2
	TestVideo2CommentCount     = 0
	TestVideo2IsFavorite       = true
	TestVideo2Title            = "抉择之战 游戏视频"
	TestVideo2PublishTimestamp = 1399548023
)

const (
	TestCommentID1         = int64(2)
	TestCommentContent1    = "抉择之战真不行"
	TestCommentCreateDate1 = "08-11"

	TestCommentID2         = int64(3)
	TestCommentContent2    = "抉择之战真垃圾"
	TestCommentCreateDate2 = "08-12"
)

const (
	TestMessageID1         = int64(2)
	TestMessageToUserID1   = int64(2)
	TestMessageFromUserID1 = int64(3)
	TestMessageContent1    = "开学不快乐"
	TestMessageCreateTime1 = "08-15"

	TestMessageID2         = int64(3)
	TestMessageToUserID2   = int64(3)
	TestMessageFromUserID2 = int64(2)
	TestMessageContent2    = "早安"
	TestMessageCreateTime2 = "08-15"

	TestMessageID3         = int64(4)
	TestMessageToUserID3   = int64(2)
	TestMessageFromUserID3 = int64(3)
	TestMessageContent3    = "午安"
	TestMessageCreateTime3 = "08-15"
)

var (
	RootUser              *pb.User
	TestUser              *pb.User
	TestUsers             []*pb.User
	TestUserFavoriteVideo *pb.Video
	TestVideos            []*pb.Video
	TestComments          []*pb.Comment
	TestMessages          []*pb.Message
)

func init() {
	// users
	RootUser = &pb.User{
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
	*RootUser.Id = RootUserID // 1 是 root
	*RootUser.Name = RootUsername
	*RootUser.FollowCount = RootUserFollowCount
	*RootUser.FollowerCount = RootUserFollowerCount
	*RootUser.IsFollow = RootUserIsFollow
	*RootUser.Avatar = RootUserAvatar
	*RootUser.BackgroundImage = RootUserBackgroundImage
	*RootUser.Signature = RootUserSignature
	*RootUser.TotalFavorited = RootUserTotalFavorited
	*RootUser.WorkCount = RootUserWorkCount
	*RootUser.FavoriteCount = RootUserFavoriteCount

	TestUser = &pb.User{
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
	*TestUser.Id = TestUserID1 // 1 是 root
	*TestUser.Name = TestUsername1
	*TestUser.FollowCount = TestUserFollowCount1
	*TestUser.FollowerCount = TestUserFollowerCount1
	*TestUser.IsFollow = TestUserIsFollow1
	*TestUser.Avatar = TestUserAvatar1
	*TestUser.BackgroundImage = TestUserBackgroundImage1
	*TestUser.Signature = TestUserSignature1
	*TestUser.TotalFavorited = TestUserTotalFavorited1
	*TestUser.WorkCount = TestUserWorkCount1
	*TestUser.FavoriteCount = TestUserFavoriteCount1

	TestUser2 := &pb.User{
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
	*TestUser2.Id = TestUserID2 // 1 是 root
	*TestUser2.Name = TestUsername2
	*TestUser2.FollowCount = TestUserFollowCount2
	*TestUser2.FollowerCount = TestUserFollowerCount2
	*TestUser2.IsFollow = TestUserIsFollow2
	*TestUser2.Avatar = TestUserAvatar2
	*TestUser2.BackgroundImage = TestUserBackgroundImage2
	*TestUser2.Signature = TestUserSignature2
	*TestUser2.TotalFavorited = TestUserTotalFavorited2
	*TestUser2.WorkCount = TestUserWorkCount2
	*TestUser2.FavoriteCount = TestUserFavoriteCount2

	TestUsers = append(TestUsers, TestUser, TestUser2)

	// videos
	video := &pb.Video{
		Id:            new(int64),
		Author:        TestUser,
		PlayUrl:       new(string),
		CoverUrl:      new(string),
		FavoriteCount: new(int64),
		CommentCount:  new(int64),
		IsFavorite:    new(bool),
		Title:         new(string),
	}
	*video.Id = TestVideo1ID
	*video.PlayUrl = TestVideo1PlayURL
	*video.CoverUrl = TestVideo1CoverURL
	*video.FavoriteCount = TestVideo1FavoriteCount
	*video.CommentCount = TestVideo1CommentCount
	*video.IsFavorite = TestVideo1IsFavorite
	*video.Title = TestVideo1Title
	TestVideos = append(TestVideos, video)
	video = &pb.Video{
		Id:            new(int64),
		Author:        TestUser,
		PlayUrl:       new(string),
		CoverUrl:      new(string),
		FavoriteCount: new(int64),
		CommentCount:  new(int64),
		IsFavorite:    new(bool),
		Title:         new(string),
	}
	*video.Id = TestVideo2ID
	*video.PlayUrl = TestVideo2PlayURL
	*video.CoverUrl = TestVideo2CoverURL
	*video.FavoriteCount = TestVideo2FavoriteCount
	*video.CommentCount = TestVideo2CommentCount
	*video.IsFavorite = TestVideo2IsFavorite
	*video.Title = TestVideo2Title
	TestVideos = append(TestVideos, video)
	TestUserFavoriteVideo = TestVideos[1]

	// comments
	TestComment := &pb.Comment{
		Id:         new(int64),
		User:       TestUser,
		Content:    new(string),
		CreateDate: new(string),
	}
	*TestComment.Id = TestCommentID1
	*TestComment.Content = TestCommentContent1
	*TestComment.CreateDate = TestCommentCreateDate1
	TestComments = append(TestComments, TestComment)
	TestComment = &pb.Comment{
		Id:         new(int64),
		User:       TestUser,
		Content:    new(string),
		CreateDate: new(string),
	}
	*TestComment.Id = TestCommentID2
	*TestComment.Content = TestCommentContent2
	*TestComment.CreateDate = TestCommentCreateDate2
	TestComments = append(TestComments, TestComment)

	// messages
	TestMessage := &pb.Message{
		Id:         new(int64),
		ToUserId:   new(int64),
		FromUserId: new(int64),
		Content:    new(string),
		CreateTime: new(string),
	}
	*TestMessage.Id = TestMessageID1
	*TestMessage.ToUserId = TestMessageToUserID1
	*TestMessage.FromUserId = TestMessageFromUserID1
	*TestMessage.Content = TestMessageContent1
	*TestMessage.CreateTime = TestMessageCreateTime1
	TestMessages = append(TestMessages, TestMessage)

	TestMessage = &pb.Message{
		Id:         new(int64),
		ToUserId:   new(int64),
		FromUserId: new(int64),
		Content:    new(string),
		CreateTime: new(string),
	}
	*TestMessage.Id = TestMessageID2
	*TestMessage.ToUserId = TestMessageToUserID2
	*TestMessage.FromUserId = TestMessageFromUserID2
	*TestMessage.Content = TestMessageContent2
	*TestMessage.CreateTime = TestMessageCreateTime2
	TestMessages = append(TestMessages, TestMessage)

	TestMessage = &pb.Message{
		Id:         new(int64),
		ToUserId:   new(int64),
		FromUserId: new(int64),
		Content:    new(string),
		CreateTime: new(string),
	}
	*TestMessage.Id = TestMessageID3
	*TestMessage.ToUserId = TestMessageToUserID3
	*TestMessage.FromUserId = TestMessageFromUserID3
	*TestMessage.Content = TestMessageContent3
	*TestMessage.CreateTime = TestMessageCreateTime3
	TestMessages = append(TestMessages, TestMessage)
}
