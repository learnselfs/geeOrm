// Package test @Author Bing
// @Date 2023/11/23 15:36:00
// @Desc
package test

import (
	"github.com/learnselfs/geeOrm/core"
	"github.com/learnselfs/geeOrm/dialect"
	"testing"
)

type T struct {
	//ID       string `orm:"AUTO_INCREMENT"`
	Username string
	Password string
}

func TestTable(t *testing.T) {
	t.Helper()
	d, _ := dialect.GetDialect("mysql")
	e := core.NewEngine("mysql", "db", "db", "192.168.101.138", "db", 30666, d)
	s, _ := e.GetSession()
	s.Model(T{Username: "user", Password: "password"})
	err := s.Create()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Drop()
	if err != nil {
		t.Fatal(err)
	}
}
