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
func (d *DB) Insert(data ReflectTable) (string, error) {
	var re *string
	var reStr *string
	var relSlice *[]interface{}

	relSlice, re, reStr, err := data.AllReflect()
	if err != nil {
		return "", err
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
func (d *DB) Update(req string, data ReflectTable, column []string) (err error) {
	if req == "" {
		return errors.New("更新条件不能为空")
	}
	req = "WHERE " + req
	fmt.Println("req:", req)
	var re string
	var relSlice *[]interface{}

	relSlice, _, err = data.Reflect(column)
	if err != nil {
		return
	}
	for index, v := range column {
		if index > 0 {
			re = re + `,"` + v + `"=$` + fmt.Sprint(index+1)
		} else {
			re = re + `"` + v + `"=$` + fmt.Sprint(index+1)
		}
	}

	table := data.TableName()
	_, err = d.Exec(`UPDATE "`+table+`" SET  `+re+` `+req, *relSlice...)
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
	sql := `DELETE FROM ` + table + ` ` + req
	_, err = d.Exec(sql)
	d.Info(sql)
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
	d.Info(`SELECT COUNT(*) FROM ` + table + ` ` + where)
	err := d.QueryRow(`SELECT COUNT(*) FROM "` + table + `" WHERE ` + where).Scan(&re)
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
