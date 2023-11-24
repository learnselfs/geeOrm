// Package dialect @Author Bing
// @Date 2023/11/22 10:22:00
// @Desc
package dialect

import (
	"github.com/learnselfs/geeOrm/utils"
	"reflect"
)

type Dialect interface {
	TypeOf(reflect.Value) utils.DataType
	ExistsTable(string) (string, []interface{})
}

var dialectMap = make(map[string]Dialect)

func RegisterDialect(db string, dialect Dialect) error {
	dialectMap[db] = dialect
	return nil
}

func GetDialect(db string) (Dialect, error) {
	dialect := dialectMap[db]
	return dialect, nil
}
