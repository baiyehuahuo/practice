package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"douyin/service/VideoService"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

// ServePublishAction handle publish action request
// 登录用户选择视频上传
// Method is POST
// token, title and file is required.
func ServePublishAction(c *gin.Context) (res *pb.DouyinPublishActionResponse, dyerr *dyerror.DouyinError) {
	var (
		token, title string
		file         *multipart.FileHeader
	)
	if dyerr = checkPublishActionParams(c, &token, &title, &file); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
	}
	author := UserService.QueryUserByID(userID)
	filePath := path.Join(constants.UploadFileDir, author.Name, file.Filename)
	if _, err := os.Stat(filePath); err == nil {
		return nil, dyerror.UploadFileExistError
	}
	// todo 下面这两步应该用 事务 来执行的
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		dyerr = dyerror.UnknownError
		dyerr.ErrMessage = err.Error()
		return nil, dyerr
	}
	if err := VideoService.CreateVideo(&entity.Video{
		AuthorID: author.ID,
		PlayURL:  filePath,
		Title:    title,
	}); err != nil {
		return nil, dyerror.DBCreateVideoError
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
		//log.Printf("token: %v, title: %v, err: %v", token, title, err)
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
