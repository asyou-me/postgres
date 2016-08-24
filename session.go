package postgres

import

//"errors"
(
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type Session struct {
	Tx *pgx.Tx
	DB *DB
}

//提交一个事务
func (s *Session) Commit() error {
	return s.Tx.Commit()
}

//回滚一个事务
func (s *Session) Rollback() error {
	return s.Tx.Rollback()
}

/*

//查询单条数据并解析到结构体
//param1:数据库表的名字
//param2:查询条件 sql写法 where xxx
//param3:查询后返回结果需要放入的对象
//param4:查询返回的结果需要包含的字段
func (s *Session) One(table string, req string, out interface{}, column []string) (err error) {
	var re *[]interface{}
	var re_str *string

	if sqlf, ok := s.DB.SqlFuncMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		if !s.DB.SqlCheckMap[table](out) {
			return errors.New("表" + table + "与传入的结构体不兼容")
		}
		re, re_str, err = sqlf(out, column)
		if err != nil {
			return err
		}
	}

	err = s.Tx.QueryRow(`SELECT ` + *re_str + ` FROM "` + table + `" ` + req).Scan((*re)...)
	if err != nil {
		return err
	}
	return nil
}

//获取单条查询结构
//param1:数据库表的名字
//param2:查询条件 sql写法 where xxx
//param3:查询后返回结果需要放入的对象
//param4:查询返回的结果需要包含的字段
//param5:排序相关的写法 order by xxx
//param6:查询数据的页数
//param7:查询数据的每页条数
func (s *Session) All(table string, req string, out interface{}, column []string, sort string, p int, limit int16) (err error) {

	var re *[]interface{}
	var re_str *string
	var out_item interface{}

	if sqlf, ok := s.DB.SqlFuncMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		if !s.DB.SqlCheck2Map[table](out) {
			return errors.New("表" + table + "与传入的结构体不兼容")
		}
		out_item = s.DB.SqlNewMap[table]()
		re, re_str, err = sqlf(out_item, column)
		if err != nil {
			return err
		}
	}
	re_slice := *re

	rows, err := s.Tx.Query(`SELECT `+*re_str+` FROM "`+table+`" `+req+` `+sort+` LIMIT $1 OFFSET $2`, limit, int(limit)*(p-1))
	if err != nil {
		s.DB.loger.Warn("该处需要错误日志系统", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(re_slice...)
		if err != nil {
			s.DB.loger.Warn("该处需要错误日志系统", err)
			return err
		}
		s.DB.SqlAddMap[table](out, out_item)
	}

	rows.Close()
	return nil
}

//插入数据到数据库
func (s *Session) Insert(table string, data interface{}) (err error) {
	var re *string
	var re_str *string
	var rel_s *[]interface{}
	if sqlr, ok := s.DB.AllReflectMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, re_str, rel_s, err = sqlr(data)
		if err != nil {
			return
		}
	}
	rels := *rel_s
	_, err = s.Tx.Exec(`INSERT INTO "`+table+`" (`+*re+`) VALUES (`+*re_str+`)`, rels...)
	if err != nil {
		return
	}
	return
}

//更新数据到数据库
//param1:更新数据的表
//param2:查询条件 sql写法 where xxx
//param3:需要更新的数据的对象
//param4:需要更新的字段
func (s *Session) Update(table string, req string, data interface{}, column []string) (err error) {
	if req == "" {
		return errors.New("更新条件不能为空")
	}
	var re *string
	var rel_s *[]interface{}
	if sqlr, ok := s.DB.UpdateReflectMap[table]; !ok {
		return errors.New("表" + table + "未初始化")
	} else {
		re, rel_s, err = sqlr(data, column)
		if err != nil {
			return
		}
	}
	rels := *rel_s

	_, err = s.Tx.Exec(`UPDATE "`+table+`" SET  `+*re+` `+req, rels...)
	if err != nil {
		return
	}
	return
}

//删除数据
//param1:删除数据的表
//param2:条件sql写法 where xxx
func (s *Session) Del(table string, req string) (err error) {
	_, err = s.Tx.Exec(`DELETE FROM ` + table + ` ` + req)
	if err != nil {
		return
	}
	return
}

//获取数据的条数
//param1:数据的表
//param2:条件sql写法 where xxx
func (s *Session) Count(table string, req string) int64 {
	var re int64
	err := s.Tx.QueryRow(`SELECT COUNT(*) FROM ` + table + ` ` + req).Scan(&re)
	if err != nil {
		return 0
	}
	return re
}

*/
