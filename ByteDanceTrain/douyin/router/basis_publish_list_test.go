package router

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strconv"
	"testing"
)

func getPublishListResponse(t *testing.T, data url.Values) (body *pb.DouyinPublishListResponse) {
	req, err := http.NewRequest(http.MethodGet, path.Join(constants.ProjectGroup, constants.RoutePublishList), nil)
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
	body = &pb.DouyinPublishListResponse{}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	return body
}

func TestPublishListSuccess(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token)
	body := getPublishListResponse(t, data)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.VideoList) != len(constants.TestVideos) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	for i := range body.VideoList {
		if !checkVideoEqual(body.VideoList[i], constants.TestVideos[i]) {
			t.Fatalf("Test result video %d are not as expected: %v", i, body.VideoList[i])
		}
	}
}

func TestPublishListParamsEmptyFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	//data.Add("token", token)
	body := getPublishListResponse(t, data)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishListParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", "jiliguala")
	data.Add("token", token)
	body := getPublishListResponse(t, data)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishListAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(constants.TestUserID)))
	data.Add("token", token[1:])
	body := getPublishListResponse(t, data)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
