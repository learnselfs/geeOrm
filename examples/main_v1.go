// Package main @Author Bing
// @Date 2023/11/16 16:25:00
// @Desc
package examples

import (
	"fmt"
	"github.com/learnselfs/geeOrm/core"
	"github.com/learnselfs/geeOrm/utils"
)

type user struct {
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
	var u user
	fmt.Println(u)
	row, _ := session.DB().Exec("select * from user")
	fmt.Println(row)
}
