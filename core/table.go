// Package core @Author Bing
// @Date 2023/11/23 11:21:00
// @Desc
package core

import (
	"fmt"
	"reflect"
	"strings"
)

func (s *Session) Model(m interface{}) *Session {
	if s.schema == nil || reflect.TypeOf(m) != reflect.TypeOf(s.schema.model) {
		s.schema = Parser(m, s.dialect)
	}
	return s
}

func (s *Session) Schema() *Schema {
	if s.schema == nil {
		panic("schema is nil")
	}
	return s.schema
}

func (s *Session) Create() error {
	table := s.Schema()
	args := make([]string, 0)
	for _, column := range table.Columns {
		arg := []string{column.Name, string(column.Type), column.Tag}
		columnSql := strings.Join(arg, " ")
		args = append(args, columnSql)
	}
	columns := strings.Join(args, ",")
	_, err := s.Row(fmt.Sprintf("create table if not  exists %s (%s);", table.Name, columns)).Exec()
	return err
}

// drop table if exists tableName
func (s *Session) Drop() error {
	table := s.schema
	_, err := s.Row(fmt.Sprintf("drop table if exists %s;", table.Name)).Exec()
	return err
}
