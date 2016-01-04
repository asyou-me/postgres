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
	//数据操作连接
	SQLDB *sql.DB
	//数据日志对象
	loger            Loger
	SqlFuncMap       map[string]ReflectFunc   = map[string]ReflectFunc{}
	SqlNewMap        map[string]NewFunc       = map[string]NewFunc{}
	SqlAddMap        map[string]AddFunc       = map[string]AddFunc{}
	SqlCheckMap      map[string]CheckFunc     = map[string]CheckFunc{}
	SqlCheck2Map     map[string]CheckFunc     = map[string]CheckFunc{}
	AllReflectMap    map[string]StructReflect = map[string]StructReflect{}
	UpdateReflectMap map[string]UpdateReflect = map[string]UpdateReflect{}

	AllColumn = []string{}
)

func Init(sqlurl string, l Loger) error {
	//初始化数据库
	var err error
	SQLDB, err = sql.Open("postgres", sqlurl)
	//传入外部日志模块
	loger = l
	return err
}

func Map_list() {
}

func Map_add() {
}
