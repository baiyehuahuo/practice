package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestMessageChatSuccess(t *testing.T) {
	messageRebuild(t)
	toUserID := int64(3)
	data := url.Values{}
	data.Add("token", token)
	data.Add("to_user_id", strconv.Itoa(int(toUserID)))
	body := &pb.DouyinMessageChatResponse{}
	getResponse(t, data, constants.RouteMessageChat, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.MessageList) != len(TestMessages) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	for i := range body.MessageList {
		if !checkMessageEqual(body.MessageList[i], TestMessages[i]) {
			t.Fatalf("Test results are not as expected: %v", body)
		}
	}
}

func TestMessageChatParamsEmptyFail(t *testing.T) {
	//toUserID := int64(3)
	data := url.Values{}
	data.Add("token", token)
	//data.Add("to_user_id", strconv.Itoa(int(toUserID)))
	body := &pb.DouyinMessageChatResponse{}
	getResponse(t, data, constants.RouteMessageChat, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage ||
		len(body.MessageList) != 1 ||
		!checkMessageEqual(body.MessageList[0], constants.DefaultMessage) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageChatParamsInputTypeFail(t *testing.T) {
	//toUserID := int64(3)
	data := url.Values{}
	data.Add("token", token)
	data.Add("to_user_id", "aaa")
	body := &pb.DouyinMessageChatResponse{}
	getResponse(t, data, constants.RouteMessageChat, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.MessageList) != 1 ||
		!checkMessageEqual(body.MessageList[0], constants.DefaultMessage) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestMessageChatAuthTokenFail(t *testing.T) {
	toUserID := int64(3)
	data := url.Values{}
	data.Add("token", token[1:])
	data.Add("to_user_id", strconv.Itoa(int(toUserID)))
	body := &pb.DouyinMessageChatResponse{}
	getResponse(t, data, constants.RouteMessageChat, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.MessageList) != 1 ||
		!checkMessageEqual(body.MessageList[0], constants.DefaultMessage) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
