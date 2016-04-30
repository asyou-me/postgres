package exmples

import (
	"fmt"
	"testing"
	//"time"

	pulic_type "github.com/asyoume/lib/pulic_type"
)

var connConfig = pulic_type.MicroSerType{
	Addr:   "xxxx",
	Id:     "xxxx",
	Secret: "xxxx",

	Attr: map[string]interface{}{
		"Database":       "test",
		"MaxConnections": 100,
		"AcquireTimeout": 10,
	},
}

type Loger struct {
}

func (this *Loger) Debug(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Info(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Print(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Warn(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Warning(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Error(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Fatal(args ...interface{}) {
	fmt.Println(args)
}

func (this *Loger) Panic(args ...interface{}) {
	fmt.Println(args)
}

func TestInit(t *testing.T) {
	err := Init(&connConfig, &Loger{})
	if err != nil {
		t.Error(err)
	}

	err = DB.Insert("test", &Test{D: &map[string]string{"xiaobai": "zheshi"}})
	fmt.Println("err:", err)

	data2 := &Test{}
	err = DB.One("test", ` WHERE d@>'{"xiaobai": "zheshi"}'`, data2)
	fmt.Println("err:", err)
	fmt.Println(data2.D)

	dataList := &[]Test{}
	err = DB.All("test", ` WHERE d@>'{"xiaobai": "zheshi"}'`, dataList, "", 1, 10)
	fmt.Println("err:", err)
	fmt.Println(dataList)
}
