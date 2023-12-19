// Package core @Author Bing
// @Date 2023/12/19 13:38:00
// @Desc
package core

import (
	"database/sql"
	"github.com/learnselfs/geeOrm/utils"
)

type ftx func(s *Session) (any, error)

func (s *Session) Begin() *sql.Tx {
	tx, err := s.db.Begin()
	if err != nil {
		utils.ErrorLog.Println(err)
	}
	s.tx = tx
	return s.tx
}

func (s *Session) Commit() {
	err := s.tx.Commit()
	if err != nil {
		utils.ErrorLog.Println(err)
	}
}

func (s *Session) Rollback() {
	err := s.tx.Rollback()
	if err != nil {
		utils.ErrorLog.Println(err)
	}
}

func (s *Session) Transaction(f ftx) (res any, err error) {
	s.Begin()
	defer func() {
		if p := recover(); p != nil {
			s.Rollback()
			panic(p)
		} else if err != nil {
			s.Rollback()
			panic(err)
		} else {
			s.Commit()
		}
	}()
	return f(s)
}
