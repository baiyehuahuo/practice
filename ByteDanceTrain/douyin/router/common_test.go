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
	log.Print("users rebuild")
	if err = videoRebuild(); err != nil {
		panic(err)
	}
	log.Print("videos rebuild")
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

func videoRebuild() error {
	return execSQLFile(path.Join(constants.Assets, constants.VideoSQLPath))
}

// if response user1 equals to user2 return true
func checkUserEqual(user1, user2 *pb.User) bool {
	return *user1.Id == *user2.Id &&
		*user1.Name == *user2.Name &&
		*user1.FollowCount == *user2.FollowCount &&
		*user1.FollowerCount == *user2.FollowerCount &&
		*user1.IsFollow == *user2.IsFollow &&
		*user1.Avatar == *user2.Avatar &&
		*user1.BackgroundImage == *user2.BackgroundImage &&
		*user1.Signature == *user2.Signature &&
		*user1.TotalFavorited == *user2.TotalFavorited &&
		*user1.WorkCount == *user2.WorkCount &&
		*user1.FavoriteCount == *user2.FavoriteCount
}

func checkVideoEqual(video1, video2 *pb.Video) bool {
	return *video1.Id == *video2.Id &&
		checkUserEqual(video1.Author, video2.Author) &&
		*video1.PlayUrl == *video2.PlayUrl &&
		*video1.CoverUrl == *video2.CoverUrl &&
		*video1.FavoriteCount == *video2.FavoriteCount &&
		*video1.CommentCount == *video2.CommentCount &&
		*video1.IsFavorite == *video2.IsFavorite &&
		*video1.Title == *video2.Title
}
