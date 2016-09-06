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
	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + q.whereStr()
	err = q.Engine.QueryRow(query).Scan(*relSlice...)

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
	var argsLen = len(q.args)

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
	q.args = append(q.args, offset, limit)

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

	query := `SELECT ` + *reStr + ` FROM "` + table + `"` + q.whereStr() + "" + ` LIMIT $` + fmt.Sprint(argsLen+1) + ` OFFSET $` + fmt.Sprint(argsLen+2)
	q.Engine.Info(query)
	rows, err := q.Engine.Query(query, q.args...)

	if err != nil {
		rows.Close()
		q.Engine.Error(err.Error())
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

// Set 设定jsonb数据
func (q *QueryBuilder) Set(out []GSTYPE) (err error) {
	var sets = ""
	var lenOut = len(out)
	var indexOut = lenOut - 1
	var values = make([]interface{}, lenOut)
	for k, v := range out {
		if k == indexOut {
			if v.Path == "" {
				sets = sets + v.Key + `=$` + fmt.Sprint(k+1)
			} else {
				sets = sets + v.Key + `=jsonb_set(` + v.Key + `,'{` + v.Path + `}',$` + fmt.Sprint(k+1) + `,true)`
			}
		} else {
			if v.Path == "" {
				sets = sets + v.Key + `=$` + fmt.Sprint(k+1) + `,`
			} else {
				sets = sets + v.Key + `=jsonb_set(` + v.Key + `,'{` + v.Path + `}',$` + fmt.Sprint(k+1) + `,true),`
			}
		}
		values[k] = v.Value
	}
	_, err = q.Engine.Exec(`UPDATE "`+q.table+`" SET `+sets+q.whereStr(), values...)
	return
}

// Get 获取jsonb数据
func (q *QueryBuilder) Get(out []GSTYPE) (err error) {
	var gets = ""
	var values = make([]interface{}, len(out))
	var indexOut = len(out) - 1
	for k, v := range out {
		if k == indexOut {
			if v.Path == "" {
				gets = gets + v.Key
			} else {
				gets = gets + v.Key + `#>>'{` + v.Path + `}'`
			}
		} else {
			if v.Path == "" {
				gets = gets + v.Key + `,`
			} else {
				gets = gets + v.Key + `#>>'{` + v.Path + `}',`
			}
		}
		values[k] = &(out[k])
	}
	err = q.Engine.QueryRow(`SELECT ` + gets + ` FROM "` + q.table + `"` + q.whereStr()).Scan(values...)
	if err != nil {
		return
	}
	return nil
}

func (q *QueryBuilder) whereStr() string {
	if q.where != "" {
		return " WHERE " + q.where
	}
	return ""
}
