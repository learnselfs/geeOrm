// Package dialect @Author Bing
// @Date 2023/11/22 14:13:00
// @Desc
package dialect

import (
	"github.com/learnselfs/geeOrm/utils"
	"reflect"
)

var _ Dialect = (*mysql)(nil)

func init() {
	err := RegisterDialect("mysql", &mysql{})
	if err != nil {
		return
	}
}

const (
	Bool   utils.DataType = "bool"
	Int                   = "integer"
	Uint                  = "uint"
	Float                 = "float"
	Double                = "double"
	String                = "varchar(255)"
	Time                  = "time"
	Bytes                 = "bytes"
)

type mysql struct{}

func (m *mysql) TypeOf(v reflect.Value) utils.DataType {
	switch v.Kind() {
	case reflect.Bool:
		return Bool

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Uint
	case reflect.Float32:
		return Float
	case reflect.Float64:
		return Double

	case reflect.String:
		return String
	default:
		return String
	}
}

func (m *mysql) ExistsTable(name string) (string, []interface{}) {
	args := []interface{}{name}
	return "create table if exists ?", args
}
