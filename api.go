package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type ReflectFunc func(s interface{}, column []string) (*[]interface{}, *string, error)
type NewFunc func() interface{}
type AddFunc func(all interface{}, s interface{})

var (
	SQLDB      *sql.DB
	loger      Loger
	SqlFuncMap map[string]ReflectFunc = map[string]ReflectFunc{}
	SqlNewMap  map[string]NewFunc     = map[string]NewFunc{}
	SqlAddMap  map[string]AddFunc     = map[string]AddFunc{}
	AllColumn                         = []string{}
)

func Init(sqlurl string, l Loger) error {
	//初始化数据库
	var err error
	SQLDB, err = sql.Open("postgres", sqlurl)
	loger = l
	return err
}

func Map_list() {
}

func Map_add() {
}
