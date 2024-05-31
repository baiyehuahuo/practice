package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql: ", viper.Get("mysql"))
	//fmt.Println(config)
}

var db *gorm.DB

func InitMySQL() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       viper.GetString("mysql.dns"),                        // DSN data source name
		DefaultStringSize:         viper.GetUint("mysql.default_string_size"),          // string 类型字段的默认长度
		DisableDatetimePrecision:  viper.GetBool("mysql.default_datetime_precision"),   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    viper.GetBool("mysql.dont_support_rename_index"),    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   viper.GetBool("mysql.dont_support_rename_column"),   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: viper.GetBool("mysql.skip_initialize_with_version"), // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("connect db failed.")
	}
}

func GetDB() *gorm.DB {
	return db
}
