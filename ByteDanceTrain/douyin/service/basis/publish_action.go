package basis

import (
	"douyin/configs"
	"douyin/constants"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

// ServePublishAction handle publish action request
// 登录用户选择视频上传
// Method is POST
// token, title and file is required.
func ServePublishAction(c *gin.Context) (res *pb.DouyinPublishActionResponse, err error) {
	token, title := c.PostForm("token"), c.PostForm("title")
	file, err := c.FormFile("file")
	if token == "" || title == "" || err == http.ErrMissingFile {
		log.Printf("token: %v, title: %v, err: %v", token, title, err)
		return nil, configs.ParamEmptyError
	}
	if err != nil {
		return nil, err
	}
	if err = c.SaveUploadedFile(file, path.Join(constants.UploadFileDir, file.Filename)); err != nil {
		return nil, err
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}, nil
}
