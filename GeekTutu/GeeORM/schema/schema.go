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
	FiledNames []string
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
		FiledNames: nil,
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
			schema.FiledNames = append(schema.FiledNames, field.Name)
			schema.fieldMap[field.Name] = field
		}
	}
	return schema
}
