package postgres

import ()

func (d *DB) Del(table string, req string) (err error) {
	_, err = d.SQLDB.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

func (d *DB) Count(table string, req string) int64 {
	var re int64
	err := d.SQLDB.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + req).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}
