// Package core @Author Bing
// @Date 2023/11/23 11:21:00
// @Desc
package core

import (
	"database/sql"
	"fmt"
	"github.com/learnselfs/geeOrm/utils"
	"reflect"
	"strings"
)

func (s *Session) Model(m interface{}) *Session {
	if s.schema == nil || reflect.TypeOf(m) != reflect.TypeOf(s.schema.model) {
		s.schema = Parser(m, s.dialect)
		s.Clause.SetTable(s.schema.Name)
	}
	return s
}

func (s *Session) Row(f string, v ...any) *Session {
	s.sql.WriteString(f)
	s.sql.WriteString(" ")
	s.args = append(s.args, v...)
	return s
}

func (s *Session) Exec() (sql.Result, error) {

	row := fmt.Sprintf(s.sql.String(), s.args...)
	utils.DebugLog.Printf("%s", row)
	res, err := s.DB().Exec(row)

	return res, err
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	s.sql, s.args, s.Data, s.method = s.Clause.Query()
	s.HookMethod(ReadBefore, s.schema.model)
	raw := fmt.Sprintf(s.sql.String(), s.args...)
	row := s.DB().QueryRow(raw)
	if row.Err() != nil {
		utils.ErrorLog.Println(row.Err())
	}
	s.HookMethod(ReadAfter, s.schema.model)
	return row
}

func (s *Session) QueryRows() *sql.Rows {
	defer s.Clear()
	s.sql, s.args, s.Data, s.method = s.Clause.Query()
	raw := fmt.Sprintf(s.sql.String(), s.args...)
	utils.InfoLog.Println(raw)
	row, err := s.DB().Query(raw)
	if err != nil {
		utils.ErrorLog.Println(err)
	}
	return row
}

func (s *Session) Find(values interface{}) {
	v := reflect.Indirect(reflect.ValueOf(values))
	kind := v.Type().Elem()
	table := s.Model(reflect.New(kind).Elem().Interface())
	fieldNames, _ := utils.ParseAllStructFieldValueUnsafe(reflect.New(kind).Elem().Interface())
	s.Clause.Select(fieldNames)
	s.HookMethod(ReadBefore, s.schema.model)
	rows := s.QueryRows()
	defer rows.Close()
	for rows.Next() {
		value := reflect.New(kind).Elem()
		var fields []any
		//for _, field := range table.Schema().ColumnNames {
		for _, field := range table.Schema().ColumnNames {
			fields = append(fields, value.FieldByName(field).Addr().Interface())
		}
		rows.Scan(fields...)
		s.HookMethod(ReadAfter, value.Addr().Interface())
		v.Set(reflect.Append(v, value))
	}
}

func (s *Session) Save() {
	defer s.Clear()
	var (
		res sql.Result
		err error
	)
	s.sql, s.args, s.Data, s.method = s.Clause.Query()
	switch s.method {
	case utils.InsertMethod:
		s.HookMethod(CreateBefore, s.Data)
	case utils.DeleteMethod:
		s.HookMethod(DeleteBefore, s.schema.model)
	case utils.UpdateMethod:
		s.HookMethod(UpdateBefore, s.schema.model)
	}
	res, err = s.Exec()
	switch s.method {
	case utils.InsertMethod:
		s.HookMethod(CreateAfter, s.Data)
	case utils.DeleteMethod:
		s.HookMethod(DeleteAfter, s.schema.model)
	case utils.UpdateMethod:
		s.HookMethod(UpdateAfter, s.schema.model)
	}
	if err != nil {
		utils.ErrorLog.Println(err)
		return
	}
	utils.InfoLog.Printf("%#v", res)
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

// Drop table if exists tableName
func (s *Session) Drop() error {
	table := s.schema
	_, err := s.Row(fmt.Sprintf("drop table if exists %s;", table.Name)).Exec()
	return err
}
