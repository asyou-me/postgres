package postgres

import "errors"

// Begin 开启一个事务
func (d *DB) Begin() (*Session, error) {
	s, err := d.Pool.Begin()
	session := &Session{}
	if err != nil {
		return session, err
	}
	session.Tx = s
	session.DB = d
	return session, nil
}

// Insert 插入数据到数据库
func (d *DB) Insert(data ReflectTable, column ...string) (string, error) {
	var re *string
	var reStr *string
	var relSlice *[]interface{}

	relSlice, re, reStr, err := data.AllReflect()
	if err != nil {
		return "", err
	}

	table := data.TableName()
	rel, err := d.Pool.Exec(`INSERT INTO "`+table+`" (`+*re+`) VALUES (`+*reStr+`)`, *relSlice...)
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
	var re *string
	var relSlice *[]interface{}

	relSlice, re, err = data.Reflect(column)

	if err != nil {
		return
	}

	table := data.TableName()
	_, err = d.Pool.Exec(`UPDATE "`+table+`" SET  `+*re+` `+req, *relSlice...)
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
	_, err = d.Pool.Exec(`DELETE FROM ` + table + ` ` + req)
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
	err := d.Pool.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + where).Scan(&re)
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
