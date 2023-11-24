// Package core @Author Bing
// @Date 2023/11/16 14:27:00
// @Desc
package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/learnselfs/geeOrm/dialect"
	"github.com/learnselfs/geeOrm/utils"
)

type Engine struct {
	session  *Session
	dbType   string
	username string
	password string
	hostname string
	port     int
	database string
	dns      string
}

func (e *Engine) newDb() (*sql.DB, error) {
	db, err := sql.Open(e.dbType, e.dns)
	if err != nil {
		utils.ErrorLog.Printf("%s", err)
		return nil, err
	}
	utils.InfoLog.Printf("connect [%s]%s success.", e.dbType, e.dns)
	return db, nil
}

func (e *Engine) GetSession() (*Session, error) {
	return e.session, nil
}

func (e *Engine) Close() error {
	err := e.session.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewEngine(dbType, username, password, hostname, database string, port int, dialect dialect.Dialect) *Engine {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)
	e := &Engine{
		dbType:   dbType,
		username: username,
		password: password,
		hostname: hostname,
		port:     port,
		database: database,
		dns:      dns,
	}
	db, err := e.newDb()
	if err != nil {
		utils.ErrorLog.Printf("%s", err)
		return nil
	}
	e.session = NewSession(db, dialect)

	return e
}
