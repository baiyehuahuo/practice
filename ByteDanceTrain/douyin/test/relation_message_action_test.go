// POST METHOD
package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"mime/multipart"
	"strconv"
	"testing"
)

func TestMessageActionSuccess(t *testing.T) {
	messageRebuild(t)
	toUserID, actionType := 3, 1
	content := "开学完蛋"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	body := &pb.DouyinMessageActionResponse{}
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}

	// TODO 查看消息个数

	// 可以重复发送相同消息
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageActionParamsEmptyFail(t *testing.T) {
	toUserID, actionType := 3, 1
	content := "开学完蛋"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	//_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	body := &pb.DouyinMessageActionResponse{}
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageActionParamsInputTypeFail(t *testing.T) {
	toUserID, actionType := "a", 1
	content := "开学完蛋"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", toUserID)
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	body := &pb.DouyinMessageActionResponse{}
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageActionUnknownActionTypeFail(t *testing.T) {
	toUserID, actionType := 3, 2
	content := "开学完蛋"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	body := &pb.DouyinMessageActionResponse{}
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != dyerror.ParamUnknownActionTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamUnknownActionTypeError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageActionAuthTokenFail(t *testing.T) {
	toUserID, actionType := 3, 1
	content := "开学完蛋"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token[1:])
	_ = writer.WriteField("to_user_id", strconv.Itoa(toUserID))
	_ = writer.WriteField("action_type", strconv.Itoa(actionType))
	_ = writer.WriteField("content", content)
	body := &pb.DouyinMessageActionResponse{}
	postResponse(t, payload, writer, constants.RouteMessageAction, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
