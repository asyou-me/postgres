package exmples

import (
	"github.com/asyou-me/postgres"

	pulic_type "github.com/asyou-me/lib.v1/pulic_type"
)

var (
	DB *postgres.DB
)

func Init(conf *pulic_type.MicroSerType, loger pulic_type.Logger) error {
	var err error
	DB, err = postgres.NewDB(conf, loger)
	DB.TableMap["test"] = NewTest
	return err
}
