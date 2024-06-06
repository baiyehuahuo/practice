package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	//c.HTML(http.StatusOK, "user/index.html", nil)
	temp, err := template.ParseFiles("views/user/index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	//temp.Execute(c.Writer, "index")
	if err = temp.Execute(c.Writer, "index"); err != nil {
		panic(err)
	}
}

// ToRegister
// @Tags 注册
// @Success 200 {string} register
// @Router /toRegister [get]
func ToRegister(c *gin.Context) {
	temp, err := template.ParseFiles("views/user/register.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	//temp.Execute(c.Writer, "index")
	if err = temp.Execute(c.Writer, "register"); err != nil {
		panic(err)
	}
}

// ToChat
// @Tags 交流页面
// @Success 200 {string} chat
// @Router /toChat [get]
func ToChat(c *gin.Context) {
	temp, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
	)
	if err != nil {
		panic(err)
	}
	uid, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		panic(err)
	}
	token := c.Query("token")
	user := &models.UserBasic{}
	user.ID = uint(uid)
	user.Identity = token
	if err = temp.Execute(c.Writer, user); err != nil {
		panic(err)
	}
}
