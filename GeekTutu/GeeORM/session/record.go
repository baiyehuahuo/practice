package session

import (
	"geeorm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	table := s.Model(values[0]).RefTable()                    // 获取模型
	s.clause.Set(clause.INSERT, table.Name, table.FieldNames) // 插入语句
	for _, value := range values {
		recordValues = append(recordValues, table.RecordValues(value)) // 插入的真实值
	}
	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)                              // 根据表名查表
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT) // 如果没有 WHERE ORDERBY LIMIT 也不会插入
	rows, err := s.Raw(sql, vars...).QueryRows()                                           // 执行查询语句
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err = rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest)) // 通过反射返回值到输入的 values 中
	}
	return rows.Close()
}
