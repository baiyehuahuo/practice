package main

import (
	"fmt"
	GeeORM "geeorm"
	"geeorm/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := GeeORM.NewEngine("sqlite3", "gee.db")
	if err != nil {
		log.Error(err)
		return
	}
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Sam").Exec()
	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Printf("Exec success, %d affected\n", count)
}
