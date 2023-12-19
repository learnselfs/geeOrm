// Package clause @Author Bing
// @Date 2023/11/24 15:10:00
// @Desc
package clause

import "strings"

type Clause interface {
	SetTable(table string)
	Query() (strings.Builder, []any, any, int)
	Insert(interface{}) Clause
	Delete() Clause
	Select(fields []string) Clause
	Update(interface{}) Clause
	Where(whereFields string, whereValues string) Clause
	Like(field, value string) Clause
	Clear()
}

var ClauseMap = make(map[string]Clause)

func RegisterClause(dbType string, clause Clause) error {
	ClauseMap[dbType] = clause
	return nil
}

func GetClause(dbType string) (Clause, error) {
	return ClauseMap[dbType], nil
}
