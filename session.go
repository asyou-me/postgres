package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Session struct {
	Tx *sql.Tx
}

func (d *DB) Begin() (*Session, error) {
	s, err := d.SQLDB.Begin()
	session := &Session{}
	if err != nil {
		return session, err
	}
	session.Tx = s
	return session, nil
}
