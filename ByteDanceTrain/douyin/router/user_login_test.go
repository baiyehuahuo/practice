package router

import (
	"bytes"
	"douyin/constants"
	"douyin/pb"
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	db  *gorm.DB
	buf = make([]byte, 1e5)
)

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

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(mysql.Open(constants.DatabasePath))
	if err != nil {
		panic(err)
	}
	if err = execSQLFile(path.Join(constants.Assets, constants.UserSQLPath)); err != nil {
		panic(err)
	}
	m.Run()
}

func TestLoginSuccess(t *testing.T) {
	r := SetupRouter()
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", constants.TestUserName)
	_ = writer.WriteField("password", constants.TestUserPassword)
	if err := writer.Close(); err != nil {
		t.Fatalf("Write params failed, err: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, path.Join(constants.ProjectGroup, constants.RouteUserLogin), payload)
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
	body := &pb.DouyinUserLoginResponse{}
	n, err := res.Body.Read(buf)
	if err != nil {
		t.Fatalf("Read respond body failed, err: %v", err)
	}
	if err = json.Unmarshal(buf[:n], body); err != nil {
		t.Fatalf("Convert respond body failed, err: %v", err)
	}
	if *body.StatusCode != constants.DefaultInt32 || *body.StatusMsg != constants.DefaultString || *body.UserId != constants.TestUserID || *body.Token == constants.DefaultString {
		t.Fatalf("Test results are not as expected: %v", body)
	}
	t.Log("Test login passed successfully")
}

//func TestLoginFail(t *testing.T) {
//
//	r := SetupRouter()
//
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	_ = writer.WriteField("username", "fwf")
//	if err := writer.Close(); err != nil {
//		t.Fatal(err)
//	}
//
//	client := &http.Client{}
//	req, err := http.NewRequest(http.MethodPost, "/douyin/user/login", payload)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//	res, err := client.Do(req)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(string(body))
//}
