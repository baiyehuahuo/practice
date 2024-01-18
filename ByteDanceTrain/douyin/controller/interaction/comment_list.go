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
)

// ServeCommentList handle comment list request
// 查看视频的所有评论，按发布时间倒序
// 打开评论列表，立即调用douyin/comment/list/获取所有评论信息
// Method is GET
// token, video_id is required
func ServeCommentList(c *gin.Context) (res *pb.DouyinCommentListResponse, err error) {
	var (
		token   string
		videoID int64
	)
	if err = checkCommentListParams(c, &token, &videoID); err != nil {
		return nil, err
	}
	userID, err := TokenService.GetUserIDFromToken(token)
	if err != nil {
		return nil, err
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

func checkCommentListParams(c *gin.Context, pToken *string, pVideoID *int64) error {
	body := struct {
		Token   string `form:"token" json:"token" binding:"required"`
		VideoID int64  `form:"video_id" json:"video_id" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pToken = body.Token
	*pVideoID = body.VideoID
	return nil
}
