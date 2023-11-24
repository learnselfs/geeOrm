// Package test @Author Bing
// @Date 2023/11/24 20:53:00
// @Desc
package test

import (
	"github.com/learnselfs/geeOrm/clause"
	"testing"
)

func TestClause(t *testing.T) {
	clause := &clause.Clause{}
	field := []string{"name", "pass"}
	insert_1 := clause.Insert("User", field, []string{"guest", "12345"}).Query()
	insert_2 := "insert into User(name,pass) values(guest,12345);"
	if insert_1 != insert_2 {
		t.Fatal(insert_1)
	}

	delete_1 := clause.Delete("User").Where("name = pass").Query()
	delete_2 := "delete from User where name = pass;"
	if delete_1 != delete_2 {
		t.Fatal(delete_2)
	}
	fields := []string{"name"}
	values := []string{"admin"}
	update_1 := clause.Update("User", fields, values).Where("name = guest").Query()
	update_2 := "update User set name = admin where name = guest;"
	if update_1 != update_2 {
		t.Fatal(update_1)
	}

	select_1 := clause.Select("User", field).Where("name = guest").Query()
	select_2 := "select name,pass from User where name = guest;"
	if select_1 != select_2 {
		t.Fatal(select_1)
	}
}
