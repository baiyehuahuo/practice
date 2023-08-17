package DBService

import (
	"douyin/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.DatabasePath), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				//SlowThreshold:             time.Microsecond, // Slow SQL threshold
				//ParameterizedQueries:      true,          // Don't include params in the SQL log
				LogLevel:                  logger.Error, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,         // Disable color
			}),
	})
	if err != nil {
		panic(err)
	}
	if err = db.Error; err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
