package clause

import "strings"

type Type int

const (
	INSERT = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// Set 设置该 SQL 一行语句
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 根据顺序生成最终 SQL 语句
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
