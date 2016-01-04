package main

import ()

type InputType struct {
	Package string       `json:"package"`
	Tables  []InputTable `json:"tables"`
}

type InputTable struct {
	Name       string        `json:"name"`
	StructName string        `json:"struct_name"`
	Columns    []InputColumn `json:"columns"`
}

type InputColumn struct {
	Name       string `json:"name"`
	StructName string `json:"struct_name"`
	Tag        string `json:"tag"`
	Type       string `json:"type"`
	Default    string `json:"default"`
}
