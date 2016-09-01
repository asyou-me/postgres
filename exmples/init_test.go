package exmples

import (
	"errors"
	"testing"
	//"time"

	pulic_type "github.com/asyoume/lib.v1/pulic_type"
	"github.com/asyoume/postgres"
)

var logger pulic_type.Logger = &pulic_type.DefalutLogger{}

var connConfig = pulic_type.MicroSerType{
	Addr:   "jxspy.com",
	Id:     "postgres",
	Secret: "Jx201501",

	Attr: map[string]interface{}{
		"Database":       "test",
		"MaxConnections": 100,
		"AcquireTimeout": 10,
	},
}

func TestInit(t *testing.T) {
	err := Init(&connConfig, logger)
	if err != nil {
		t.Error(err)
	}

	_, err = DB.Insert(&Test{D: &map[string]string{"xiaobai": "zheshi"}})
	if err != nil {
		t.Log(err)
	}

	data2 := &Test{}
	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Scan(data2)
	if err != nil || (*data2.D)["xiaobai"] != "zheshi" {
		t.Log(err)
	}

	dataList := &[]Test{}
	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Scans(dataList, 1, 10)
	if err != nil {
		t.Log(err)
	}

	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Set([]postgres.GSTYPE{
		postgres.GSTYPE{
			Path:  "xiaobai",
			Key:   "d",
			Value: "\"xiugaihou\"",
		},
	})
	if err != nil {
		t.Log(err)
	}

	data3 := []postgres.GSTYPE{
		postgres.GSTYPE{
			Path:  "xiaobai",
			Key:   "d",
			Value: "",
		},
	}
	err = DB.Table("test").Get(data3)
	if err != nil {
		t.Log(err)
	}

	if data3[0].Value != "xiugaihou" {
		t.Log(errors.New("Get error"))
	}

	err = DB.Del("test", `d@>'{"xiaobai": "zheshi"}'`)
	if err != nil {
		t.Log(err)
	}
}
