package router

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func getUserLoginResponse(t *testing.T, payload *bytes.Buffer, writer *multipart.Writer) *pb.DouyinUserLoginResponse {
	if err := writer.Close(); err != nil {
		t.Fatalf("Write params failed, err: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, path.Join(constants.ProjectGroup, constants.RouteUserLogin), payload)
	if err != nil {
		t.Fatalf("Build request failed, err: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	record := httptest.NewRecorder()
	r.ServeHTTP(record, req)
	res := record.Result()
	if res.StatusCode != 200 {
		t.Fatalf("Request status code is not as expected")
	}
	body := &pb.DouyinUserLoginResponse{}
	n, err := res.Body.Read(buf)
	if err != nil {
		t.Fatalf("Read respond body failed, err: %v", err)
	}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	return body
}

func TestUserLoginSuccess(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	_ = writer.WriteField("password", constants.TestUserPassword)
	body := getUserLoginResponse(t, payload, writer)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.UserId != constants.TestUserID ||
		len(*body.Token) != constants.TokenLength {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user login success passed successfully")
}

func TestUserLoginUsernameFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername[1:]) // 账号为 一位的话就会报错
	_ = writer.WriteField("password", constants.TestUserPassword)
	body := getUserLoginResponse(t, payload, writer)
	if *body.StatusCode != dyerror.AuthUsernameOrPasswordFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthUsernameOrPasswordFailError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user login username fail passed successfully")
}

func TestUserLoginPasswordFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	_ = writer.WriteField("password", constants.TestUserPassword[1:]) // 密码为 一位的话就会报错
	body := getUserLoginResponse(t, payload, writer)
	if *body.StatusCode != dyerror.AuthUsernameOrPasswordFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthUsernameOrPasswordFailError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user login password fail passed successfully")
}

func TestUserLoginParamsEmptyFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	//_ = writer.WriteField("password", constants.TestUserPassword) // 密码缺失
	body := getUserLoginResponse(t, payload, writer)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user login password empty passed successfully")
}
