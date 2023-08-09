package router

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strconv"
	"testing"
)

func getUserInfoResponse(t *testing.T, data url.Values) (body *pb.DouyinUserResponse) {
	req, err := http.NewRequest(http.MethodGet, path.Join(constants.ProjectGroup, constants.RouteUserInfo), nil)
	if err != nil {
		t.Fatalf("Build request failed, err: %v", err)
	}
	req.URL.RawQuery = data.Encode()
	//req.Header.Set("Content-Type", writer.FormDataContentType())
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
	body = &pb.DouyinUserResponse{}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	return body
}

func TestUserInfoSuccess(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token)
	body := getUserInfoResponse(t, data)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		!checkUserInfoSuccess(body.User) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user message success passed successfully")
}

func TestUserInfoParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	//data.Add("token", token)
	body := getUserInfoResponse(t, data)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		!checkUserInfoFail(body.User) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user info params empty passed successfully")
}

func TestUserInfoInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "hello")
	data.Add("token", token)
	body := getUserInfoResponse(t, data)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		!checkUserInfoFail(body.User) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user info params type passed successfully")
}

func TestUserInfoTokenAuthFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token[1:])
	body := getUserInfoResponse(t, data)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		!checkUserInfoFail(body.User) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	log.Printf("Test user info token auth passed successfully")
}
