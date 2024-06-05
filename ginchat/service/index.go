package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
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

// Register
// @Tags 注册
// @Success 200 {string} register
// @Router /toRegister [get]
func Register(c *gin.Context) {
	temp, err := template.ParseFiles("views/user/register.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	//temp.Execute(c.Writer, "index")
	if err = temp.Execute(c.Writer, "register"); err != nil {
		panic(err)
	}
}
