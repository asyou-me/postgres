package postgres

import (
	"errors"
)

func One(table string, req string, out interface{}, column []string) (err error) {
	var re *[]interface{}
	var re_str *string
	if sqlf, ok := SqlFuncMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, re_str, err = sqlf(out, column)
		if err != nil {
			return err
		}
	}
	err = SQLDB.QueryRow(`SELECT ` + *re_str + ` FROM ` + table + ` ` + req).Scan((*re)...)
	if err != nil {
		return err
	}
	return nil
}

func All(table string, req string, out interface{}, column []string, sort string, p int, limit int16) (err error) {

	var re *[]interface{}
	var re_str *string

	out_item := SqlNewMap[table]()

	if sqlf, ok := SqlFuncMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, re_str, err = sqlf(out_item, column)
		if err != nil {
			return err
		}
	}

	re_slice := *re

	rows, err := SQLDB.Query(`SELECT `+*re_str+` FROM `+table+` `+req+` `+sort+` LIMIT $1 OFFSET $2`, limit, int(limit)*(p-1))
	if err != nil {
		loger.Warn("该处需要错误日志系统", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(re_slice...)
		if err != nil {
			loger.Warn("该处需要错误日志系统", err)
			return err
		}
		SqlAddMap[table](out, out_item)
	}

	rows.Close()
	return nil
}
