package basis

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"douyin/service/VideoService"
	"fmt"
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
	if dyerr = VideoService.CreateVideo(&entity.Video{
		AuthorID: author.ID,
		PlayURL:  filePath,
		Title:    title,
	}); dyerr != nil {
		return nil, dyerr
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkPublishActionParams(c *gin.Context, pToken, pTitle *string, pFile **multipart.FileHeader) *dyerror.DouyinError {
	body := struct {
		common.TokenAuthFields
		Title string `form:"title" json:"title"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", body)
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
