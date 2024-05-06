package service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mashibing/model"
)

var db *gorm.DB

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mashibing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connect mysql success")
}

func InsertCourse(course *model.Course) (int, error) {
	result := db.Create(course)
	return int(result.RowsAffected), result.Error
}

func DeleteCourse(course *model.Course) {
	db.Delete(course)
}

func UpdateCourse(course *model.Course) {
	db.Save(course)
}

func QueryCourse(course *model.Course) {
	db.First(course, course.Cid)
}

func QueryMultiCourse(course *model.Course) ([]*model.Course, error) {
	var ans []*model.Course
	db.Where("cid > ?", course.Cid).Find(&ans)
	return ans, nil
}
