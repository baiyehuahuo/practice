package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList
// @Summary      List Users
// @Description  get all user messages in database
// @Tags         用户服务
// @Success      200  {string}   json{"code", "data"}
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	var datas = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": datas,
	})
}
