package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//代替反射的方法,通过字段名获取结构体的值，基于预编译的方法
type ReflectFunc func(s interface{}, column []string) (*[]interface{}, *string, error)

//基于预编译的方法获取基本元对象
type NewFunc func() interface{}

//数组添加元素
type AddFunc func(all interface{}, s interface{})

//数组添加元素
type CheckFunc func(s interface{}) bool

//插入结构反射
type StructReflect func(s interface{}) (*string, *string, *[]interface{}, error)

//插入结构反射
type UpdateReflect func(s interface{}, column []string) (*string, *[]interface{}, error)

var (
	AllColumn = []string{}
)

type DB struct {
	//数据操作连接
	SQLDB *sql.DB
	loger Loger
	//数据库结构反射方法map
	SqlFuncMap       map[string]ReflectFunc
	SqlNewMap        map[string]NewFunc
	SqlAddMap        map[string]AddFunc
	SqlCheckMap      map[string]CheckFunc
	SqlCheck2Map     map[string]CheckFunc
	AllReflectMap    map[string]StructReflect
	UpdateReflectMap map[string]UpdateReflect
}

func (d *DB) Open(sqlurl string, l Loger) error {
	//初始化数据库
	var err error
	d.SQLDB, err = sql.Open("postgres", sqlurl)
	//传入外部日志模块
	d.loger = l
	return err
}

//定义新建数据库操作对象的当法
func NewDB() *DB {
	db := DB{}
	db.SqlFuncMap = map[string]ReflectFunc{}
	db.SqlNewMap = map[string]NewFunc{}
	db.SqlAddMap = map[string]AddFunc{}
	db.SqlCheckMap = map[string]CheckFunc{}
	db.SqlCheck2Map = map[string]CheckFunc{}
	db.AllReflectMap = map[string]StructReflect{}
	db.UpdateReflectMap = map[string]UpdateReflect{}
	return &db
}
