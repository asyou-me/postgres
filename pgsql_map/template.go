package main

var tmp = `package {{.Package}}

import (
  "errors"
  "github.com/asyoume/postgres"
  "strings"
  "fmt"
)

{{range $_, $v := .Tables}}type {{$v.StructName}} struct { {{range $_, $v2 := $v.Columns}}
    {{$v2.StructName}}   {{$v2.Type}} ` + "`" + `{{$v2.Tag}}` + "`" + `{{end}}
}

var {{$v.StructName}}_all_column = "{{range $k2, $v2 := $v.Columns}}{{$v2.Name}}{{if $k2}}{{else}},{{end}}{{end}}"
var {{$v.StructName}}_all_column_index = "{{range $k2, $v2 := $v.Columns}}${{add $k2}}{{if $k2}}{{else}},{{end}}{{end}}"

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

func {{$v.StructName}}UpdateReflect(s interface{}, column []string) (*string, *[]interface{}, error) {
  rel := ""
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*{{$v.StructName}})
  for k, v := range column {
    rel = rel+v+"=$"+fmt.Sprintf("%d",k+1)+","
    switch v { {{range $_, $v2 := $v.Columns}}
    case "{{$v2.Name}}":
      rel_s = append(rel_s,sr.{{$v2.StructName}}){{end}}
    default:
      return &rel, &rel_s, errors.New(v + ",字段不存在")
    }
  }
  rel = strings.TrimRight(rel, ",")
  return &rel, &rel_s, nil
}

func {{$v.StructName}}AllReflect(s interface{}) (*string,*string, *[]interface{}, error) {
  rel := &{{$v.StructName}}_all_column
  rel_str := &{{$v.StructName}}_all_column_index
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*{{$v.StructName}})
  {{range $k2, $v2 := $v.Columns}}
  if sr.{{$v2.StructName}} != {{column_def $v2.Type}} {
    rel_s = append(rel_s,sr.{{$v2.StructName}})
  }else{
    rel_s = append(rel_s,{{column_def_replace $v2.Type $v2.Default}})
  }{{end}}
  return rel, rel_str,&rel_s, nil
}

func {{$v.StructName}}New() interface{}{
  return &{{$v.StructName}}{}
}

func {{$v.StructName}}Check(s interface{}) bool{
  _, ok := s.(*{{$v.StructName}})
  return ok
}

func {{$v.StructName}}Check2(s interface{}) bool{
  _, ok := s.(*[]{{$v.StructName}})
  return ok
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
  {{range $_, $v := .Tables}}postgres.SqlCheckMap["{{$v.Name}}"] = {{$v.StructName}}Check{{end}}
  {{range $_, $v := .Tables}}postgres.SqlCheck2Map["{{$v.Name}}"] = {{$v.StructName}}Check2{{end}}
  {{range $_, $v := .Tables}}postgres.AllReflectMap["{{$v.Name}}"] = {{$v.StructName}}AllReflect{{end}}
  {{range $_, $v := .Tables}}postgres.UpdateReflectMap["{{$v.Name}}"] = {{$v.StructName}}UpdateReflect{{end}}
}

func test() {
    fmt.Println("start sqlmap")
}
`
