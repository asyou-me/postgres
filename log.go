package postgres

import (
	pulic_type "github.com/asyoume/lib/pulic_type"
)

type DBLog struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Msg   string `json:"type"`
	Level string `json:"level"`
	Time  int64  `json:"time"`
}

func (l *DBLog) GetLevel() string {
	return l.Level
}

func (l *DBLog) SetLevel(level string) {
	l.Level = level
}

func (l *DBLog) SetTime(t int64) {
	l.Time = t
}

func (this *DB) SetLog(log pulic_type.Loger) {
	this.loger = log
}

//传入debug日志
func (this *DB) Debug(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Debug(log)
}

//传入info日志
func (this *DB) Info(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Info(log)
}

//传入Print日志
func (this *DB) Print(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Print(log)
}

//传入Warn日志
func (this *DB) Warn(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Warn(log)
}

//传入Error日志
func (this *DB) Error(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Error(log)
}

//传入Fatal日志
func (this *DB) Fatal(str string) {
	log := &DBLog{
		Msg: str,
	}
	this.loger.Fatal(log)
}