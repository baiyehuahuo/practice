// GET METHOD
package router

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestUserInfoSuccess(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token)
	body := &pb.DouyinUserResponse{}
	getResponse(t, data, constants.RouteUserInfo, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		!checkUserEqual(body.User, constants.TestUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserInfoParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	//data.Add("token", token)
	body := &pb.DouyinUserResponse{}
	getResponse(t, data, constants.RouteUserInfo, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		!checkUserEqual(body.User, constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserInfoInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "hello")
	data.Add("token", token)
	body := &pb.DouyinUserResponse{}
	getResponse(t, data, constants.RouteUserInfo, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		!checkUserEqual(body.User, constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserInfoTokenAuthFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token[1:])
	body := &pb.DouyinUserResponse{}
	getResponse(t, data, constants.RouteUserInfo, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		!checkUserEqual(body.User, constants.DefaultUser) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
