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
	clause.Clause
	db      *sql.DB
	tx      *sql.Tx
	dialect dialect.Dialect
	schema  *Schema
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

type CommonDB interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func NewSession(db *sql.DB, dialect dialect.Dialect, clause clause.Clause) *Session {
	return &Session{db: db, dialect: dialect, Clause: clause}
}
