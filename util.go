package postgres

import ()

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
