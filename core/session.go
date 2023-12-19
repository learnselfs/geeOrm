// Package core @Author Bing
// @Date 2023/11/16 14:28:00
// @Desc
package core

import (
	"database/sql"
	"github.com/learnselfs/geeOrm/clause"
	"github.com/learnselfs/geeOrm/dialect"
	"strings"
)

type Session struct {
	db      *sql.DB
	dialect dialect.Dialect
	schema  *Schema
	Clause  clause.Clause
	method  int
	sql     strings.Builder
	args    []interface{}
	// model
	Data interface{}
}

func (s *Session) Clear() {
	var i int
	var d any
	s.Data = d
	s.method = i
	s.sql.Reset()
	s.args = make([]interface{}, 0)
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func NewSession(db *sql.DB, dialect dialect.Dialect, clause clause.Clause) *Session {
	return &Session{db: db, dialect: dialect, Clause: clause}
}
