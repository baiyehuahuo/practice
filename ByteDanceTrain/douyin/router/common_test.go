package router

import (
	"bytes"
	"douyin/constants"
	"douyin/pb"
	"douyin/service/DBService"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	buf   = make([]byte, 1e5)
	r     = gin.New()
	err   error
	token string
)

func TestMain(m *testing.M) {
	SetupRouter(r)
	var err error
	if err = userRebuild(); err != nil {
		panic(err)
	}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUsername)
	_ = writer.WriteField("password", constants.TestUserPassword)
	token = *getUserLoginResponse(nil, payload, writer).Token
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
		}
		//log.Println(sql, "\t success!")
	}
	return nil
}

func userRebuild() error {
	return execSQLFile(path.Join(constants.Assets, constants.UserSQLPath))
}

// if response user success and user info is test data return true
func checkUserInfoSuccess(user *pb.User) bool {
	return *user.Id == constants.TestUserID &&
		*user.Name == constants.TestUsername &&
		*user.FollowCount == constants.TestUserFollowCount &&
		*user.FollowerCount == constants.TestUserFollowerCount &&
		*user.IsFollow == constants.TestUserIsFollow &&
		*user.Avatar == constants.TestUserAvatar &&
		*user.BackgroundImage == constants.TestUserBackgroundImage &&
		*user.Signature == constants.TestUserSignature &&
		*user.TotalFavorited == constants.TestUserTotalFavorited &&
		*user.WorkCount == constants.TestUserWorkCount &&
		*user.FavoriteCount == constants.TestUserFavoriteCount
}

// if response user fail and user info is default return true
func checkUserInfoFail(user *pb.User) bool {
	return *user.Id == constants.DefaultInt64 &&
		*user.Name == constants.DefaultString &&
		*user.FollowCount == constants.DefaultInt64 &&
		*user.FollowerCount == constants.DefaultInt64 &&
		*user.IsFollow == constants.DefaultBool &&
		*user.Avatar == constants.DefaultString &&
		*user.BackgroundImage == constants.DefaultString &&
		*user.Signature == constants.DefaultString &&
		*user.TotalFavorited == constants.DefaultInt64 &&
		*user.WorkCount == constants.DefaultInt64 &&
		*user.FavoriteCount == constants.DefaultInt64
}
