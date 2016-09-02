package postgres

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

// Session 数据库会话
type Session struct {
	Tx     *pgx.Tx
	Engine *DB

	table string
	// 查询条件
	where string
	// 查询参数
	args []interface{}
	// 需要查询的字段 为空时列出所有数据
	column []string
	// 排序条件
	order []string
}

// Commit 提交一个事务
func (s *Session) Commit() error {
	return s.Tx.Commit()
}

// Rollback 回滚一个事务
func (s *Session) Rollback() error {
	return s.Tx.Rollback()
}

// Table 数据查询的表
func (s *Session) Table(table string) *Session {
	s.table = table
	s.where = ""
	s.args = []interface{}{}
	s.column = []string{}
	s.order = []string{}
	return s
}

// Where 数据查询条件
func (s *Session) Where(sql string, args ...interface{}) *Session {
	s.where = sql
	s.args = args
	return s
}

// OrderBy 数据排序条件
func (s *Session) OrderBy(sql ...string) *Session {
	s.order = sql
	return s
}

// Scan 查询一条数据
func (s *Session) Scan(out ReflectTable) error {
	var relSlice *[]interface{}
	var reStr *string
	var err error

	if len(s.column) == 0 {
		relSlice, reStr, _, err = out.AllReflect()
	} else {
		relSlice, reStr, err = out.Reflect(s.column)
	}

	if err != nil {
		return err
	}

	table := out.TableName()
	query := `SELECT ` + *reStr + ` FROM "` + table + `" ` + s.whereStr()
	err = s.Tx.QueryRow(query).Scan(*relSlice...)

	s.Engine.Info(query)

	return err
}

// Scans 查询多条数据
func (s *Session) Scans(out interface{}, args ...int64) error {
	var relSlice *[]interface{}
	var reStr *string

	var newFunc NewFunc
	var ok bool
	var err error
	var limit int64
	var offset int64
	var table = s.table
	var column = s.column
	var argsLen = len(s.args)

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
	s.args = append(s.args, offset, limit)

	if newFunc, ok = s.Engine.TableMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	}

	item := newFunc()
	if len(column) == 0 {
		relSlice, reStr, _, err = item.AllReflect()
	} else {
		relSlice, reStr, err = item.Reflect(column)
	}

	if err != nil {
		s.Engine.Warn(err.Error())
		return err
	}

	query := `SELECT ` + *reStr + ` FROM "` + table + `"` + s.whereStr() + "" + ` LIMIT $` + fmt.Sprint(argsLen+1) + ` OFFSET $` + fmt.Sprint(argsLen+2)
	s.Engine.Info(query)
	rows, err := s.Tx.Query(query, s.args...)

	if err != nil {
		rows.Close()
		s.Engine.Error(err.Error())
		return err
	}

	var i = 0
	for rows.Next() {
		if i != 0 {
			item = newFunc()
			if len(s.column) == 0 {
				relSlice, reStr, _, err = item.AllReflect()
			} else {
				relSlice, reStr, err = item.Reflect(s.column)
			}

			if err != nil {
				s.Engine.Warn(err.Error())
				continue
			}
		}
		i = i + 1

		err = rows.Scan(*relSlice...)
		if err != nil {
			s.Engine.Warn(err.Error())
			continue
		}
		item.AppendSelf(out)
	}
	rows.Close()
	return nil
}

// Set 修改部分数据 (兼容jsonb)
func (s *Session) Set(out []GSTYPE) (err error) {
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
	_, err = s.Tx.Exec(`UPDATE "`+s.table+`" SET `+sets+s.whereStr(), values...)
	return
}

// Get 获取部分数据 (兼容jsonb)
func (s *Session) Get(out []GSTYPE) (err error) {
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
		values[k] = &out[k].Value
	}
	return s.Tx.QueryRow(`SELECT ` + gets + ` FROM "` + s.table + `"` + s.whereStr()).Scan(values...)
}

func (s *Session) whereStr() string {
	if s.where != "" {
		return " WHERE " + s.where
	}
	return ""
}
