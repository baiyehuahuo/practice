package interaction

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/CommentService"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeCommentList handle comment list request
// 查看视频的所有评论，按发布时间倒序
// 打开评论列表，立即调用douyin/comment/list/获取所有评论信息
// Method is GET
// token, video_id is required
func ServeCommentList(c *gin.Context) (res *pb.DouyinCommentListResponse, dyerr *dyerror.DouyinError) {
	var (
		token   string
		videoID int64
	)
	if dyerr = checkCommentListParams(c, &token, &videoID); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
	}
	commentLists := CommentService.QueryVideoCommentsByVideoID(videoID)
	pbCommentLists := make([]*pb.Comment, 0, len(commentLists))
	for i := range commentLists {
		user := common.ConvertToPBUser(UserService.QueryUserByID(commentLists[i].UserID))
		*user.IsFollow = RelationService.QueryFollowByIDs(userID, *user.Id)
		pbCommentLists = append(pbCommentLists, common.ConvertToPBComment(commentLists[i], user))
	}
	return &pb.DouyinCommentListResponse{
		StatusCode:  &constants.DefaultInt32,
		StatusMsg:   &constants.DefaultString,
		CommentList: pbCommentLists,
	}, nil
}

func checkCommentListParams(c *gin.Context, pToken *string, pVideoID *int64) *dyerror.DouyinError {
	token, videoID := c.Query("token"), c.Query("video_id")
	if token == "" || videoID == "" {
		return dyerror.ParamEmptyError
	}
	id, err := strconv.Atoi(videoID)
	if err != nil {
		return dyerror.ParamInputTypeError
	}
	*pToken = token
	*pVideoID = int64(id)
	return nil
}
