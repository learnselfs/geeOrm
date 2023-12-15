// Package core @Author Bing
// @Date 2023/11/16 14:28:00
// @Desc
package core

import (
	"database/sql"
	"github.com/learnselfs/geeOrm/clause"
	"github.com/learnselfs/geeOrm/dialect"
	"github.com/learnselfs/geeOrm/utils"
	"strings"
)

type Session struct {
	clause  clause.Clause
	db      *sql.DB
	dialect dialect.Dialect
	schema  *Schema
	sql     strings.Builder
	args    []interface{}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.args = make([]interface{}, 0)
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Row(f string, v ...any) *Session {
	s.sql.WriteString(f)
	s.sql.WriteString(" ")
	s.args = append(s.args, v...)
	return s
}

func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear()
	utils.InfoLog.Println(s.sql.String(), s.args)
	res, err := s.db.Exec(s.sql.String(), s.args...)
	if err != nil {
		utils.ErrorLog.Println(err)
	}
	return res, err
}

func (s *Session) QueryRow(f string, v ...any) *sql.Row {
	defer s.Clear()
	return s.db.QueryRow(f, v...)
}

func (s *Session) QueryRows(f string, v ...any) (*sql.Rows, error) {
	defer s.Clear()
	row, err := s.db.Query(f, v...)
	if err != nil {
		utils.ErrorLog.Println(err)
	}
	return row, err
}

func NewSession(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}
