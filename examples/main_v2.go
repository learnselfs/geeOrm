// Package main @Author Bing
// @Date 2023/11/16 16:25:00
// @Desc
package main

import (
	"fmt"
	"github.com/learnselfs/geeOrm/core"
	"github.com/learnselfs/geeOrm/utils"
)

type User struct {
	id   int
	name string
	pass string
}

func main() {
	engine := core.NewEngine("mysql", "db", "db", "192.168.101.138", "db", 30666)
	session, err := engine.GetSession()
	if err != nil {
		utils.ErrorLog.Println(err)
		return
	}
	var user User
	fmt.Println(user)
	//row, _ := session.DB().Exec("select * from user")
	u := session.Model(user)
	u.Insert()
	fmt.Println(row)
}
