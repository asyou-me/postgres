package exmples

import (
	"errors"
	"fmt"
	"testing"
	//"time"

	pulic_type "github.com/asyou-me/lib.v1/pulic_type"
	"github.com/asyou-me/postgres"
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

	// 测试插入数据
	_, err = DB.Insert(&Test{
		D:    &map[string]string{"xiaobai": "zheshi", "qq": "422145328"},
		Nick: "xiaobai",
	})
	if err != nil {
		t.Error(err)
	}

	data2 := &Test{}
	// 查询数据
	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Scan(data2)
	if err != nil || (*data2.D)["xiaobai"] != "zheshi" {
		t.Error(err)
	}

	// 查询列表数据
	dataList := &[]Test{}
	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Scans(dataList, 1, 10)
	if err != nil {
		t.Error(err)
	}

	// 设置部分字段
	err = DB.Table("test").Where(`d@>'{"xiaobai": "zheshi"}'`).Set([]postgres.GSTYPE{
		postgres.GSTYPE{
			Key:   "d",
			Path:  "qq",
			Value: "\"4228\"",
		}, postgres.GSTYPE{
			Key:   "nick",
			Value: "xiaobai1",
		},
	})
	if err != nil {
		t.Error(err)
	}

	data3 := []postgres.GSTYPE{
		postgres.GSTYPE{
			Key:  "d",
			Path: "xiaobai,4",
		}, postgres.GSTYPE{
			Key: "nick",
		},
	}

	// 获取部分字段
	err = DB.Table("test").Get(data3)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("get data succeed:", data3)
	}

	if data3[0].Value != "xiugaihou" {
		t.Error(errors.New("Get error"))
	}

	err = DB.Del("test", `d@>'{"xiaobai": "xiugaihou"}'`)
	if err != nil {
		t.Error(err)
	}
}
