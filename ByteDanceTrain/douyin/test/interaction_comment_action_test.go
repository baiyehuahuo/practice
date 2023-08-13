// POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestCommentAction1Success(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 2, 1
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
		*body.Comment.Id != *TestComments[1].Id+1 ||
		!checkUserEqual(body.Comment.User, TestUser) ||
		*body.Comment.Content != commentText ||
		*body.Comment.CreateDate != commentDate {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// 检测评论个数
	data := url.Values{}
	data.Add("token", token)
	data.Add("video_id", strconv.Itoa(int(*TestVideos[0].Id)))
	commentListBody := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, commentListBody)
	if *commentListBody.StatusCode != constants.DefaultInt32 ||
		*commentListBody.StatusMsg != constants.DefaultString ||
		len(commentListBody.CommentList) != len(TestComments)+1 {
		t.Fatalf("Test results are not as expected: %v", body)
	}

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
		*body.Comment.Id != *TestComments[1].Id+2 ||
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
	_ = writer.WriteField("comment_id", strconv.Itoa(int(*TestComments[1].Id)))
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		!checkCommentEqual(body.Comment, TestComments[1]) {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// 检测评论个数
	data := url.Values{}
	data.Add("token", token)
	data.Add("video_id", strconv.Itoa(int(*TestVideos[0].Id)))
	commentListBody := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, commentListBody)
	if *commentListBody.StatusCode != constants.DefaultInt32 ||
		*commentListBody.StatusMsg != constants.DefaultString ||
		len(commentListBody.CommentList) != len(TestComments)-1 {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_id", strconv.Itoa(int(*TestComments[1].Id)))
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.DBDeleteCommentEventError.ErrCode ||
		*body.StatusMsg != dyerror.DBDeleteCommentEventError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

// 检测越权可能性
func TestCommentAction2SafeSuccess(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 2, 2
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_id", strconv.Itoa(1)) // root's comment
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.DBDeleteCommentEventError.ErrCode ||
		*body.StatusMsg != dyerror.DBDeleteCommentEventError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionParamsEmptyFail(t *testing.T) {
	commentRebuild(t)
	videoID := 2
	//commentText, commentDate := "抉择之战 傻逼游戏", time.Now().Format("01-02")
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(1))
	//_ = writer.WriteField("comment_text", commentText)
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(2))
	//_ = writer.WriteField("comment_id", strconv.Itoa(int(*TestComment.Id)))
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionParamsInputTypeFail(t *testing.T) {
	commentRebuild(t)
	actionType := 1
	commentText := "抉择之战 傻逼游戏"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", "a")
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_text", commentText)
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionUnknownActionTypeFail(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 3, 3
	commentText := "抉择之战 傻逼游戏"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_text", commentText)
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.ParamUnknownActionTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamUnknownActionTypeError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentActionAuthTokenFail(t *testing.T) {
	commentRebuild(t)
	videoID, actionType := 3, 1
	commentText := "抉择之战 傻逼游戏"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token[1:])
	_ = writer.WriteField("video_id", strconv.Itoa(videoID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("comment_text", commentText)
	body := &pb.DouyinCommentActionResponse{}
	postResponse(t, payload, writer, constants.RouteCommentAction, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		!checkCommentEqual(body.Comment, constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
