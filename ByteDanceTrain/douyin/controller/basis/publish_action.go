package basis

import (
	"douyin/constants"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"path"
)

// ServePublishAction handle publish action request
// 登录用户选择视频上传
// Method is POST
// token, title and file is required.
func ServePublishAction(c *gin.Context) (res *pb.DouyinPublishActionResponse, err error) {
	var (
		token, title string
		file         *multipart.FileHeader
	)
	if err = checkPublishActionParams(c, &token, &title, &file); err != nil {
		return nil, err
	}
	if err = c.SaveUploadedFile(file, path.Join(constants.UploadFileDir, file.Filename)); err != nil {
		return nil, err
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkPublishActionParams(c *gin.Context, pToken, pTitle *string, pFile **multipart.FileHeader) error {
	token, title := c.PostForm("token"), c.PostForm("title")
	file, err := c.FormFile("file")
	if token == "" || title == "" || err == http.ErrMissingFile {
		log.Printf("token: %v, title: %v, err: %v", token, title, err)
		return constants.ParamEmptyError
	}
	if err != nil {
		return err
	}
	*pToken = token
	*pTitle = title
	*pFile = file
	return nil
}
