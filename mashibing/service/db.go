package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mashibing/model"
)

var db *sql.DB

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mashibing"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect mysql success")
}

func InsertCourse(course *model.Course) (int, error) {
	stmt, err := db.Prepare("INSERT INTO course (Cname, Tid) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(course.Cname, course.Tid)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return int(affected), err
	}
	if err = stmt.Close(); err != nil {
		return int(affected), err
	}
	return int(affected), err
}

func DeleteCourse(course *model.Course) (int, error) {
	stmt, err := db.Prepare("DELETE FROM course WHERE Cid=?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(course.Cid)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return int(affected), err
	}
	if err = stmt.Close(); err != nil {
		return int(affected), err
	}
	return int(affected), err
}

func UpdateCourse(course *model.Course) (int, error) {
	stmt, err := db.Prepare("UPDATE course SET Cname=?, Tid=? WHERE Cid=?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(course.Cname, course.Tid, course.Cid)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return int(affected), err
	}
	if err = stmt.Close(); err != nil {
		return int(affected), err
	}
	return int(affected), err
}

func QueryCourse(course *model.Course) (int, error) {
	err := db.QueryRow("SELECT * FROM course WHERE Cid=?", course.Cid).Scan(&course.Cid, &course.Cname, &course.Tid)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func QueryMultiCourse(course *model.Course) ([]*model.Course, error) {
	rows, err := db.Query("SELECT * FROM course WHERE Cid>?", course.Cid)
	if err != nil {
		return nil, err
	}
	var courses []*model.Course
	for rows.Next() {
		course = &model.Course{}
		if err = rows.Scan(&course.Cid, &course.Cname, &course.Tid); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
