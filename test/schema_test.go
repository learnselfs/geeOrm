// Package test @Author Bing
// @Date 2023/11/23 10:14:00
// @Desc
package test

import (
	"github.com/learnselfs/geeOrm/core"
	"github.com/learnselfs/geeOrm/dialect"
	"testing"
)

type User struct {
	name string `orm:"name"`
	info string `orm:"info"`
}

func TestSchema(t *testing.T) {
	u := User{"Guest", "this is info ."}
	db, err := dialect.GetDialect("mysql")
	if err != nil {
		return
	}
	s := core.Parser(u, db)
	if s.Name != "User" || len(s.Columns) != 2 {
		t.Fatalf("failed to parse struct ")
	}
	if s.Column("name").Tag != "name" {
		t.Fatalf("failed to parse tag ")
	}
}
