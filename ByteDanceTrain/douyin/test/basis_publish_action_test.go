package test

import (
	"bytes"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const testUploadTitle = "test.title"
const testUploadFile = "/Users/weifengfan/Documents/Practice/ByteDanceTrain/douyin/assets/user.sql"
const testUploadExistFile = "/Users/weifengfan/Documents/Practice/ByteDanceTrain/douyin/assets/video.sql"

func TestPublishActionSuccess(t *testing.T) {
	videoRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("title", testUploadTitle)
	var (
		file *os.File
		part io.Writer
	)
	if file, err = os.Open(testUploadFile); err != nil {
		t.Fatalf("Test open file fail: %v", err)
	}
	defer file.Close()

	if part, err = writer.CreateFormFile("file", filepath.Base(testUploadFile)+strconv.Itoa(rand.Int())); err != nil { // 似乎是并发请求，同文件名会返回文件已存在的错误
		t.Fatalf("Test create form file fail: %v", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		t.Fatalf("Test copy file to post param fail: %v", err)
	}
	body := &pb.DouyinPublishActionResponse{}
	postResponse(t, payload, writer, constants.RoutePublishAction, body)
	if *body.StatusCode != constants.DefaultInt32 ||
		*body.StatusMsg != constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishActionTokenAuthFail(t *testing.T) {
	videoRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token[1:])
	_ = writer.WriteField("title", testUploadTitle)
	var (
		file *os.File
		part io.Writer
		err  error
	)
	if file, err = os.Open(testUploadFile); err != nil {
		t.Fatalf("Test open file fail: %v", err)
	}
	defer file.Close()

	if part, err = writer.CreateFormFile("file", filepath.Base(testUploadFile)); err != nil {
		t.Fatalf("Test create form file fail: %v", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		t.Fatalf("Test copy file to post param fail: %v", err)
	}

	body := &pb.DouyinPublishActionResponse{}
	postResponse(t, payload, writer, constants.RoutePublishAction, body)
	if *body.StatusCode != dyerror.AuthTokenFailError.ErrCode ||
		*body.StatusMsg != dyerror.AuthTokenFailError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishActionFileExistFail(t *testing.T) {
	videoRebuild(t)
	body := &pb.DouyinPublishActionResponse{}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("title", testUploadTitle)
	var (
		file *os.File
		part io.Writer
	)
	if file, err = os.Open(testUploadExistFile); err != nil {
		t.Fatalf("Test open file fail: %v", err)
	}
	defer file.Close()

	if part, err = writer.CreateFormFile("file", filepath.Base(testUploadExistFile)); err != nil {
		t.Fatalf("Test create form file fail: %v", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		t.Fatalf("Test copy file to post param fail: %v", err)
	}
	postResponse(t, payload, writer, constants.RoutePublishAction, body)
	if *body.StatusCode != dyerror.UploadFileExistError.ErrCode ||
		*body.StatusMsg != dyerror.UploadFileExistError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}

func TestPublishActionParamsEmptyFail(t *testing.T) {
	videoRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("title", testUploadTitle)
	body := &pb.DouyinPublishActionResponse{}
	postResponse(t, payload, writer, constants.RoutePublishAction, body)
	if *body.StatusCode != dyerror.ParamEmptyError.ErrCode ||
		*body.StatusMsg != dyerror.ParamEmptyError.ErrMessage {
		t.Fatalf("Test results are not as expected: %v", body)
	}
}
