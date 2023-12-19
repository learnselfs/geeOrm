// Package clause @Author Bing
// @Date 2023/11/24 15:10:00
// @Desc
package clause

import (
	"github.com/learnselfs/geeOrm/utils"
	"strings"
)

func init() {
	RegisterClause("mysql", &ClauseMysql{})
}

var _ Clause = (*ClauseMysql)(nil)

type ClauseMysql struct {
	method         int
	table          string
	fields         []string
	values         []string
	limit          string
	whereFields    []string
	whereValues    []string
	whereCondition []string
	whereArgs      []string
	orderBy        string
	sql            strings.Builder
	args           []any
	// models
	insertModel interface{}
}

func (c *ClauseMysql) SetTable(table string) {
	c.table = table
}

func (c *ClauseMysql) Clear() {
	c.whereArgs = make([]string, 0)
	c.values = make([]string, 0)
	c.fields = make([]string, 0)
	c.sql.Reset()
	c.args = make([]any, 0)
}

func (c *ClauseMysql) Insert(s interface{}) Clause {
	c.method = utils.InsertMethod
	c.fields, c.values = utils.ParseStructFieldValueUnsafe(s)
	c.insertModel = s
	return c
}

func (c *ClauseMysql) Delete() Clause {
	c.method = utils.DeleteMethod
	return c
}

func (c *ClauseMysql) Select(fields []string) Clause {
	c.method = utils.SelectMethod
	c.fields = fields
	return c
}

func (c *ClauseMysql) Update(s interface{}) Clause {
	c.method = utils.UpdateMethod
	var fields []string
	var values []string
	fields, values = utils.ParseStructFieldValueUnsafe(s)
	c.fields = fields
	c.values = values
	return c
}

func (c *ClauseMysql) Where(whereFields string, whereValues string) Clause {
	c.whereFields = append(c.whereFields, whereFields)
	c.whereValues = append(c.whereValues, whereValues)
	c.condition(whereFields, "=", whereValues)
	return c
}

func (c *ClauseMysql) And(field, condition, value string) Clause {
	c.special("and", field, condition, value)
	return c
}

func (c *ClauseMysql) Not(field, condition, value string) Clause {
	c.special("not", field, condition, value)
	return c
}

func (c *ClauseMysql) Or(field, condition, value string) Clause {
	c.special("or", field, condition, value)
	return c
}

func (c *ClauseMysql) In(field, value string) Clause {
	c.condition(field, "in", value)
	return c
}

func (c *ClauseMysql) Between(field, value string) Clause {
	c.special("between", field, "and", value)
	return c
}

func (c *ClauseMysql) Like(field, value string) Clause {
	c.condition(field, "like", value)
	return c
}

func (c *ClauseMysql) special(specialCondition, field, condition, value string) Clause {
	c.whereArgs = append(c.whereArgs, specialCondition)
	c.condition(field, condition, value)
	return c
}

func (c *ClauseMysql) condition(field, condition, value string) Clause {
	c.whereCondition = append(c.whereCondition, condition)
	c.whereFields = append(c.whereFields, field)
	c.whereValues = append(c.whereValues, value)
	field = utils.EscapeQuote(field)
	value = utils.EscapeQuote(value)
	arg_ := utils.AddString([]string{"`", field, "`"})
	value_ := utils.AddString([]string{"'", value, "'"})
	c.whereArgs = append(c.whereArgs, utils.AddString([]string{arg_, " ", condition, " ", value_}))
	return c
}

func (c *ClauseMysql) Query() (strings.Builder, []any, any, int) {
	defer c.Clear()
	var (
		fields string
		values string
		model  any
	)
	fields, values = c.parseSql()
	switch c.method {
	case utils.InsertMethod:
		// insert into #{c.table}(#{c.fields}) values(#{c.fields});}
		c.sql.WriteString("insert into `%s`(%s) values(%s);")
		args := []any{c.table, fields, values}
		c.args = append(c.args, args...)
		//return fmt.Sprintf("insert into %s(%s) values(%s);", c.table, fields, values)
		model = c.insertModel
	case utils.DeleteMethod:
		// delete from #{c.table} where #{c.whereArgs}
		if len(c.whereArgs) > 0 {
			args := c.ParseWhere()
			c.sql.WriteString("delete from `%s` where %s;")
			c.args = []any{c.table}
			c.args = append(c.args, args)
			break
			//return fmt.Sprintf("delete from %s where %s;", c.table, args)
		}

		c.sql.WriteString("delete from `%s`;")
		c.args = append(c.args, c.table)
		//return fmt.Sprintf("delete from %s;", c.table)

	case utils.SelectMethod:
		//
		if len(c.whereArgs) > 0 {
			args := c.ParseWhere()
			c.sql.WriteString("select %s from `%s` where %s;")
			c.args = append(c.args, fields, c.table, args)
			//return fmt.Sprintf("select %s from %s where %s;", fields, c.table, args)
			break
		}

		c.sql.WriteString("select %s from `%s`;")
		args := []any{fields, c.table}
		c.args = append(c.args, args...)
		//return fmt.Sprintf("select %s from %s;", fields, c.table)

	case utils.UpdateMethod:
		//
		setArgs := c.parseUpdate()
		whereArgs := c.ParseWhere()
		if len(setArgs) > 0 {
			c.sql.WriteString("update `%s` set %s where %s;")
			c.args = []any{c.table, setArgs, whereArgs}
			//return fmt.Sprintf("update %s set %s where %s;", c.table, setArgs, args)
			break
		}
		c.sql.WriteString("update `%s` set %s;")
		c.args = []any{c.table, setArgs}
		//return fmt.Sprintf("update %s set %s;", c.table, setArgs)
	}
	//fmt.Println(c.sql.String(), c.args)

	return c.sql, c.args, model, c.method
}

func (c *ClauseMysql) parseSql() (string, string) {
	var (
		field  string
		value  string
		fields []string
		values []string
		//args   []any
	)
	for i, _ := range c.values {
		v := c.values[i]
		_v := utils.AddString([]string{"'", utils.EscapeQuote(v), "'"})
		values = append(values, _v)
	}
	for i, _ := range c.fields {
		f := c.fields[i]
		_f := utils.AddString([]string{"`", utils.EscapeQuote(f), "`"})
		fields = append(fields, _f)
	}
	field = strings.Join(fields, ", ")
	value = strings.Join(values, ", ")
	//temp := strings.Join(c.whereArgs, ", ")
	//args = []any{temp}

	return field, value
}

func (c *ClauseMysql) parseUpdate() string {
	if len(c.fields) != len(c.values) {
		panic("fields and values must equal")
	}
	var updateFields []string
	for i, _ := range c.fields {
		field := c.fields[i]
		value := c.values[i]
		field = utils.EscapeBackQuote(field)
		value = utils.EscapeSingleQuote(value)
		f := utils.AddString([]string{"`", field, "`"})
		v := utils.AddString([]string{"'", value, "'"})
		updateField := []string{f, "=", v}
		uf := strings.Join(updateField, " ")
		updateFields = append(updateFields, uf)
	}

	setArgs := strings.Join(updateFields, ", ")
	return setArgs
}

func (c *ClauseMysql) ParseWhere() string {
	args := strings.Join(c.whereArgs, " ")
	return args
}
