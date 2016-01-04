package main

import (
	"text/template"
)

var template_func = template.FuncMap{
	"column": func(name, ty string) string {
		rel := ""
		switch ty {
		case "string":
			rel = rel + "\"'\"+sr." + name + "+\"'\""
		case "int64":
			rel = rel + "fmt.Sprintf(\"%d\",sr." + name + ")"
		case "int32":
			rel = rel + "fmt.Sprintf(\"%d\",sr." + name + ")"
		case "int16":
			rel = rel + "fmt.Sprintf(\"%d\",sr." + name + ")"
		case "int8":
			rel = rel + "fmt.Sprintf(\"%d\",sr." + name + ")"
		default:

		}
		return rel
	}, "column_def": func(ty string) string {
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
		default:
			return ""
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
		default:
			return ""
		}
	}, "add": func(i int) int {
		return i + 1
	},
}
