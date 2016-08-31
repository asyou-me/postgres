package postgres

import (
	"errors"
	"fmt"
)

// QueryBuilder 数据查询构造器
type QueryBuilder struct {
	Engine *DB
	table  string
	// 查询条件
	where  string
	args   []interface{}
	column []string
	// 排序条件
	order []string
}

// Table 数据查询的表
func (q *QueryBuilder) Table(table string) *QueryBuilder {
	q.table = table
	return q
}

// Where 数据查询条件
func (q *QueryBuilder) Where(sql string, args ...interface{}) *QueryBuilder {
	q.where = sql
	q.args = args
	return q
}

// OrderBy 数据排序条件
func (q *QueryBuilder) OrderBy(sql ...string) *QueryBuilder {
	q.order = sql
	return q
}

// Scan 查询一条数据
func (q *QueryBuilder) Scan(out ReflectTable) error {
	var relSlice *[]interface{}
	var reStr *string
	var err error

	if len(q.column) == 0 {
		relSlice, reStr, _, err = out.AllReflect()
	} else {
		relSlice, reStr, err = out.Reflect(q.column)
	}

	if err != nil {
		return err
	}

	table := out.TableName()
	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + q.where
	err = q.Engine.Pool.QueryRow(query).Scan(*relSlice...)

	q.Engine.Info(query)

	return err
}

// Scans 查询多条数据
func (q *QueryBuilder) Scans(out interface{}, args ...int64) error {
	var relSlice *[]interface{}
	var reStr *string

	var newFunc NewFunc
	var ok bool
	var err error
	var limit int64
	var offset int64
	var table = q.table
	var column = q.column

	argLen := len(args)
	if argLen == 2 {
		limit = args[0]
		offset = args[1]
	} else if argLen == 1 {
		limit = args[0]
		offset = 0
	} else {
		limit = 10
		offset = 0
	}

	if newFunc, ok = q.Engine.TableMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	}

	item := newFunc()
	if len(column) == 0 {
		relSlice, reStr, _, err = item.AllReflect()
	} else {
		relSlice, reStr, err = item.Reflect(column)
	}

	if err != nil {
		q.Engine.Warn(err.Error())
		return err
	}

	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + q.where + ` ` + "" + ` LIMIT $1 OFFSET $2`
	rows, err := q.Engine.Pool.Query(query, args, limit, offset)

	q.Engine.Info(query)
	if err != nil {
		rows.Close()
		q.Engine.Warn(err.Error())
		return err
	}

	var i = 0
	for rows.Next() {
		if i != 0 {
			item = newFunc()
			if len(q.column) == 0 {
				relSlice, reStr, _, err = item.AllReflect()
			} else {
				relSlice, reStr, err = item.Reflect(q.column)
			}

			if err != nil {
				q.Engine.Warn(err.Error())
				//q.Engine.loger.Warn("该处需要错误日志系统", err)
				continue
			}
		}
		i = i + 1

		err = rows.Scan(*relSlice...)
		if err != nil {
			q.Engine.Warn(err.Error())
			continue
		}
		item.AppendSelf(out)
	}
	rows.Close()
	return nil
}

// Set 设定数据
func (q *QueryBuilder) Set(out []GSTYPE) (err error) {
	var sets = ""
	var values = make([]interface{}, len(out))
	for k, v := range out {
		sets = sets + `jsonb_ set(` + v.Key + `,'{` + v.Path + `}','$` + fmt.Sprint(k+1) + `'::jsonb,true)`
		values[k] = v
	}
	_, err = q.Engine.Pool.Exec(`UPDATE "`+q.table+`" SET `+sets, values...)
	return
}

// Get 获取数据
func (q *QueryBuilder) Get(out []GSTYPE) (err error) {
	var sets = ""
	var values = make([]interface{}, len(out))
	for k, v := range out {
		sets = sets + `jsonb_ set(` + v.Key + `,'{` + v.Path + `}','$` + fmt.Sprint(k+1) + `'::jsonb,true)`
		values[k] = &v
	}
	err = q.Engine.Pool.QueryRow(`SELECT ` + sets + ` FROM "` + q.table + `" `).Scan(values)
	return
}
