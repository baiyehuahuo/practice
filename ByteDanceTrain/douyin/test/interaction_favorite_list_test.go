package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestFavoriteListSuccess(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	body := &pb.DouyinFavoriteListResponse{}
	getResponse(t, data, constants.RouteFavoriteList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], TestUserFavoriteVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestFavoriteListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	//data.Add("token", token)
	body := &pb.DouyinFavoriteListResponse{}
	getResponse(t, data, constants.RouteFavoriteList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestFavoriteListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "aaa")
	data.Add("token", token)
	body := &pb.DouyinFavoriteListResponse{}
	getResponse(t, data, constants.RouteFavoriteList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestFavoriteListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token[1:])
	body := &pb.DouyinFavoriteListResponse{}
	getResponse(t, data, constants.RouteFavoriteList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
