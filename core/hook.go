// Package core @Author Bing
// @Date 2023/12/15 13:55:00
// @Desc
package core

import (
	"github.com/learnselfs/geeOrm/utils"
	"reflect"
)

var ()

const (
	CreateBefore = "CreateBefore"
	CreateAfter  = "CreateAfter"

	ReadBefore = "ReadBefore"
	ReadAfter  = "ReadAfter"

	UpdateBefore = "UpdateBefore"
	UpdateAfter  = "UpdateAfter"

	DeleteBefore = "DeleteBefore"
	DeleteAfter  = "DeleteAfter"
)

type ICreateBefore interface {
	CreateBefore(*Session)
}

type ICreateAfter interface {
	CreateAfter(*Session)
}

type IReadBefore interface {
	ReadBefore(*Session)
}

type IReadAfter interface {
	ReadAfter(*Session)
}

type IUpdateBefore interface {
	UpdateBefore(*Session)
}

type IUpdateAfter interface {
	UpdateAfter(*Session)
}

type IDeleteBefore interface {
	DeleteBefore(*Session)
}

type IDeleteAfter interface {
	DeleteAfter(*Session)
}

func (s *Session) HookMethod(method string, model interface{}) {
	//m := reflect.Indirect(reflect.ValueOf(model))
	//i := m.Type().Elem()
	var i interface{}
	if reflect.TypeOf(model).Kind() == reflect.Ptr {
		i = model
	} else if reflect.TypeOf(model).Kind() == reflect.Struct {
		kind := reflect.Indirect(reflect.ValueOf(model)).Type()
		fields, values := utils.ParseAllStructFieldValueUnsafe(model)
		e := reflect.New(kind).Elem()
		for index := 0; index < e.NumField(); index++ {
			f := fields[index]
			v := values[index]
			v_ := reflect.ValueOf(v)
			e.FieldByName(f).Set(v_)
		}
		i = e.Addr().Interface()
	}
	switch method {
	case CreateBefore:
		I, ok := i.(ICreateBefore)
		if ok {
			I.CreateBefore(s)
		}
	case CreateAfter:
		if I, ok := i.(ICreateAfter); ok {
			I.CreateAfter(s)
		}
	case ReadBefore:
		if I, ok := i.(IReadBefore); ok {
			I.ReadBefore(s)
		}

	case ReadAfter:
		if I, ok := i.(IReadAfter); ok {
			I.ReadAfter(s)
		}

	case UpdateBefore:
		if I, ok := i.(IUpdateBefore); ok {
			I.UpdateBefore(s)
		}

	case UpdateAfter:
		if I, ok := i.(IUpdateAfter); ok {
			I.UpdateAfter(s)
		}

	case DeleteBefore:
		if I, ok := i.(IDeleteBefore); ok {
			I.DeleteBefore(s)
		}

	case DeleteAfter:
		if I, ok := i.(IDeleteAfter); ok {
			I.DeleteAfter(s)
		}

	default:
		panic("not method hook")
	}
}
