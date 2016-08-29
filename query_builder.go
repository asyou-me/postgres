package postgres

import "errors"

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

func (q *QueryBuilder) Table(table string) {
	q.table = table
}

func (q *QueryBuilder) Where(sql string, args ...interface{}) {
	q.where = sql
	q.args = args
}

func (q *QueryBuilder) OrderBy(sql ...string) {
	q.order = sql
}

func (q *QueryBuilder) Scan(out ReflectTableInterface) error {
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

func (q *QueryBuilder) Set(out GSTYPE) error {
	return nil
}

func (q *QueryBuilder) Get(out GSTYPE) error {
	return nil
}
