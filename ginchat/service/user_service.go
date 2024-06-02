package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

// CreateUser
// @Summary      Create a new user
// @Description  insert a user data into database
// @Tags         用户注册
// @param        name formData string true "用户名"
// @param        password formData string true "密码"
// @param        repassword formData string true "确认密码"
// @param        email formData string false "邮箱"
// @param        phone formData string false "电话号码"
// @Success      200  {string}   json{"message"}
// @Router       /user/createUser [post]
func CreateUser(c *gin.Context) {
	var err error
	user := models.UserBasic{
		Name:          c.PostForm("name"),
		Password:      c.PostForm("password"),
		Phone:         c.PostForm("phone"),
		Email:         c.PostForm("email"),
		LoginTime:     time.Now(),
		HeartBeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}
	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name or password is empty",
		})
		return
	}

	repassword := c.PostForm("repassword")
	if repassword != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password is not equals to repassword",
		})
		return
	}

	var msg string = "create success"
	if err = models.CreateUser(user).Error; err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

// DeleteUser
// @Summary      Delete a user
// @Description  Delete a user from database
// @Tags         用户删除
// @param        id formData int true "用户id"
// @Success      200  {string}   json{"message"}
// @Router       /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	var err error
	id, err := strconv.ParseUint(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is empty",
		})
		return
	}
	user := models.UserBasic{}
	user.ID = uint(id)
	msg := "delete success"
	if err = models.DeleteUser(user).Error; err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
