package main

import (
	"text/template"
)

var template_func = template.FuncMap{
	"column_def": func(ty string) string {
		switch ty {
		case "string":
			return "\"\""
		case "int64":
			return "0"
		case "int32":
			return "0"
		case "int16":
			return "0"
		case "int8":
			return "0"
		case "float32":
			return "0"
		case "float64":
			return "0"
		case "bool":
			return "false"
		default:
			return "nil"
		}
	}, "column_def_replace": func(ty string, def string) string {
		switch ty {
		case "string":
			return "\"" + def + "\""
		case "int64":
			return def
		case "int32":
			return def
		case "int16":
			return def
		case "int8":
			return def
		case "float32":
			return def
		case "float64":
			return def
		case "bool":
			return def
		default:
			return def
		}
	}, "add": func(i int) int {
		return i + 1
	},
}
