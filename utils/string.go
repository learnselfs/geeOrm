// Package utils @Author Bing
// @Date 2023/12/6 12:19:00
// @Desc
package utils

import (
	"strings"
)

func AddString(s []string) string {
	//var str strings.Builder
	//for _, v := range s {
	//	str.WriteString(v)
	//}
	//return str.String()
	return AddSpaceString(s, "")
}
func AddCommaString(s []string) string {
	return AddSpaceString(s, ", ")
}
func AddBlankString(s []string) string {
	return AddSpaceString(s, " ")
}
func AddSpaceString(s []string, space string) string {
	var str strings.Builder
	for i := 0; i < len(s)-1; i++ {
		str.WriteString(s[i])
		str.WriteString(space)
	}
	str.WriteString(s[len(s)-1])
	return str.String()
}

func EscapeSingleQuote(s string) string {
	return strings.Join(strings.Split(s, "'"), `\'`)
	//return strings.ReplaceAll(s, "'", "")
}

func EscapeDoubleQuote(s string) string {
	return strings.Join(strings.Split(s, `"`), `\"`)
	//return strconv.Quote(s)
}

func EscapeBackQuote(s string) string {
	return strings.Join(strings.Split(s, "`"), "")
}

func EscapeQuote(s string) string {
	s1 := EscapeDoubleQuote(s)
	s2 := EscapeSingleQuote(s1)
	s3 := EscapeBackQuote(s2)
	return s3
}
