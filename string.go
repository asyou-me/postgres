package postgres

import (
	"errors"
	"strings"
)

func OneString(table string, req string, re *string) (err error) {
	err = SQLDB.QueryRow(`SELECT d FROM ` + table + ` ` + req).Scan(re)
	if err != nil {
		return errors.New("查询出错，该错误已记录，请联系管理员")
	}
	return
}

func AllString(table string, req string, data *string, sort string, p int, limit int8) {
	rows, err := SQLDB.Query(`SELECT d FROM `+table+` `+req+` `+sort+` LIMIT $1 OFFSET $2`, limit, int(limit)*(p-1))
	if err != nil {
		loger.Warn("该处需要错误日志系统", err)
		return
	}
	var d string
	*data = "["
	for rows.Next() {
		err = rows.Scan(&d)
		*data = *data + d + `,`
	}
	rows.Close()
	*data = strings.TrimRight(*data, ",")
	*data = *data + "]"
}
