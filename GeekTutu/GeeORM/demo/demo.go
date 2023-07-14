package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // 注册 sqlite3 驱动
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err = db.Exec("DROP TABLE IF EXISTS User;"); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec("CREATE TABLE User(Name text);"); err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Sam")
	if err != nil {
		log.Fatal(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	log.Println(affected)
	row := db.QueryRow("SELECT Name From User LIMIT 1")
	var name string
	if err = row.Scan(&name); err != nil {
		log.Fatal(err)
	}
	log.Println(name)
}
