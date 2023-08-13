// POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"testing"
)

func TestUserLoginSuccess(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	_ = writer.WriteField("password", TestUserPassword)
	body := &pb.DouyinUserLoginResponse{}
	postResponse(t, payload, writer, constants.RouteUserLogin, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.UserId != TestUserID ||
		len(*body.Token) != constants.TokenLength {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserLoginUsernameFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername[1:]) // 账号为 一位的话就会报错
	_ = writer.WriteField("password", TestUserPassword)
	body := &pb.DouyinUserLoginResponse{}
	postResponse(t, payload, writer, constants.RouteUserLogin, body)
	if *body.StatusCode != dyerror.AuthUsernameOrPasswordFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthUsernameOrPasswordFailError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserLoginPasswordFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	_ = writer.WriteField("password", TestUserPassword[1:]) // 密码为 一位的话就会报错
	body := &pb.DouyinUserLoginResponse{}
	postResponse(t, payload, writer, constants.RouteUserLogin, body)
	if *body.StatusCode != dyerror.AuthUsernameOrPasswordFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthUsernameOrPasswordFailError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserLoginParamsEmptyFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	//_ = writer.WriteField("password", constants.TestUserPassword) // 密码缺失
	body := &pb.DouyinUserLoginResponse{}
	postResponse(t, payload, writer, constants.RouteUserLogin, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
