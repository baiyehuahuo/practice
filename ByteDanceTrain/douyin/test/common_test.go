package test

import (
	"bytes"
	"douyin/constants"
	"douyin/pb"
	"douyin/router"
	"douyin/service/DBService"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	router.SetupRouter(r)
	t := new(testing.T)
	userRebuild(t)
	videoRebuild(t)
	favoriteRebuild(t)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername1)
	_ = writer.WriteField("password", TestUserPassword1)
	body := &pb.DouyinUserLoginResponse{}
	postResponse(new(testing.T), payload, writer, constants.RouteUserLogin, body)
	token = *body.Token
	m.Run()
	userRebuild(t)
	videoRebuild(t)
	favoriteRebuild(t)
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

// if response user1 equals to user2 return true
func checkUserEqualWithoutFollow(user1, user2 *pb.User) bool {
	return *user1.Id == *user2.Id &&
		*user1.Name == *user2.Name &&
		*user1.FollowCount == *user2.FollowCount &&
		*user1.FollowerCount == *user2.FollowerCount &&
		//*user1.IsFollow == *user2.IsFollow &&
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
		*video1.IsFavorite == *video2.IsFavorite && // TODO  how to check it？
		*video1.Title == *video2.Title
}

func checkVideoEqualWithoutFavorite(video1, video2 *pb.Video) bool {
	return *video1.Id == *video2.Id &&
		checkUserEqualWithoutFollow(video1.Author, video2.Author) &&
		*video1.PlayUrl == *video2.PlayUrl &&
		*video1.CoverUrl == *video2.CoverUrl &&
		*video1.FavoriteCount == *video2.FavoriteCount &&
		*video1.CommentCount == *video2.CommentCount &&
		*video1.Title == *video2.Title
}

func checkCommentEqual(comment1, comment2 *pb.Comment) bool {
	return *comment1.Id == *comment2.Id &&
		checkUserEqual(comment1.User, comment2.User) &&
		*comment1.Content == *comment2.Content &&
		*comment1.CreateDate == *comment2.CreateDate
}

func checkMessageEqual(msg1, msg2 *pb.Message) bool {
	return *msg1.Id == *msg2.Id &&
		*msg1.ToUserId == *msg2.ToUserId &&
		*msg1.FromUserId == *msg2.FromUserId &&
		*msg1.Content == *msg2.Content &&
		*msg1.CreateTime == *msg2.CreateTime
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

func userRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.UserSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func videoRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.VideoSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func favoriteRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.FavoriteSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func commentRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.CommentSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func relationRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.RelationSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func messageRebuild(t *testing.T) {
	if err = execSQLFile(path.Join("../", constants.Assets, constants.MessageSQLFileName)); err != nil {
		t.Fatal(err)
	}
}

func getResponse(t *testing.T, data url.Values, routePath string, body interface{}) {
	req, err := http.NewRequest(http.MethodGet, path.Join(constants.ProjectGroup, routePath), nil)
	if err != nil {
		t.Fatalf("Build request failed, err: %v", err)
	}
	req.URL.RawQuery = data.Encode()
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	record := httptest.NewRecorder()
	r.ServeHTTP(record, req)
	res := record.Result()
	if res.StatusCode != 200 {
		t.Fatalf("Request status code is not as expected, code: %v", res.StatusCode)
	}
	n, err := res.Body.Read(buf)
	if err != nil {
		t.Fatalf("Read respond body failed, err: %v", err)
	}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
}

func postResponse(t *testing.T, payload *bytes.Buffer, writer *multipart.Writer, routePath string, body interface{}) {
	if err := writer.Close(); err != nil {
		t.Fatalf("Write params failed, err: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, path.Join(constants.ProjectGroup, routePath), payload)
	if err != nil {
		t.Fatalf("Build request failed, err: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	record := httptest.NewRecorder()
	r.ServeHTTP(record, req)
	res := record.Result()
	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode, path.Join(constants.ProjectGroup, routePath))
		t.Fatalf("Request status code is not as expected")
	}
	n, err := res.Body.Read(buf)
	if err != nil {
		t.Fatalf("Read respond body failed, err: %v", err)
	}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	return
}

/*
func TestSuccess(t *testing.T) {}

func TestParamsEmptyFail(t *testing.T) {}

func TestParamsInputTypeFail(t *testing.T) {}

func TestUnknownActionTypeFail(t *testing.T) {}

func TestAuthTokenFail(t *testing.T) {}
*/
