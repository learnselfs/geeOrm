// Package main @Author Bing
// @Date 2023/11/16 16:25:00
// @Desc
package main

import (
	"fmt"
	"github.com/learnselfs/geeOrm/core"
	"github.com/learnselfs/geeOrm/utils"
)

type user struct {
	Id       int
	Name     string
	Password string
}

var u *core.Session

func main() {
	engine := core.NewEngine("mysql", "db", "db", "192.168.101.138", "db", 30666)
	session, err := engine.GetSession()
	if err != nil {
		utils.ErrorLog.Println(err)
		return
	}
	var uTable user
	//row, _ := session.DB().Exec("select * from user")
	u = session.Model(uTable)
	//
	//insert()

	//find()
	//selectOne()
	//selectAll()

	//update()

	deleteOne()
	//deleteAll()

	//fmt.Println(strconv.Quote(`""`))
	//fmt.Println(strconv.Quote("``"))
	//fmt.Println(strconv.Quote("''"))

	//fmt.Println(utils.EscapeSingleQuote(" afi'   "))
	//fmt.Println(utils.EscapeDoubleQuote(` afi \' ""'`))
	//fmt.Println(utils.EscapeBackQuote(" afi ` "))
	//fmt.Println(utils.EscapeQuote(` afi "''"?<>/!@#$ `))

}
func insert() {
	name := `u1'--#"`
	pass := "p1"
	u1 := user{Name: name, Password: pass}

	u.Clause.Insert(u1)
	u.Save()

}

func selectOne() {
	var uTables []user
	// todo: select
	//u.Clause.Like("name", "admin1")
	//row := u.QueryRow()
	//u.Clause.Like("id", "102")
	u.Clause.Where("id`=101;-- +#`", "102")
	//_ = row.Scan(&uTable.Id, &uTable.Name, &uTable.Password)
	u.Find(&uTables)
	utils.InfoLog.Printf("%#v", uTables)
}

func selectAll() {
	// todo: select many
	u.Clause.Select([]string{"id", "password"})
	rows := u.QueryRows()
	defer rows.Close()
	for rows.Next() {
		var us user
		rows.Scan(&us.Id, &us.Password)
		fmt.Printf("%#v\n", us)
	}
}

func deleteAll() {

	//// todo: delete all
	u.Clause.Delete()
	u.Save()
}

func deleteOne() {
	// todo: select
	u.Clause.Delete().Where("id", "103")
	u.Save()
}

func update() {
	// todo: update
	var us = user{Name: "name", Password: "password' where id=104;'-- "}
	u.Clause.Update(us).Where("id", "103")
	u.Save()
}

func find() {
	var us []user
	u.Find(&us)
	fmt.Printf("%#v\n", us)
}
