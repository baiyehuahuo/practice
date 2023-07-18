// Package schema convert object to table
package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string // 约束条件
}

// Schema represents a table of database
type Schema struct {
	Model      interface{} // 被映射的对象
	Name       string      // 表名
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// reflect TypeOf 和 ValueOf 返回入参的 类型 和 值。 Indirect 返回指向对象的指针
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:      dest,
		Name:       modelType.Name(),
		Fields:     nil,
		FieldNames: nil,
		fieldMap:   make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.fieldMap[field.Name] = field
		}
	}
	return schema
}

// RecordValues 根据数据库中列的顺序，从对象中找到对应的值，按顺序平铺
/*
	比如
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	拆解为 [Tom, 18, Sam, 25]
*/
func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
