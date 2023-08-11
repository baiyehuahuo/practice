package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestPublishListSuccess(t *testing.T) {
	videoRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token)
	body := &pb.DouyinPublishListResponse{}
	getResponse(t, data, constants.RoutePublishList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.VideoList) != len(constants.TestVideos) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	for i := range body.VideoList {
		if !checkVideoEqual(body.VideoList[i], constants.TestVideos[i]) {
			t.Fatalf("Test result video %d are not as expected: %v", i, body.VideoList[i])
		}
	}
}

func TestPublishListParamsEmptyFail(t *testing.T) {
	videoRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	//data.Add("token", token)
	body := &pb.DouyinPublishListResponse{}
	getResponse(t, data, constants.RoutePublishList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishListParamsInputTypeFail(t *testing.T) {
	videoRebuild(t)
	data := url.Values{}
	data.Add("user_id", "jiliguala")
	data.Add("token", token)
	body := &pb.DouyinPublishListResponse{}
	getResponse(t, data, constants.RoutePublishList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishListAuthTokenFail(t *testing.T) {
	videoRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token[1:])
	body := &pb.DouyinPublishListResponse{}
	getResponse(t, data, constants.RoutePublishList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
