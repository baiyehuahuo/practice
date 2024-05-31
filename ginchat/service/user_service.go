package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var datas = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": datas,
	})
}
