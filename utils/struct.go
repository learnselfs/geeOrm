// Package utils @Author Bing
// @Date 2023/12/12 15:35:00
// @Desc
package utils

import "reflect"

func ParseStructFieldValue(s interface{}) ([]string, []string) {
	var fields []string
	var values []string
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			_v := v.Field(i)
			_t := t.Field(i)
			if reflect.DeepEqual(_v.Interface(), reflect.Zero(_t.Type).Interface()) {
				continue
			}
			_t_ := []string{"`", _t.Name, "`"}
			t_ := AddString(_t_)
			fields = append(fields, t_)
			_v_ := []string{"'", _v.String(), "'"}
			v_ := AddString(_v_)
			values = append(values, v_)
		}
	}
	return fields, values
}

func ParseStructFieldValueUnsafe(s interface{}) ([]string, []string) {
	var fields []string
	var values []string
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			_v := v.Field(i)
			_t := t.Field(i)
			if reflect.DeepEqual(_v.Interface(), reflect.Zero(_t.Type).Interface()) {
				continue
			}
			fields = append(fields, _t.Name)
			values = append(values, _v.String())
		}
	}
	return fields, values
}
