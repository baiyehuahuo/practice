package GeeORM

import (
	"database/sql"
	"fmt"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
	"strings"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Close database success")
}

// NewSession 返回 Session，共用同一个 DB
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}

type TxFunc func(session2 *session.Session) (interface{}, error)

func (e *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(err)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()

	return f(s)
}

// difference returns a - b
func difference(a []string, b []string) (diff []string) {
	hashB := make(map[string]struct{}, len(b))
	for _, s := range b {
		hashB[s] = struct{}{}
	}
	for _, s := range a {
		if _, ok := hashB[s]; !ok {
			diff = append(diff, s)
		}
	}
	return
}

// Migrate table 结构体变更时，数据库表字段自动迁移
func (e *Engine) Migrate(value interface{}) error {
	_, err := e.Transaction(func(s *session.Session) (result interface{}, err error) {
		if !s.Model(value).HasTable() { // 当前表不存在 不需要迁移 直接创建新的
			log.Infof("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}
		table := s.RefTable()
		rows, _ := s.Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		columns, _ := rows.Columns()                     // 查看字段信息
		addCols := difference(table.FieldNames, columns) // 查看新增的字段
		delCols := difference(columns, table.FieldNames) // 查看应该被删除的字段
		log.Infof("added cols %v, delete cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col) // 新增字段的执行语句
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table.Name, f.Name, f.Type)
			if _, err = s.Raw(sqlStr).Exec(); err != nil {
				return
			}
		}
		if len(delCols) == 0 {
			return
		}
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")
		// 创建新表 复制旧表内容
		s.Raw(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s;", tmp, fieldStr, table.Name))
		// 删除旧表
		s.Raw(fmt.Sprintf("DROP TABLE %s;", table.Name))
		// 重命名
		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmp, table.Name))
		return s.Exec()
	})
	return err
}
