package main

var tmp = `package {{.Package}}

import (
  "errors"
  "github.com/asyoume/postgres"
  "strings"
)

{{range $_, $v := .Tables}}type {{$v.StructName}} struct { {{range $_, $v2 := $v.Columns}}
    {{$v2.StructName}}   {{$v2.Type}} ` + "`" + `{{$v2.Tag}}` + "`" + `{{end}}
}

func {{$v.StructName}}Reflect(s interface{}, column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""
  sr := s.(*{{$v.StructName}})
  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { {{range $_, $v2 := $v.Columns}}
    case "{{$v2.Name}}":
      rel = append(rel, &(sr.{{$v2.StructName}})){{end}}
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func {{$v.StructName}}New() interface{}{
  return &{{$v.StructName}}{}
}

func {{$v.StructName}}Add(all interface{},s interface{}){
   all_data := all.(*[]{{$v.StructName}})
   sr := *s.(*{{$v.StructName}})

   new_sr := {{$v.StructName}}{}
   new_sr = sr
  *all_data = append(*all_data, new_sr)
}{{end}}

func init() {
  {{range $_, $v := .Tables}}postgres.SqlFuncMap["{{$v.Name}}"] = {{$v.StructName}}Reflect{{end}}
  {{range $_, $v := .Tables}}postgres.SqlNewMap["{{$v.Name}}"] = {{$v.StructName}}New{{end}}
  {{range $_, $v := .Tables}}postgres.SqlAddMap["{{$v.Name}}"] = {{$v.StructName}}Add{{end}}
}
`
