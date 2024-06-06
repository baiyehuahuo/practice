package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
		"code": 0,
		"data": datas,
	})
}

// FindUserByNameAndPassword
// @Summary      Get User message by name and password
// @Description  get a user messages from database
// @Tags         用户服务
// @param        name formData string true "用户名"
// @param        password formData string true "密码"
// @Success      200  {string}   json{"code", "message", "data"}
// @Router       /user/findUserByNameAndPassword [post]
func FindUserByNameAndPassword(c *gin.Context) {
	var (
		code    = http.StatusOK
		msgCode = 0
		msg     = "get user failed"
		data    *models.UserBasic
		err     error
	)
	defer func() {
		if code != http.StatusOK {
			msgCode = -1
			data = nil
		}
		c.JSON(code, gin.H{
			"code":    msgCode,
			"message": msg,
			"data":    data,
		})
	}()

	input := struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
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
	user := models.FindUserByName(input.Name)
	if utils.ValidPassword(input.Password, user.Salt, user.Password) {
		msg = "get user success"
		data = models.FindUserByNameAndPwd(user.Name, user.Password)
	}
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
// @Success      200  {string}   json{"code", "message", "data"}
// @Router       /user/createUser [post]
func CreateUser(c *gin.Context) {
	var (
		code    = http.StatusOK
		msgCode = 0
		msg     = "create success"
		data    *models.UserBasic
		err     error
	)
	defer func() {
		if code != http.StatusOK {
			msgCode = -1
			data = nil
		}
		c.JSON(code, gin.H{
			"code":    msgCode,
			"message": msg,
			"data":    data,
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
	data = &models.UserBasic{
		Name:          input.Name,
		Password:      utils.MakePassword(input.Password, salt),
		Salt:          salt,
		Phone:         input.Phone,
		Email:         input.Email,
		LoginTime:     time.Now(),
		HeartBeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}

	if err = models.CreateUser(*data).Error; err != nil {
		code = http.StatusInternalServerError
		msg = err.Error()
	}
}

// DeleteUser
// @Summary      Delete a user
// @Description  Delete a user from database
// @Tags         用户删除
// @param        id formData int true "用户id"
// @Success      200  {string}   json{"code", "message", "data"}
// @Router       /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	var (
		code    = http.StatusOK
		msgCode = 0
		msg     = "delete success"
		data    *models.UserBasic
		err     error
	)
	defer func() {
		if code != http.StatusOK {
			msgCode = -1
			data = nil
		}
		c.JSON(code, gin.H{
			"code":    msgCode,
			"message": msg,
			"data":    data,
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

	data = models.FindUserByID(input.ID)
	if err = models.DeleteUser(*data).Error; err != nil {
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
// @Success      200  {string}   json{"code", "message", "data"}
// @Router       /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	var (
		code    = http.StatusOK
		msgCode = 0
		msg     = "update success"
		data    *models.UserBasic
		err     error
	)
	defer func() {
		if code != http.StatusOK {
			msgCode = -1
			data = nil
		}
		c.JSON(code, gin.H{
			"code":    msgCode,
			"message": msg,
			"data":    data,
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

	data = models.FindUserByID(input.ID)
	data.Name = input.Name
	data.Password = input.Password
	data.Phone = input.Phone
	data.Email = input.Email

	if err = models.UpdateUser(*data).Error; err != nil {
		code = http.StatusInternalServerError
		msg = err.Error()
	}
}

// SearchFriends
// @Summary      get user friends
// @Description  get all user friends message from database
// @Tags         从数据库中获取好友信息
// @param        id formData int true "用户id"
// @Success      200  {string}   json{"code", "message", "data"}
// @Router       /user/searchFriends [post]
func SearchFriends(c *gin.Context) {
	var (
		code = http.StatusOK
		data []*models.UserBasic
		err  error
	)
	defer func() {
		if code != http.StatusOK {
			data = nil
		}
		utils.RespOKList(c.Writer, data, len(data))
	}()

	input := struct {
		ID uint `form:"id" json:"id" binding:"required"`
	}{}

	if err = c.ShouldBind(&input); err != nil {
		code = http.StatusBadRequest
		return
	}
	if _, err = govalidator.ValidateStruct(input); err != nil {
		code = http.StatusBadRequest
		return
	}

	data = models.SearchFriends(input.ID)
}

// 防止跨域站点 伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", timeStr, msg)
		if err = ws.WriteMessage(1, []byte(m)); err != nil {
			fmt.Println(err)
			return
		}
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
