package constants

import "fmt"

const (
	UploadFileDir = "uploadfiles"
)

const (
	// Database
	userName     = "root"
	password     = "rootpwd"
	ip           = "127.0.0.1"
	port         = 3306
	dbName       = "douyin"
	DataBaseType = "mysql"
)

var (
	DatabasePath = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, ip, port, dbName)
)
