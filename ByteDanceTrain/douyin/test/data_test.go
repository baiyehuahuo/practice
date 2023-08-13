package test

import "douyin/pb"

const (
	TestUserID              = int64(2) // 1 是 root
	TestUsername            = "fwf"
	TestUserPassword        = "fwf233"
	TestUserFollowCount     = 1
	TestUserFollowerCount   = 2
	TestUserIsFollow        = false
	TestUserAvatar          = "/uploadfiles/fwf/avatar.png"
	TestUserBackgroundImage = "/uploadfiles/fwf/background.png"
	TestUserSignature       = "hello world"
	TestUserTotalFavorited  = 3
	TestUserWorkCount       = 2
	TestUserFavoriteCount   = 1
)

const (
	TestVideo1ID               = int64(2)
	TestVideo1PlayURL          = "uploadfiles/fwf/抉择之战.war3"
	TestVideo1CoverURL         = "uploadfiles/fwf/抉择之战.png"
	TestVideo1FavoriteCount    = 1
	TestVideo1CommentCount     = 2
	TestVideo1IsFavorite       = true
	TestVideo1Title            = "抉择之战 记录视频"
	TestVideo1PublishTimestamp = 1691686968

	TestVideo2ID               = int64(3)
	TestVideo2PlayURL          = "uploadfiles/fwf/抉择之战.mp4"
	TestVideo2CoverURL         = "uploadfiles/fwf/抉择之战.jpg"
	TestVideo2FavoriteCount    = 2
	TestVideo2CommentCount     = 4
	TestVideo2IsFavorite       = false
	TestVideo2Title            = "抉择之战 游戏视频"
	TestVideo2PublishTimestamp = 1399548023
)

const (
	TestCommentID         = int64(2)
	TestCommentContent    = "抉择之战真垃圾"
	TestCommentCreateDate = "08-12"
)

var (
	TestUser              *pb.User
	TestVideos            []*pb.Video
	TestUserFavoriteVideo *pb.Video
	TestComment           *pb.Comment
)

func init() {
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
	*TestUser.Id = TestUserID // 1 是 root
	*TestUser.Name = TestUsername
	*TestUser.FollowCount = TestUserFollowCount
	*TestUser.FollowerCount = TestUserFollowerCount
	*TestUser.IsFollow = TestUserIsFollow
	*TestUser.Avatar = TestUserAvatar
	*TestUser.BackgroundImage = TestUserBackgroundImage
	*TestUser.Signature = TestUserSignature
	*TestUser.TotalFavorited = TestUserTotalFavorited
	*TestUser.WorkCount = TestUserWorkCount
	*TestUser.FavoriteCount = TestUserFavoriteCount

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

	TestComment = &pb.Comment{
		Id:         new(int64),
		User:       TestUser,
		Content:    new(string),
		CreateDate: new(string),
	}
	*TestComment.Id = TestCommentID
	*TestComment.Content = TestCommentContent
	*TestComment.CreateDate = TestCommentCreateDate
}
