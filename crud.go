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

// One 获取单条查询结构
// param1:数据库表的名字
// param2:查询条件 sql写法 where xxx
// param3:查询后返回结果需要放入的对象
// param4:查询返回的结果需要包含的字段
func (d *DB) One(req string, out ReflectTableInterface, column ...string) (err error) {
	var relSlice *[]interface{}
	var reStr *string

	if len(column) == 0 {
		relSlice, reStr, _, err = out.AllReflect()
	} else {
		relSlice, reStr, err = out.Reflect(column)
	}

	if err != nil {
		return err
	}

	table := out.TableName()
	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + req
	err = d.Pool.QueryRow(query).Scan(*relSlice...)

	d.Info(query)

	return err
}

// All 获取单条查询结构
// param1:数据库表的名字
// param2:查询条件 sql写法 where xxx
// param3:查询后返回结果需要放入的对象
// param4:查询返回的结果需要包含的字段
// param5:排序相关的写法 order by xxx
// param6:查询数据的页数
// param7:查询数据的每页条数
func (d *DB) All(table string, req string, sort string,
	out interface{}, limit int16, offset int64,
	column ...string) (err error) {

	var relSlice *[]interface{}
	var reStr *string

	var newFunc NewFunc
	var ok bool

	if newFunc, ok = d.TableMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	}

	item := newFunc()
	if len(column) == 0 {
		relSlice, reStr, _, err = item.AllReflect()
	} else {
		relSlice, reStr, err = item.Reflect(column)
	}

	if err != nil {
		d.Warn(err.Error())
		return err
	}

	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + req + ` ` + sort + ` LIMIT $1 OFFSET $2`
	rows, err := d.Pool.Query(query, limit, offset)

	d.Info(query)
	if err != nil {
		rows.Close()
		d.Warn(err.Error())
		return err
	}

	var i = 0
	for rows.Next() {

		if i != 0 {
			item = newFunc()
			if len(column) == 0 {
				relSlice, reStr, _, err = item.AllReflect()
			} else {
				relSlice, reStr, err = item.Reflect(column)
			}

			if err != nil {
				d.Warn(err.Error())
				//d.loger.Warn("该处需要错误日志系统", err)
				continue
			}
		}
		i = i + 1

		err = rows.Scan(*relSlice...)
		if err != nil {
			d.Warn(err.Error())
			continue
		}

		item.AppendSelf(out)
	}

	rows.Close()
	return nil
}

// Insert 插入数据到数据库
func (d *DB) Insert(data ReflectTableInterface, column ...string) (string, error) {
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
// req:查询条件 sql写法 where xxx
// data:需要更新的数据的对象
// column:需要更新的字段
func (d *DB) Update(req string, data ReflectTableInterface, column []string) (err error) {
	if req == "" {
		return errors.New("更新条件不能为空")
	}
	var re *string
	var relSlice *[]interface{}

	relSlice, re, err = data.Reflect(column)

	if err != nil {
		return
	}

	table := data.TableName()
	_, err = d.Pool.Exec(`UPDATE "`+table+`" SET  `+*re+` `+req, *relSlice...)
	if err != nil {
		return
	}
	return
}

// Del 删除数据
// table:删除数据的表
// req:条件sql写法 where xxx
func (d *DB) Del(table string, req string) (err error) {
	_, err = d.Pool.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

// Count 获取数据的条数
// table:数据的表
// req:条件sql写法 where xxx
func (d *DB) Count(table string, req string) int64 {
	var re int64
	err := d.Pool.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + req).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}
