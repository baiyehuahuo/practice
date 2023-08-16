package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestRelationFriendListSuccess(t *testing.T) {
	relationRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	body := &pb.DouyinRelationFriendListResponse{}
	getResponse(t, data, constants.RouteRelationFriendList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], TestUsers[1]) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFriendListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	//data.Add("token", token)
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFriendList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFriendListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "aa")
	data.Add("token", token)
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFriendList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFriendListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token[1:])
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFriendList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
