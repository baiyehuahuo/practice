package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
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
func ServePublishAction(c *gin.Context) (res *pb.DouyinPublishActionResponse, err *dyerror.DouyinError) {
	var (
		token, title string
		file         *multipart.FileHeader
	)
	if err = checkPublishActionParams(c, &token, &title, &file); err != nil {
		return nil, err
	}
	if saveErr := c.SaveUploadedFile(file, path.Join(constants.UploadFileDir, file.Filename)); saveErr != nil {
		err = dyerror.UnknownError
		err.ErrMessage = saveErr.Error()
		return nil, err
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkPublishActionParams(c *gin.Context, pToken, pTitle *string, pFile **multipart.FileHeader) *dyerror.DouyinError {
	token, title := c.PostForm("token"), c.PostForm("title")
	file, err := c.FormFile("file")
	if token == "" || title == "" || err == http.ErrMissingFile {
		log.Printf("token: %v, title: %v, err: %v", token, title, err)
		return dyerror.ParamEmptyError
	}
	if err != nil {
		dyerr := dyerror.UnknownError
		dyerr.ErrMessage = err.Error()
		return dyerr
	}
	*pToken = token
	*pTitle = title
	*pFile = file
	return nil
}
