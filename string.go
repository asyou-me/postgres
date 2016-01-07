package postgres

import (
	"errors"
	"strings"
)

func (d *DB) OneString(table string, req string, re *string) (err error) {
	err = d.SQLDB.QueryRow(`SELECT d FROM ` + table + ` ` + req).Scan(re)
	if err != nil {
		return errors.New("查询出错，该错误已记录，请联系管理员")
	}
	return
}

func (d *DB) AllString(table string, req string, data *string, sort string, p int, limit int8) {
	rows, err := d.SQLDB.Query(`SELECT d FROM `+table+` `+req+` `+sort+` LIMIT $1 OFFSET $2`, limit, int(limit)*(p-1))
	if err != nil {
		d.loger.Warn("该处需要错误日志系统", err)
		return
	}
	var ds string
	*data = "["
	for rows.Next() {
		err = rows.Scan(&ds)
		*data = *data + ds + `,`
	}
	rows.Close()
	*data = strings.TrimRight(*data, ",")
	*data = *data + "]"
}
