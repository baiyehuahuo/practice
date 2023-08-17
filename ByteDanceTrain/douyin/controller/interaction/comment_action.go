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
	"log"
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
		common.TokenAuthFields
		common.VideoIDField
		common.ActionTypeField
		common.ContentFields
	}{}
	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", body)
	token, videoID, actionType := c.PostForm("token"), c.PostForm("video_id"), c.PostForm("action_type")
	commentText, commentID := c.PostForm("comment_text"), c.PostForm("comment_id")
	if token == "" || videoID == "" || actionType == "" {
		log.Printf("token: %v, videoID: %v, actionType: %v", token, videoID, actionType)
		return dyerror.ParamEmptyError
	}
	action, _ := strconv.Atoi(actionType)
	if action != 1 && action != 2 {
		return dyerror.ParamUnknownActionTypeError
	}
	if action == 1 && commentText == "" || action == 2 && commentID == "" {
		return dyerror.ParamEmptyError
	}
	vid, err1 := strconv.Atoi(videoID)
	cid, err2 := strconv.Atoi(commentID)
	if err1 != nil || action == 2 && err2 != nil {
		return dyerror.ParamInputTypeError
	}

	*pToken = token
	*pVideoID = int64(vid)
	*pAction = action
	*pCommentText = commentText
	*pCommentID = int64(cid)
	return nil
}
