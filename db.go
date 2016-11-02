/*
Package postgres 数据库处理对象
*/
package postgres

import (
	"time"

	"github.com/jackc/pgx"

	pulic_type "github.com/asyou-me/lib.v1/pulic_type"
)

// NewFunc 创建一个新的数据库对象的方法
type NewFunc func() ReflectTable

var (
	// AllColumn 数据库表全部字段
	AllColumn = []string{}
)

// DB 数据库处理对象
type DB struct {
	//数据操作连接池
	*pgx.ConnPool
	loger    pulic_type.Logger
	TableMap map[string]NewFunc
}

// Open 创建新的数据库对象
func (d *DB) Open(conf *pulic_type.MicroSerType, loger pulic_type.Logger) error {
	//初始化数据库
	var err error

	// 映射数据库连接参数
	var connConfig = pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     conf.Addr,
			User:     conf.Id,
			Password: conf.Secret,
			Database: conf.Attr["Database"].(string),
		},
		MaxConnections: int(conf.Attr["MaxConnections"].(int)),
		AcquireTimeout: time.Second * time.Duration(conf.Attr["AcquireTimeout"].(int)),
	}

	d.ConnPool, err = pgx.NewConnPool(connConfig)
	//传入外部日志模块
	d.loger = loger
	return err
}

// NewDB 定义新建数据库操作对象的当法
func NewDB(conf *pulic_type.MicroSerType, loger pulic_type.Logger) (*DB, error) {
	db := DB{}
	db.TableMap = map[string]NewFunc{}
	err := db.Open(conf, loger)
	return &db, err
}

// SetLog 定义数据库记录对象
func (d *DB) SetLog(log pulic_type.Logger) {
	d.loger = log
}

// Debug 传入debug日志
func (d *DB) Debug(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Debug(log)
}

// Info 传入info日志
func (d *DB) Info(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Info(log)
}

// Print 传入Print日志
func (d *DB) Print(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Print(log)
}

// Warn 传入Warn日志
func (d *DB) Warn(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Warn(log)
}

// Error 传入Error日志
func (d *DB) Error(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Error(log)
}

// Fatal 传入Fatal日志
func (d *DB) Fatal(str string) {
	log := &dbLog{
		Msg: str,
	}
	d.loger.Fatal(log)
}
