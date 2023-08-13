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
)

func TestFavoriteAction1Success(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(2))
	_ = writer.WriteField("action_type", strconv.Itoa(1))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestFavoriteAction2Success(t *testing.T) {
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("video_id", strconv.Itoa(3))
	_ = writer.WriteField("action_type", strconv.Itoa(2))
	body := &pb.DouyinFavoriteActionResponse{}
	postResponse(t, payload, writer, constants.RouteFavoriteAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token)
	userBody := &pb.DouyinUserResponse{}
	getResponse(t, data, constants.RouteUserInfo, userBody)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*userBody.User.TotalFavorited != *constants.TestUser.TotalFavorited-1 ||
		*userBody.User.FavoriteCount != *constants.TestUser.FavoriteCount-1 {
		t.Fatalf("Test results are not as expected: %v", userBody)
	}
}

func TestFavoriteActionParamsEmptyFail(t *testing.T) {
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

func TestFavoriteActionParamsInputTypeFail(t *testing.T) {
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

func TestFavoriteActionUnknownActionTypeFail(t *testing.T) {
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

func TestFavoriteActionAuthTokenFail(t *testing.T) {
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
