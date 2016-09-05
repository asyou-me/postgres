package postgres

import (
	"errors"
	"fmt"
)

// Begin 开启一个事务
func (d *DB) Begin() (*Session, error) {
	s, err := d.ConnPool.Begin()
	session := &Session{}
	if err != nil {
		return session, err
	}
	session.Tx = s
	session.Engine = d
	return session, nil
}

// Insert 插入数据到数据库
func (d *DB) Insert(data ReflectTable, columns ...GSTYPE) (string, error) {
	var re *string
	var reStr *string
	var relSlice *[]interface{}

	relSlice, re, reStr, err := data.AllReflect()
	if err != nil {
		return "", err
	}

	// jsonb 数据
	var indexOut = len(columns) - 1
	var lenSlice = len(*relSlice)
	for k, v := range columns {
		if k == indexOut {
			*reStr = *reStr + v.Key + `$` + fmt.Sprint(k+lenSlice)
		} else {
			*reStr = *reStr + v.Key + `$` + fmt.Sprint(k+lenSlice) + `,`
		}
		*relSlice = append(*relSlice, v.Value)
	}

	table := data.TableName()
	rel, err := d.Exec(`INSERT INTO "`+table+`" (`+*re+`) VALUES (`+*reStr+`)`, *relSlice...)
	relStr := string(rel)

	return relStr, err
}

// Update 更新数据到数据库
//
// req:查询条件 sql写法 where xxx
//
// data:需要更新的数据的对象
//
// column:需要更新的字段
func (d *DB) Update(req string, data ReflectTable, column []string, columns ...GSTYPE) (err error) {
	if req == "" {
		return errors.New("更新条件不能为空")
	}
	req = "WHERE " + req
	var re *string
	var relSlice *[]interface{}

	relSlice, re, err = data.Reflect(column)
	if err != nil {
		return
	}

	var indexOut = len(columns) - 1
	for k, v := range columns {
		if k == indexOut {
			if v.Path == "" {
				*re = *re + v.Key + `=$` + fmt.Sprint(k+1)
			} else {
				*re = *re + v.Key + `=jsonb_set(` + v.Key + `,'{` + v.Path + `}',$` + fmt.Sprint(k+1) + `,true)`
			}
		} else {
			if v.Path == "" {
				*re = *re + v.Key + `=$` + fmt.Sprint(k+1) + `,`
			} else {
				*re = *re + v.Key + `=jsonb_set(` + v.Key + `,'{` + v.Path + `}',$` + fmt.Sprint(k+1) + `,true),`
			}
		}
		*relSlice = append(*relSlice, v.Value)
	}

	table := data.TableName()
	_, err = d.Exec(`UPDATE "`+table+`" SET  `+*re+` `+req, *relSlice...)
	return
}

// Del 删除数据
//
// table:删除数据的表
//
// req:条件sql写法 where xxx
func (d *DB) Del(table string, req string) (err error) {
	if req != "" {
		req = "WHERE " + req
	}
	_, err = d.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

// Count 获取数据的条数
//
// table:数据的表
//
// req:条件sql写法 where xxx
func (d *DB) Count(table string, where string) int64 {
	var re int64
	err := d.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + where).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}

// Table 建立一个 针对于 table 表的数据库查询对象
//
// table:数据库表名
func (d *DB) Table(table string) *QueryBuilder {
	builder := &QueryBuilder{
		Engine: d,
	}
	builder.Table(table)
	return builder
}
