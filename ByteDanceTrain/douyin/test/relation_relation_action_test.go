// POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"net/url"
	"strconv"
	"testing"
)

func TestRelationAction1Success(t *testing.T) {
	relationRebuild(t)
	toUserID, actionType := 1, 1 // 1 is root
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	followBody := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, followBody)
	if *followBody.StatusCode != constants.DefaultInt32 ||
		*followBody.StatusMsg != constants.DefaultString ||
		len(followBody.UserList) != 2 {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// 不可重复关注
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.DBCreateRelationEventError.ErrCode ||
		*body.StatusMsg != dyerror.DBCreateRelationEventError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationAction2Success(t *testing.T) {
	relationRebuild(t)
	toUserID, actionType := 3, 2
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	data := url.Values{}
	data.Add("user_id", strconv.Itoa(int(TestUserID1)))
	data.Add("token", token)
	followBody := &pb.DouyinRelationFollowListResponse{}
	getResponse(t, data, constants.RouteRelationFollowList, followBody)
	if *followBody.StatusCode != constants.DefaultInt32 ||
		*followBody.StatusMsg != constants.DefaultString ||
		len(followBody.UserList) != 0 {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// 不可取消关注
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.DBDeleteRelationEventError.ErrCode ||
		*body.StatusMsg != dyerror.DBDeleteRelationEventError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationActionParamsEmptyFail(t *testing.T) {
	actionType := 2
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	//_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationActionParamsInputTypeFail(t *testing.T) {
	actionType := 1
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", "AAA")
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationActionUnknownActionTypeFail(t *testing.T) {
	toUserID, actionType := 1, 3
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.ParamUnknownActionTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamUnknownActionTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestRelationActionAuthTokenFail(t *testing.T) {
	toUserID, actionType := 1, 1 // 1 is root
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token[1:])
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	body := &pb.DouyinRelationActionResponse{}
	postResponse(t, payload, writer, constants.RouteRelationAction, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
