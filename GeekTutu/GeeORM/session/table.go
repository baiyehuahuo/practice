package session

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

// New create a instance for Session
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:       db,
		dialect:  dialect,
		refTable: nil,
		sql:      strings.Builder{},
		sqlVars:  nil,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = s.sqlVars[:0]
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session { // ??? 为什么要返回 Session 他不是调用者吗
	s.sql.WriteString(sql)
	s.sql.WriteByte(' ')
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 封装是为了打印日志和清空变量，以便执行多次 SQL

// Exec raw sql with sqlVals
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
