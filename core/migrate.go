// Package core @Author Bing
// @Date 2023/12/20 11:03:00
// @Desc
package core

import (
	"errors"
	"fmt"
	"github.com/learnselfs/geeOrm/utils"
	"strings"
)

func (e *Engine) Migrate(model interface{}) error {
	Model := e.session.Model(model).Schema()
	raws, _ := e.session.DB().Query(fmt.Sprintf("select * from %s limit 1;", Model.Name))
	oldColumns, _ := raws.Columns()
	if e.session.Schema().Name != Model.Name {
		return errors.New("model struct name is error")
	}
	add, del := DifferentMaps(Model.ColumnMap, oldColumns)

	if len(add) > 0 {
		addColumns := GetNameType(add)
		addColumnString := utils.AddCommaString(addColumns)

		addRaw := fmt.Sprintf("alter table %s add column (%s);", Model.Name, addColumnString)
		_, err := e.session.DB().Exec(addRaw)
		if err != nil {
			return err
		}
	}

	if len(del) > 0 {
		delColumnString := utils.AddCommaString(del)

		delRaw := fmt.Sprintf("alter table %s del column (%s);", Model.Name, delColumnString)
		_, err := e.session.DB().Exec(delRaw)
		if err != nil {
			return err
		}
	}
	return nil

}

func GetNameType(columns []*Column) []string {
	var Columns []string
	for _, v := range columns {
		var s string
		l := []string{v.Name, string(v.Type)}
		s = utils.AddBlankString(l)
		Columns = append(Columns, s)
	}
	return Columns
}

func DifferentMaps(s1 map[string]*Column, s2 []string) ([]*Column, []string) {
	m := make(map[string]struct{})
	for _, v := range s2 {
		m[v] = struct{}{}
	}
	var add []*Column
	var del []string
	for k, v := range s1 {
		k = strings.ToLower(k)
		if _, ok := m[k]; !ok {
			add = append(add, v)
		} else {
			delete(s1, k)
			delete(m, k)
		}
	}

	for k, _ := range m {
		del = append(del, k)
	}
	return add, del
}
