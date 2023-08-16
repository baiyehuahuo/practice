package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestRelationFollowListSuccess(t *testing.T) {
	relationRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	body := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], TestUsers[1]) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	//data.Add("token", token)
	body := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "aaa")
	data.Add("token", token)
	body := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token[1:])
	body := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
