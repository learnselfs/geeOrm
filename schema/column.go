// Package schema @Author Bing
// @Date 2023/11/22 16:19:00
// @Desc
package schema

import "github.com/learnselfs/geeOrm/utils"

type Column struct {
	Name string
	Type utils.DataType
	Tag  string
}

func NewColumn(name string, t utils.DataType, tag string) *Column {
	return &Column{name, t, tag}
}
