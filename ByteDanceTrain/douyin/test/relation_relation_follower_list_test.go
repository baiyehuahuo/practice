package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestRelationFollowerListSuccess(t *testing.T) {
	relationRebuild(t)
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFollowerList, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.UserList) != 2 ||
		!checkUserEqual(body.UserList[0], RootUser) ||
		!checkUserEqual(body.UserList[1], TestUsers[1]) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowerListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	//data.Add("token", token)
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFollowerList, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowerListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "aa")
	data.Add("token", token)
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFollowerList, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationFollowerListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token[1:])
	body := &pb.DouyinRelationFollowerListResponse{}
	getResponse(t, data, constants.RouteRelationFollowerList, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.UserList) != 1 ||
		!checkUserEqual(body.UserList[0], constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
