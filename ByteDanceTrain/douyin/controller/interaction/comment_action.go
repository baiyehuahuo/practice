package interaction

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/CommentService"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

// ServeCommentAction handle comment action request
// 登录用户对视频进行评论
// 提交评论会调用接口 /douyin/comment/action/
// Method is POST
// token, video_id, action_type is required
// comment_text comment_id is optional
func ServeCommentAction(c *gin.Context) (res *pb.DouyinCommentActionResponse, dyerr *dyerror.DouyinError) {
	var (
		token, commentText string
		videoID, commentID int64
		actionType         int
	)
	if dyerr = checkCommentActionParams(c, &token, &videoID, &actionType, &commentText, &commentID); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
	}

	var comment = &entity.Comment{VideoID: videoID}
	switch actionType {
	case 1:
		// 发表评论
		comment.UserID = userID
		comment.Content = commentText
		comment.CreateDate = time.Now().Format("01-02")
		if dyerr = CommentService.CreateCommentEvent(comment); dyerr != nil {
			return nil, dyerr
		}
	case 2:
		// 删除评论
		comment = CommentService.QueryCommentByIDLimitByIDs(commentID, userID, videoID)
		//if comment.ID == 0 {
		//	return nil, dyerror.DBDeleteCommentEventError
		//} todo 是不是加个错误未找到类型比较好？
		if dyerr = CommentService.DeleteCommentEvent(comment); dyerr != nil {
			return nil, dyerr
		}
	}
	user := common.ConvertToPBUser(UserService.QueryUserByID(comment.UserID))
	*user.IsFollow = RelationService.QueryFollowByIDs(userID, *user.Id)
	return &pb.DouyinCommentActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		Comment:    common.ConvertToPBComment(comment, user),
	}, nil
}

type queryCommentActionBody struct {
	Token       string `form:"token" json:"token"`
	VideoID     int64  `form:"video_id" json:"video_id"`
	ActionType  int    `form:"action_type" json:"action_type"`
	CommentText string `form:"comment_text" json:"comment_text"`
	CommentID   int64  `form:"comment_id" json:"comment_id"`
}

func checkCommentActionParams(c *gin.Context, pToken *string, pVideoID *int64, pAction *int, pCommentText *string, pCommentID *int64) *dyerror.DouyinError {
	body := struct {
		Token       string `form:"token" json:"token" binding:"required"`
		VideoID     int64  `form:"video_id" json:"video_id" binding:"required"`
		ActionType  int    `form:"action_type" json:"action_type" binding:"required"`
		CommentID   int64  `form:"comment_id" json:"comment_id"`
		CommentText string `form:"comment_text" json:"comment_text"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			return dyerror.ParamEmptyError
		case *strconv.NumError:
			return dyerror.ParamInputTypeError
		default:
			fmt.Printf("%T\n", err)
			dyerr := dyerror.UnknownError
			dyerr.ErrMessage = err.Error()
			return dyerr
		}
	}

	actionType, commentText, commentID := body.ActionType, body.CommentText, body.CommentID
	if actionType != 1 && actionType != 2 {
		return dyerror.ParamUnknownActionTypeError
	}
	if actionType == 1 && commentText == "" || actionType == 2 && commentID == 0 {
		return dyerror.ParamEmptyError
	}

	*pToken = body.Token
	*pVideoID = body.VideoID
	*pAction = actionType
	*pCommentText = commentText
	*pCommentID = commentID
	return nil
}
