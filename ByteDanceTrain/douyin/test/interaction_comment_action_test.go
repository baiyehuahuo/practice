// POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"strconv"
	"testing"
	"time"
)

// todo 检测越权可能性

func TestCommentAction1Success(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 3, 1
	commentText, commentDate := "抉择之战 傻逼游戏", time.Now().Format("01-02")
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_text", commentText)
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.Comment.Id != *TestComment.Id+1 ||
		!checkUserEqual(body.Comment.User, TestUser) ||
		*body.Comment.Content != commentText ||
		*body.Comment.CreateDate != commentDate {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// todo 检测评论个数
	//data := url.Values{}
	//data.Add("user_id", strconv.Itoa(int(TestUserID)))
	//data.Add("token", token)
	//userBody := &pb.DouyinUserResponse{}
	//getResponse(t, data, constants.RouteUserInfo, userBody)
	//if *body.StatusCode != constants.DefaultInt32 ||
	//	*body.StatusMsg != constants.DefaultString ||
	//	*userBody.User.TotalFavorited != *TestUser.TotalFavorited+1 ||
	//	*userBody.User.FavoriteCount != *TestUser.FavoriteCount+1 {
	//	t.Fatalf("Test results are not as expected: %v", userBody)
	//}

	// 可重复评论
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_text", commentText)
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.Comment.Id != *TestComment.Id+2 ||
		!checkUserEqual(body.Comment.User, TestUser) ||
		*body.Comment.Content != commentText ||
		*body.Comment.CreateDate != commentDate {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentAction2Success(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 2, 2
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_id", strconv.Itoa(int(*TestComment.Id)))
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		!checkCommentEqual(body.Comment, TestComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// todo 检测评论个数
	//data := url.Values{}
	//data.Add("user_id", strconv.Itoa(int(TestUserID)))
	//data.Add("token", token)
	//userBody := &pb.DouyinUserResponse{}
	//getResponse(t, data, constants.RouteUserInfo, userBody)
	//if *body.StatusCode != constants.DefaultInt32 ||
	//	*body.StatusMsg != constants.DefaultString ||
	//	*userBody.User.TotalFavorited != *TestUser.TotalFavorited-1 ||
	//	*userBody.User.FavoriteCount != *TestUser.FavoriteCount-1 {
	//	t.Fatalf("Test results are not as expected: %v", userBody)
	//}

	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_id", strconv.Itoa(int(*TestComment.Id)))
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.DBDeleteCommentEventError.ErrCode ||
		*body.StatusMsg != dyerror.DBDeleteCommentEventError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionParamsEmptyFail(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(2))
	//_ = writer.WriteField("action_type", strconv.Itoa(1))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionParamsInputTypeFail(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", "a")
	_ = writer.WriteField("action_type", strconv.Itoa(1))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionUnknownActionTypeFail(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(2))
	_ = writer.WriteField("action_type", strconv.Itoa(3))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != dyerror.ParamUnknownActionTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamUnknownActionTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionAuthTokenFail(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token[1:])
	_ = writer.WriteField("video_id", strconv.Itoa(2))
	_ = writer.WriteField("action_type", strconv.Itoa(1))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
