package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var (
		code = http.StatusOK
		msg  = "create success"
		err  error
	)
	defer func() {
		c.JSON(code, gin.H{
			"message": msg,
		})
	}()

	input := struct {
		Name       string `form:"name" json:"name" binding:"required"`
		Password   string `form:"password" json:"password" binding:"required"`
		Repassword string `form:"repassword" json:"repassword" binding:"required"`
		Phone      string `form:"phone" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
		Email      string `form:"email" json:"email" valid:"email"`
	}{}

	if err = c.ShouldBind(&input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}
	if _, err = govalidator.ValidateStruct(input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}
	if input.Repassword != input.Password {
		code = http.StatusBadRequest
		msg = "password is not equals to repassword"
		return
	}

	if checkCode, checkMsg := checkExist(input.Name, input.Phone, input.Email); checkCode != http.StatusOK {
		code = checkCode
		msg = checkMsg
		return
	}

	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	user := &models.UserBasic{
		Name:          input.Name,
		Password:      utils.MakePassword(input.Password, salt),
		Salt:          salt,
		Phone:         input.Phone,
		Email:         input.Email,
		LoginTime:     time.Now(),
		HeartBeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}

	if err = models.CreateUser(*user).Error; err != nil {
		code = http.StatusInternalServerError
		msg = err.Error()
	}
}

// DeleteUser
// @Summary      Delete a user
// @Description  Delete a user from database
// @Tags         用户删除
// @param        id formData int true "用户id"
// @Success      200  {string}   json{"message"}
// @Router       /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	var (
		code = http.StatusOK
		msg  = "delete success"
		err  error
	)
	defer func() {
		c.JSON(code, gin.H{
			"message": msg,
		})
	}()

	input := struct {
		ID int `form:"id" json:"id" binding:"required"`
	}{}

	if err = c.ShouldBind(&input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}
	if _, err = govalidator.ValidateStruct(input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}

	user := models.UserBasic{}
	user.ID = uint(input.ID)
	if err = models.DeleteUser(user).Error; err != nil {
		code = http.StatusInternalServerError
		msg = err.Error()
	}
}

// UpdateUser
// @Summary      Update user message
// @Description  Update user message in database
// @Tags         更新用户信息
// @param        id formData int true "用户id"
// @param        name formData string false "用户名"
// @param        password formData string false "密码"
// @param        phone formData string false "电话号码"
// @param        email formData string false "邮箱"
// @Success      200  {string}   json{"message"}
// @Router       /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	var (
		code = http.StatusOK
		msg  = "update success"
		err  error
	)
	defer func() {
		c.JSON(code, gin.H{
			"message": msg,
		})
	}()

	input := struct {
		ID       int    `form:"id" json:"id" binding:"required"`
		Name     string `form:"name" json:"name"`
		Password string `form:"password" json:"password"`
		Phone    string `form:"phone" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
		Email    string `form:"email" json:"email" valid:"email"`
	}{}

	if err = c.ShouldBind(&input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}
	if _, err = govalidator.ValidateStruct(input); err != nil {
		code = http.StatusBadRequest
		msg = err.Error()
		return
	}

	if checkCode, checkMsg := checkExist(input.Name, input.Phone, input.Email); checkCode != http.StatusOK {
		code = checkCode
		msg = checkMsg
		return
	}

	user := &models.UserBasic{}
	user.ID = uint(input.ID)
	user.Name = input.Name
	user.Password = input.Password
	user.Phone = input.Phone
	user.Email = input.Email

	if err = models.UpdateUser(*user).Error; err != nil {
		code = http.StatusInternalServerError
		msg = err.Error()
	}
}

func checkExist(name, phone, email string) (statusCode int, msg string) {
	var user *models.UserBasic
	if name != "" {
		if user = models.FindUserByName(name); user.ID != 0 {
			return http.StatusBadRequest, "name is exists"
		}
	}
	if phone != "" {
		if user = models.FindUserByPhone(phone); user.ID != 0 {
			return http.StatusBadRequest, "phone is exists"
		}
	}
	if email != "" {
		if user = models.FindUserByEmail(email); user.ID != 0 {
			return http.StatusBadRequest, "email is exists"
		}
	}
	return http.StatusOK, ""
}
