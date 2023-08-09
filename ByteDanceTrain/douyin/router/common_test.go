package router

import (
	"douyin/constants"
	"douyin/service/DBService"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	buf = make([]byte, 1e5)
	r   = gin.New()
	err error
)

func TestMain(m *testing.M) {
	SetupRouter(r)
	var err error
	if err = userRebuild(); err != nil {
		panic(err)
	}
	m.Run()
}

func execSQLFile(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	n, err := file.Read(buf)
	if err != nil {
		return err
	}

	sqlArr := strings.Split(string(buf[:n]), ";")
	db := DBService.GetDB()
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = db.Exec(sql).Error
		if err != nil {
			log.Println("数据库导入失败:" + err.Error())
			return err
		} else {
			log.Println(sql, "\t success!")
		}
	}
	return nil
}

func userRebuild() error {
	return execSQLFile(path.Join(constants.Assets, constants.UserSQLPath))
}
