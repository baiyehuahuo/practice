package basis

import (
	"douyin/configs"
	"douyin/pb"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeUserInfo handle user request
// 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
// Method is GET
// user_id, token is required
func ServeUserInfo(c *gin.Context) (res *pb.DouyinUserResponse, err error) {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		fmt.Printf("user_id: %v, token: %v", userID, token)
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(userID); err != nil {
		fmt.Printf("user_id: %v, token: %v", userID, token)
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinUserResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		User:       configs.DefaultUser,
	}, nil
}
