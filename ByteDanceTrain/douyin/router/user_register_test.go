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
	"strings"
	"testing"
)

func getUserRegisterResponse(t *testing.T, payload *bytes.Buffer, writer *multipart.Writer) (body *pb.DouyinUserRegisterResponse) {
	if err = writer.Close(); err != nil {
		t.Fatalf("Write params failed, err: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, path.Join(constants.ProjectGroup, constants.RouteUserRegister), payload)
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
	n, err := res.Body.Read(buf)
	if err != nil {
		t.Fatalf("Read respond body failed, err: %v", err)
	}
	body = &pb.DouyinUserRegisterResponse{}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	return body
}

func TestRegisterSuccess(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername+"2")
	_ = writer.WriteField("password", constants.TestUserPassword+"3")
	body := getUserRegisterResponse(t, payload, writer)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		*body.UserId != constants.TestUserID+1 ||
		len(*body.Token) != constants.TokenLength {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user register success passed successfully")
}

func TestRegisterParamsEmptyFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	body := getUserRegisterResponse(t, payload, writer)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test register params empty passed successfully")
}

func TestCreateUserFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	_ = writer.WriteField("password", constants.TestUserPassword)
	body := getUserRegisterResponse(t, payload, writer)
	if *body.StatusCode != dyerror.DBCreateUserError.ErrCode ||
		*body.StatusMsg != dyerror.DBCreateUserError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user register create user error passed successfully")
}

func TestRegisterParamInputLengthExceededFail(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", strings.Repeat(constants.TestUsername, 15))
	_ = writer.WriteField("password", constants.TestUserPassword)
	body := getUserRegisterResponse(t, payload, writer)
	if *body.StatusCode != dyerror.ParamInputLengthExceededError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputLengthExceededError.ErrMessage ||
		*body.UserId != constants.DefaultInt64 ||
		*body.Token != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test register input length exceeded error passed successfully")
}
