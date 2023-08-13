// USER REGISTER POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"strings"
	"testing"
)

func TestUserRegisterSuccess(t *testing.T) {
	userRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername+"2")
	_ = writer.WriteField("password", TestUserPassword+"3")
	body := &pb.DouyinUserRegisterResponse{}
	postResponse(t, payload, writer, constants.RouteUserRegister, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.UserId != TestUserID+1 ||
		len(*body.Token) != constants.TokenLength {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserRegisterParamsEmptyFail(t *testing.T) {
	userRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	body := &pb.DouyinUserRegisterResponse{}
	postResponse(t, payload, writer, constants.RouteUserRegister, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserRegisterCreateUserFail(t *testing.T) {
	userRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	_ = writer.WriteField("password", TestUserPassword)
	body := &pb.DouyinUserRegisterResponse{}
	postResponse(t, payload, writer, constants.RouteUserRegister, body)
	if *body.StatusCode != dyerror.DBCreateUserError.ErrCode ||
		*body.StatusMsg != dyerror.DBCreateUserError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestUserRegisterParamInputLengthExceededFail(t *testing.T) {
	userRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", strings.Repeat(TestUsername, 15))
	_ = writer.WriteField("password", TestUserPassword)
	body := &pb.DouyinUserRegisterResponse{}
	postResponse(t, payload, writer, constants.RouteUserRegister, body)
	if *body.StatusCode != dyerror.ParamInputLengthExceededError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputLengthExceededError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
