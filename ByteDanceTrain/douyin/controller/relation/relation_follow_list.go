package relation

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// ServeRelationFollowList handle relation follow list request
// 登录用户关注的所有用户列表
// 在个人页点击 粉丝/关注，能够打开该页面，会立即调用接口拉取关注用户和粉丝用户列表
// Method is GET
// user_id, token is required
func ServeRelationFollowList(c *gin.Context) (res *pb.DouyinRelationFollowListResponse, dyerr *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if dyerr = checkRelationFollowListParams(c, &userID, &token); dyerr != nil {
		return nil, dyerr
	}
	if dyerr = TokenService.CheckToken(token, userID); dyerr != nil {
		return nil, dyerr
	}
	relations := RelationService.QueryRelationEventByUserID(userID)
	pbUsers := make([]*pb.User, 0, len(relations))
	for i := range relations {
		user := common.ConvertToPBUser(UserService.QueryUserByID(relations[i].ToUserID))
		*user.IsFollow = true // RelationService.QueryFollowByIDs(userID, *user.Id) 本来就是关注目标
		pbUsers = append(pbUsers, user)
	}
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   pbUsers,
	}, nil
}

func checkRelationFollowListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
	body := struct {
		UserID int64  `form:"user_id" json:"user_id" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
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

	*pUserID = body.UserID
	*pToken = body.Token
	return nil
}
