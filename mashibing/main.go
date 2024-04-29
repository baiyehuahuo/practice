package main

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"mashibing/model"
	"mashibing/service"
	"net/http"
)

var adminUsers = gin.H{
	"fwf": gin.H{"email": "1770194225@163.com", "phone": 123321123},
	"xmy": gin.H{"email": "mfsnxy@qq.com", "phone": 66668888},
	"jhm": gin.H{"email": "none@fff.com", "phone": 88886666},
}

func getParams(ctx *gin.Context) {
	name := ctx.Param("username")
	age := ctx.Param("age")
	ctx.String(http.StatusOK, "Get Param Key name = %s, age = %s", name, age)
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func toRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func register(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.String(http.StatusOK, "get User %v", user)
}

func MyHandlerB() func(ctx *gin.Context) {
	counter := 0
	return func(ctx *gin.Context) {
		counter++
		pth := ctx.FullPath()
		method := ctx.Request.Method
		fmt.Printf("My handlerB ... called times: %d\tfull path: %s\tmethod: %s\n", counter, pth, method)
	}
}

func HandleSecret(ctx *gin.Context) {
	user := ctx.MustGet(gin.AuthUserKey).(string)
	if secret, ok := adminUsers[user]; ok {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret"})
	}
}

func HandleTestCookie(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		username = "fwf"
		ctx.SetCookie("username", username, 60*60, "/", "127.0.0.1", true, true)
		fmt.Println("set cookie")
	}
	ctx.String(http.StatusOK, "test cookie")
}

func HandleTestInsert(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBind(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	affected, err := service.InsertCourse(&course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Insert success. Affected: %d", affected))
}

func HandleTestDelete(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBind(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	affected, err := service.DeleteCourse(&course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Delete success. Affected: %d", affected))
}

func HandleTestUpdate(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBind(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	affected, err := service.UpdateCourse(&course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Update success. Affected: %d", affected))
}

func HandleTestQuery(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBind(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if _, err = service.QueryCourse(&course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Query success. Class name: %s  Tid: %d", course.Cname, course.Tid))
}

func HandleTestMultiQuery(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBind(&course)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	courses, err := service.QueryMultiCourse(&course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	var msg bytes.Buffer
	msg.WriteString("Query success.  ")
	for _, course := range courses {
		msg.WriteString(fmt.Sprintf("Class ID: %d  Class name: %s  Tid: %d  ", course.Cid, course.Cname, course.Tid))
	}
	ctx.JSON(http.StatusOK, msg.String())
}

func main() {
	r := gin.Default()  // 拿到一个 Engine 引擎 在New的基础上加入 Logger 与 Recovery 中间件
	r.Use(MyHandlerB()) // 与加入 engine 的顺序有关

	r.GET("/getParam/:username/:age", getParams)
	r.GET("/toRegister", toRegister)
	r.POST("/register", register)
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "assets")
	r.GET("/index", index)

	adminGroup := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"fwf": "xmy",
		"xmy": "jhm",
		"jhm": "fwf",
	}))
	adminGroup.GET("/secret", HandleSecret)

	r.GET("/testCookie", HandleTestCookie)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			err := session.Save()
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	r.POST("/testInsert", HandleTestInsert)
	r.POST("/testDelete", HandleTestDelete)
	r.POST("/testUpdate", HandleTestUpdate)
	r.POST("/testQuery", HandleTestQuery)
	r.POST("/testMultiQuery", HandleTestMultiQuery)

	if err := r.Run(); err != nil { // 开启服务 默认监听127.0.0.1:8080
		log.Fatal(err)
	}
}
