package postgres

import (
	"time"

	"github.com/jackc/pgx"

	pulic_type "github.com/asyoume/lib.v1/pulic_type"
)

type NewFunc func() ReflectTableInterface

var (
	AllColumn = []string{}
)

type DB struct {
	//数据操作连接
	Pool     *pgx.ConnPool
	loger    pulic_type.Logger
	TableMap map[string]NewFunc
}

func (d *DB) Open(conf *pulic_type.MicroSerType, loger pulic_type.Logger) error {
	//初始化数据库
	var err error

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

	d.Pool, err = pgx.NewConnPool(connConfig)
	//传入外部日志模块
	d.loger = loger
	return err
}

//定义新建数据库操作对象的当法
func NewDB(conf *pulic_type.MicroSerType, loger pulic_type.Logger) (*DB, error) {
	db := DB{}
	db.TableMap = map[string]NewFunc{}
	err := db.Open(conf, loger)
	return &db, err
}
