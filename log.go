package postgres

import (
	pulic_type "github.com/asyoume/lib.v1/pulic_type"
)

type dbLog struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Msg   string `json:"type"`
	Level string `json:"level"`
	Time  int64  `json:"time"`
}

func (l *dbLog) GetLevel() string {
	return l.Level
}

func (l *dbLog) SetLevel(level string) {
	l.Level = level
}

func (l *dbLog) SetTime(t int64) {
	l.Time = t
}

// SetLog 定义数据库记录对象
func (db *DB) SetLog(log pulic_type.Logger) {
	db.loger = log
}

// Debug 传入debug日志
func (db *DB) Debug(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Debug(log)
}

// Info 传入info日志
func (db *DB) Info(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Info(log)
}

// Print 传入Print日志
func (db *DB) Print(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Print(log)
}

// Warn 传入Warn日志
func (db *DB) Warn(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Warn(log)
}

// Error 传入Error日志
func (db *DB) Error(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Error(log)
}

// Fatal 传入Fatal日志
func (db *DB) Fatal(str string) {
	log := &dbLog{
		Msg: str,
	}
	db.loger.Fatal(log)
}
