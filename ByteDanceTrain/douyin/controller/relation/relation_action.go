package relation

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"github.com/gin-gonic/gin"
)

// ServeRelationAction handle relation action request
// 在Feed首页点击头像上的+号 和个人页点击关注，都会调用该接口执行关注和取消关注的逻辑
// Method is POST
// token, to_user_id, action_type is required
func ServeRelationAction(c *gin.Context) (res *pb.DouyinRelationActionResponse, err error) {
	var (
		token      string
		toUserID   int64
		actionType int
	)
	if err = checkRelationActionParams(c, &token, &toUserID, &actionType); err != nil {
		return nil, err
	}
	userID, err := TokenService.GetUserIDFromToken(token)
	if err != nil {
		return nil, err
	}
	relation := &entity.Relation{UserID: userID, ToUserID: toUserID}
	switch actionType {
	case 1:
		if err = RelationService.CreateRelationEvent(relation); err != nil {
			return nil, err
		}
	case 2:
		if err = RelationService.DeleteRelationEvent(relation); err != nil {
			return nil, err
		}
	}
	return &pb.DouyinRelationActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkRelationActionParams(c *gin.Context, pToken *string, pToUserID *int64, pActionType *int) error {
	body := struct {
		Token      string `form:"token" json:"token" binding:"required"`
		ToUserID   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
		ActionType int    `form:"action_type" json:"action_type" binding:"required,oneof=1 2"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pToken = body.Token
	*pToUserID = body.ToUserID
	*pActionType = body.ActionType
	return nil
}
