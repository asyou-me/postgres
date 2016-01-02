package postgres

import (
	"encoding/json"
)

func Insert(table string, data interface{}) (err error) {
	jsonString, _ := json.Marshal(data)
	_, err = SQLDB.Exec(`INSERT INTO `+table+` (d) VALUES($1)`, string(jsonString))
	if err != nil {
		return
	}
	return
}

func Update(table string, req string, data interface{}) (err error) {
	var jsonString []byte
	jsonString, err = json.Marshal(data)
	_, err = SQLDB.Exec(`UPDATE `+table+` SET d =jsonb_merge(d,$1) `+req, string(jsonString))
	if err != nil {
		return
	}
	return
}

func Del(table string, req string) (err error) {
	_, err = SQLDB.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

func Count(table string, req string) int64 {
	var re int64
	err := SQLDB.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + req).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}
