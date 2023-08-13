package test

import (
	"bytes"
	"douyin/constants"
	"douyin/pb"
	"douyin/router"
	"douyin/service/DBService"
	"encoding/json"
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
	userRebuild(nil)
	log.Print("users rebuild")
	videoRebuild(nil)
	log.Print("videos rebuild")
	favoriteRebuild(nil)
	log.Print("favorite events rebuild")
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", TestUsername)
	_ = writer.WriteField("password", TestUserPassword)
	body := &pb.DouyinUserLoginResponse{}
	postResponse(nil, payload, writer, constants.RouteUserLogin, body)
	token = *body.Token
	m.Run()
	userRebuild(nil)
	log.Print("users rebuild")
	videoRebuild(nil)
	log.Print("videos rebuild")
	favoriteRebuild(nil)
	log.Print("favorite events rebuild")
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
