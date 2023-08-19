// GET METHOD
package test

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"net/url"
	"strconv"
	"testing"
)

func TestFeedSuccessWithAllParams(t *testing.T) {
	data := url.Values{}
	data.Add("latest_time", strconv.Itoa(TestVideo2PublishTimestamp))
	data.Add("token", token)
	body := &pb.DouyinFeedResponse{}
	getResponse(t, data, constants.RouteFeed, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.VideoList) != len(TestVideos[1:]) { // reverse timestamp
		t.Fatalf("Test results are not as expected: %v %v %v", body.StatusCode, body.StatusMsg, len(body.VideoList))
	}
	for i := range body.VideoList {
		if !checkVideoEqual(body.VideoList[i], TestVideos[i+1]) {
			t.Fatalf("Test results are not as expected: %v %v ", body.VideoList[i], TestVideos[i+1])
		}
	}
}

func TestFeedSuccessWithoutParams(t *testing.T) {
	data := url.Values{}
	//data.Add("latest_time", strconv.Itoa(constants.TestVideo2PublishTimestamp))
	//data.Add("token", token)
	body := &pb.DouyinFeedResponse{}
	getResponse(t, data, constants.RouteFeed, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString ||
		len(body.VideoList) != len(TestVideos) {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	for i := range body.VideoList {
		if !checkVideoEqualWithoutFavorite(body.VideoList[i], TestVideos[i]) {
			t.Fatalf("Test results are not as expected: %d %v %v", i, body.VideoList[i], TestVideos[i])
		}
	}
}

func TestFeedParamsInputTypeFail(t *testing.T) {
	data := url.Values{}
	data.Add("latest_time", token)
	data.Add("token", token)
	body := &pb.DouyinFeedResponse{}
	getResponse(t, data, constants.RouteFeed, body)
	if *body.StatusCode != dyerror.ParamInputTypeError.ErrCode ||
		*body.StatusMsg != dyerror.ParamInputTypeError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) { // reverse timestamp
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestFeedAuthTokenFail(t *testing.T) {
	data := url.Values{}
	data.Add("latest_time", strconv.Itoa(TestVideo2PublishTimestamp))
	data.Add("token", token[1:])
	body := &pb.DouyinFeedResponse{}
	getResponse(t, data, constants.RouteFeed, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage ||
		len(body.VideoList) != 1 ||
		!checkVideoEqual(body.VideoList[0], constants.DefaultVideo) { // reverse timestamp
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
