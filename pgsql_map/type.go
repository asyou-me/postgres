package main

import (
	"strings"
)

type InputType struct {
	Package string       `json:"package"`
	Tables  []InputTable `json:"tables"`
	Decorat string       `json:"decorat"`
}

type InputTable struct {
	Name          string   `json:"name"`
	StructName    string   `json:"struct_name"`
	ColumnsString []string `json:"columns"`
	Columns       []InputColumn
}

type InputColumn struct {
	Name       string `json:"name"`
	StructName string `json:"struct_name"`
	Tag        string `json:"tag"`
	Type       string `json:"type"`
	Default    string `json:"default"`
}

func ColumnsStringToCloumn(obj *InputType) {
	for k1, table := range obj.Tables {
		for _, v := range table.ColumnsString {
			c := strings.Split(v, ",")
			obj.Tables[k1].Columns = append(obj.Tables[k1].Columns, InputColumn{
				Name:       c[0],
				StructName: c[1],
				Type:       c[2],
				Default:    c[3],
				Tag:        c[4],
			})
		}
	}
}
