package postgres

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
