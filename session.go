package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Session struct {
	Tx *sql.Tx
}

func Begin() (*Session, error) {
	s, err := SQLDB.Begin()
	session := &Session{}
	if err != nil {
		return session, err
	}
	session.Tx = s
	return session, nil
}
