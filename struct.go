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
		if !SqlCheckMap[table](out) {
			return errors.New("表" + table + "与传入的结构体不兼容")
		}
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
	var out_item interface{}

	if sqlf, ok := SqlFuncMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		if !SqlCheck2Map[table](out) {
			return errors.New("表" + table + "与传入的结构体不兼容")
		}
		out_item = SqlNewMap[table]()
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

func Insert(table string, data interface{}) (err error) {
	var re *string
	var re_str *string
	var rel_s *[]interface{}
	if sqlr, ok := AllReflectMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, re_str, rel_s, err = sqlr(data)
		if err != nil {
			return
		}
	}
	rels := *rel_s
	_, err = SQLDB.Exec(`INSERT INTO `+table+` (`+*re+`) VALUES(`+*re_str+`)`, rels...)
	if err != nil {
		return
	}
	return
}

func Update(table string, req string, data interface{}, column []string) (err error) {
	if req == "" {
		return errors.New("更新条件不能为空")
	}
	var re *string
	var rel_s *[]interface{}
	if sqlr, ok := UpdateReflectMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, rel_s, err = sqlr(data, column)
		if err != nil {
			return
		}
	}
	rels := *rel_s

	_, err = SQLDB.Exec(`UPDATE `+table+` SET  `+*re+` `+req, rels...)
	if err != nil {
		return
	}
	return
}
