package service

import (
	"fmt"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	w, req := c.Writer, c.Request
	srcFile, head, err := req.FormFile("file")
	if err != nil {
		log.Println("form file err:", err)
		utils.RespFail(w, err.Error())
		return
	}
	defer srcFile.Close()
	suffix := ".png"
	fileName := head.Filename
	tem := strings.Split(fileName, ".")
	if len(tem) > 1 {
		suffix = tem[len(tem)-1]
	}
	fileName = fmt.Sprintf("%d%04d.%s", time.Now().Unix(), rand.Int31(), suffix)
	savePath := "asset/upload/" + fileName
	dstFile, err := os.Create(savePath)
	if err != nil {
		log.Println("create file err:", err)
		utils.RespFail(w, err.Error())
		return
	}
	if _, err = io.Copy(dstFile, srcFile); err != nil {
		log.Println("copy file err:", err)
		utils.RespFail(w, err.Error())
		return
	}
	utils.RespOK(w, savePath, "发送图片成功")
}
