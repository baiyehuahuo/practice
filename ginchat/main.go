package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:password@tcp(127.0.0.1:3306)/mashibing?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                 // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("connect db failed.")
	}

	user := &models.UserBasic{}
	db.First(user, 1) // 根据整型主键查找
	fmt.Println(user)
	user = &models.UserBasic{}
	db.First(user, "name = ?", "fwf") // 查找 name 字段值为 fwf 的记录
	fmt.Println(user)

	// Update - 将 password 更新为 8888888
	db.Model(user).Update("password", "88888888")
}
