// Package core @Author Bing
// @Date 2023/11/22 16:14:00
// @Desc
package core

import (
	"github.com/learnselfs/geeOrm/dialect"
	"github.com/learnselfs/geeOrm/utils"
	"reflect"
)

type Schema struct {
	model       interface{}
	Name        string
	Columns     []*Column
	ColumnNames []string
	ColumnMap   map[string]*Column
}

func newSchema(model interface{}) *Schema {
	t := reflect.Indirect(reflect.ValueOf(model)).Type()
	return &Schema{model: model, Name: t.Name(), ColumnNames: make([]string, 0), ColumnMap: make(map[string]*Column)}
}

func (s *Schema) Column(name string) *Column {
	return s.ColumnMap[name]
}

func Parser(model interface{}, d dialect.Dialect) *Schema {
	S := newSchema(model)
	t := reflect.TypeOf(S.model)
	for i := 0; i < t.NumField(); i++ {
		p := t.Field(i)
		_type := d.TypeOf(reflect.Indirect(reflect.ValueOf(p.Type)))
		column := NewColumn(p.Name, _type, "")
		if tag := p.Tag.Get("orm"); len(tag) > 0 {
			column.Tag = tag
		}
		S.Columns = append(S.Columns, column)
		S.ColumnNames = append(S.ColumnNames, p.Name)
		S.ColumnMap[p.Name] = column
	}
	return S
}

type Column struct {
	Name string
	Type utils.DataType
	Tag  string
}

func NewColumn(name string, t utils.DataType, tag string) *Column {
	return &Column{name, t, tag}
}
