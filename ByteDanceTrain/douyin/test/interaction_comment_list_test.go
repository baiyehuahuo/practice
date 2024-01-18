package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestCommentListSuccess(t *testing.T) {
	data := url.Values{}
	data.Add("token", token)
	data.Add("video_id", strconv.Itoa(int(*TestVideos[0].Id)))
	body := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.CommentList) != len(TestComments) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	for i := range body.CommentList {
		if !checkCommentEqual(body.CommentList[i], TestComments[i]) {
			t.Fatalf("Test results are not as expected: %v", body)
		}
	}
}

func TestCommentListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("token", token)
	//data.Add("video_id", strconv.Itoa(int(*TestVideos[0].Id)))
	body := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.CommentList) != 1 ||
		!checkCommentEqual(body.CommentList[0], constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("token", token)
	data.Add("video_id", "aaa")
	body := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.CommentList) != 1 ||
		!checkCommentEqual(body.CommentList[0], constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestCommentListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("token", token[1:])
	data.Add("video_id", strconv.Itoa(int(*TestVideos[0].Id)))
	body := &pb.DouyinCommentListResponse{}
	getResponse(t, data, constants.RouteCommentList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.CommentList) != 1 ||
		!checkCommentEqual(body.CommentList[0], constants.DefaultComment) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
