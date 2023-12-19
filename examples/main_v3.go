// Package main @Author Bing
// @Date 2023/12/19 14:39:00
// @Desc
package main

import (
	"errors"
	"github.com/learnselfs/geeOrm/core"
)

var (
	engine  *core.Engine
	session *core.Session
	err     error
)

type user struct {
	Id       int
	Name     string
	Password string
}

func init() {
	engine = core.NewEngine("mysql", "db", "db", "192.168.101.138", "db", 30666)
	session, err = engine.GetSession()

}

func crud(s *core.Session) (any, error) {
	u1 := user{Name: "addd", Password: "addd"}
	s.Insert(u1)
	s.Save()

	//var u2 user
	//s.Find(&u2)
	return nil, errors.New("eeeeee")
}

func main() {
	var u user
	t := session.Model(u)

	t.Transaction(crud)
}
