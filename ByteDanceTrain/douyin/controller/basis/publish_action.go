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
// 发布视频，在本页完成信息填充后，点击发布就会调用该接口，发布完成会弹出发布完成提示信息并自动关闭本页面
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
	userID, err := TokenService.GetUserIDFromToken(token)
	if err != nil {
		return nil, err
	}
	author := UserService.QueryUserByID(userID)
	filePath := path.Join(constants.UploadFileDir, author.Name, file.Filename)
	if _, err = os.Stat(filePath); err == nil {
		return nil, dyerror.UploadFileExistError
	}
	// todo 下面这两步应该用 事务 来执行的
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		dyerr := dyerror.UnknownError
		dyerr.ErrMessage = err.Error()
		return nil, dyerr
	}
	if err = VideoService.CreateVideo(&entity.Video{
		AuthorID: author.ID,
		PlayURL:  filePath,
		Title:    title,
	}); err != nil {
		return nil, err
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkPublishActionParams(c *gin.Context, pToken, pTitle *string, pFile **multipart.FileHeader) error {
	body := struct {
		Token string `form:"token" json:"token" binding:"required"`
		Title string `form:"title" json:"title" binding:"required"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}
	file, err := c.FormFile("data")
	if err == http.ErrMissingFile {
		return dyerror.ParamEmptyError
	}
	if err != nil {
		dyerr := dyerror.UnknownError
		dyerr.ErrMessage = err.Error()
		return dyerr
	}
	*pToken = body.Token
	*pTitle = body.Title
	*pFile = file
	return nil
}
