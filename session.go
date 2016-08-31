package postgres

import (
	"github.com/jackc/pgx"
)

// Session 数据库会话
type Session struct {
	Tx *pgx.Tx
	DB *DB
}

// Commit 提交一个事务
func (s *Session) Commit() error {
	return s.Tx.Commit()
}

// Rollback 回滚一个事务
func (s *Session) Rollback() error {
	return s.Tx.Rollback()
}
