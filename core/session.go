// Package core @Author Bing
// @Date 2023/11/16 14:28:00
// @Desc
package core

import (
	"database/sql"
	"strings"
)

type Session struct {
	db   *sql.DB
	sql  strings.Builder
	args []interface{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.args = make([]interface{}, 0)
}

func (s *Session) Row(f string, v ...any) {
	s.sql.WriteString(f)
	s.args = append(s.args, v...)
}
func newSession(db *sql.DB) *Session {
	return &Session{db: db}
}
