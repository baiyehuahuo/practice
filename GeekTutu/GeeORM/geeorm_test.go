package GeeORM

import (
	"errors"
	"geeorm/log"
	"geeorm/session"
	"reflect"
	"testing"
)
import _ "github.com/mattn/go-sqlite3"

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	_, err := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if err != nil {
		t.Fatal(err)
	}

	log.Info("Before migrate:")
	var name string
	var age int
	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	for rows.Next() {
		err = rows.Scan(&name, &age)
		if err != nil {
			log.Info(err)
		}
		log.Info(name, age)
	}

	engine.Migrate(&User{})

	log.Info("After migrate:")
	rows, _ = s.Raw("SELECT * FROM User").QueryRows()
	columns, _ := rows.Columns()
	for rows.Next() {
		err = rows.Scan(&name, &age)
		if err != nil {
			log.Info(err)
		}
		log.Info(name, age)
	}
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("failed to migrate table user, got columns ", columns)
	}
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, errors.New("TestError") // 故意返回自定义错误 使事务回滚
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback", err, s.HasTable())
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if u.Name != "Tom" || u.Age != 18 || err != nil {
		t.Fatal("failed to commit ")
	}
}
