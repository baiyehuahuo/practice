package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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

var (
	db  *gorm.DB
	red *redis.Client
)

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)

	var err error
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	ip := viper.GetString("mysql.ip")
	port := viper.GetInt("mysql.port")
	dataset := viper.GetString("mysql.dataset")
	charset := viper.GetString("mysql.charset")
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, password, ip, port, dataset, charset),
		DefaultStringSize:         viper.GetUint("mysql.default_string_size"),          // string 类型字段的默认长度
		DisableDatetimePrecision:  viper.GetBool("mysql.default_datetime_precision"),   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    viper.GetBool("mysql.dont_support_rename_index"),    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   viper.GetBool("mysql.dont_support_rename_column"),   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: viper.GetBool("mysql.skip_initialize_with_version"), // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("connect db failed.", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func InitRedis() {
	ip := viper.GetString("redis.ip")
	port := viper.GetInt("redis.port")
	dbCode := viper.GetInt("redis.db")
	poolSize := viper.GetInt("redis.pool_size")
	minPoolConn := viper.GetInt("redis.min_pool_conn")
	red = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		Password:     "",
		DB:           dbCode,
		PoolSize:     poolSize,
		MinIdleConns: minPoolConn,
	})
	if err := red.Ping(context.Background()).Err(); err != nil {
		log.Fatal("connect redis failed.", err)
	}
}

func GetRedis() *redis.Client {
	return red
}

const (
	PublishKey = "websocket"
)

// Publish send message to redis
func Publish(ctx *gin.Context, channel string, msg string) error {
	if err := red.Publish(ctx, channel, msg).Err(); err != nil {
		return err
	}
	fmt.Println("Publish message to channel: ", channel, msg)
	return nil
}

// Subscribe scribe redis message
func Subscribe(ctx *gin.Context, channel string) (string, error) {
	sub := red.Subscribe(ctx, channel)
	defer func(sub *redis.PubSub) {
		err := sub.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(sub)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("subscribe from channel: ", msg)
	return msg.Payload, nil
}
