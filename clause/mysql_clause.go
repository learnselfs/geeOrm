// Package clause @Author Bing
// @Date 2023/11/24 15:10:00
// @Desc
package clause

import (
	"strings"
)

type Clause struct {
	method         int
	table          string
	fields         []string
	values         []string
	limit          string
	whereArgs      []string
	whereCondition []string
	orderBy        string
	sql            strings.Builder
	args           []string
}

func (c Clause) Clear() {
	c.whereArgs = make([]string, 0)
	c.sql.Reset()
	c.args = make([]string, 0)
}

const (
	insert_method int = iota
	delete_method
	select_method
	update_method
)

func (c *Clause) Insert(table string, fields []string, values []string) *Clause {
	c.method = insert_method
	c.table = table
	c.fields = fields
	c.values = values
	return c
}

func (c *Clause) Delete(table string) *Clause {
	c.method = delete_method
	c.table = table
	return c
}

func (c *Clause) Select(table string, fields []string) *Clause {
	c.method = select_method
	c.table = table
	c.fields = fields
	return c
}

func (c *Clause) Update(table string, fields []string, values []string) *Clause {
	c.method = update_method
	c.table = table
	c.fields = fields
	c.values = values
	return c
}

func (c *Clause) Where(args string) *Clause {
	c.whereArgs = []string{args}
	return c
}

func (c *Clause) And(args string) *Clause {
	c.condition("and", args)
	return c
}

func (c *Clause) Not(args string) *Clause {
	c.condition("not", args)
	return c
}

func (c *Clause) Or(args string) *Clause {
	c.condition("or", args)
	return c
}

func (c *Clause) In(field, value string) *Clause {
	c.special(field, "in", value)
	return c
}

func (c *Clause) Between(field, value string) *Clause {
	c.whereArgs = append(c.whereArgs, "between")
	c.special(field, "and", value)
	return c
}

func (c *Clause) Like(field, value string) *Clause {
	c.special(field, "like", value)
	return c
}

func (c *Clause) special(field, condition, value string) *Clause {
	c.whereArgs = append(c.whereArgs, field)
	c.whereArgs = append(c.whereArgs, condition)
	c.whereArgs = append(c.whereArgs, value)
	return c
}

func (c *Clause) condition(condition, args string) *Clause {
	c.whereArgs = append(c.whereArgs, condition)
	c.whereArgs = append(c.whereArgs, args)
	return c
}

func (c *Clause) Query() (strings.Builder, []string) {
	var (
		fields string
		values string
		args   string
	)
	fields, values, args = c.parseSql()
	defer c.Clear()
	switch c.method {
	case insert_method:
		// insert into #{c.table}(#{c.fields}) values(#{c.fields});}
		c.sql.WriteString("insert into %s(%s) values(%s);")
		args := []string{c.table, fields, values}
		c.args = append(c.args, args...)
		//return fmt.Sprintf("insert into %s(%s) values(%s);", c.table, fields, values)
	case delete_method:
		// delete from #{c.table} where #{c.whereArgs}
		if len(c.whereArgs) > 0 {

			c.sql.WriteString("delete from %s where %s;")
			args := []string{c.table, args}
			c.args = append(c.args, args...)
			//return fmt.Sprintf("delete from %s where %s;", c.table, args)
		}
		c.sql.WriteString("delete from %s;")
		c.args = append(c.args, c.table)
		//return fmt.Sprintf("delete from %s;", c.table)

	case select_method:
		//
		if len(c.whereArgs) > 0 {
			c.sql.WriteString("select %s from %s where %s;")
			args := []string{fields, c.table, args}
			c.args = append(c.args, args...)
			//return fmt.Sprintf("select %s from %s where %s;", fields, c.table, args)
		}
		c.sql.WriteString("select %s from %s;")
		args := []string{fields, c.table}
		c.args = append(c.args, args...)
		//return fmt.Sprintf("select %s from %s;", fields, c.table)

	case update_method:
		//
		setArgs := c.parseUpdate()
		if len(setArgs) > 0 {
			c.sql.WriteString("update %s set %s where %s;")
			args := []string{c.table, setArgs, args}
			c.args = append(c.args, args...)
			//return fmt.Sprintf("update %s set %s where %s;", c.table, setArgs, args)
		}
		c.sql.WriteString("update %s set %s;")
		args := []string{c.table, setArgs}
		c.args = append(c.args, args...)
		//return fmt.Sprintf("update %s set %s;", c.table, setArgs)
	}
	return c.sql, c.args
}

func (c *Clause) parseSql() (string, string, string) {
	var (
		fields string
		values string
		args   string
	)
	fields = strings.Join(c.fields, ",")
	values = strings.Join(c.values, ",")
	args = strings.Join(c.whereArgs, ",")
	return fields, values, args
}

func (c *Clause) parseUpdate() string {
	if len(c.fields) != len(c.values) {
		panic("fields and values must equal")
	}
	var updateFields []string
	for i, _ := range c.fields {
		var updateField []string
		updateField = append(updateField, c.fields[i])
		updateField = append(updateField, "=")
		updateField = append(updateField, c.values[i])
		uf := strings.Join(updateField, " ")
		updateFields = append(updateFields, uf)
	}
	setArgs := strings.Join(updateFields, " ")
	return setArgs
}
