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

var {{$v.StructName}}_all_column = "{{range $k2, $v2 := $v.Columns}}{{if $k2}},{{else}}{{end}}{{$v2.Name}}{{end}}"
var {{$v.StructName}}_all_column_index = "{{range $k2, $v2 := $v.Columns}}{{if $k2}},{{else}}{{end}}${{add $k2}}{{end}}"

func (this *{{$v.StructName}}) Reflect(column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""

  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { {{range $_, $v2 := $v.Columns}}
    case "{{$v2.Name}}":
      rel = append(rel, &(this.{{$v2.StructName}})){{end}}
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

/*func (this *{{$v.StructName}}) ColumnReflect(column []string) (*[]interface{}, *string, error) {
  rel := ""
  rel_s :=make([]interface{}, 0, 10)

  for k, v := range column {
    rel = rel+v+"=$"+fmt.Sprintf("%d",k+1)+","
    switch v { {{range $_, $v2 := $v.Columns}}
    case "{{$v2.Name}}":
      rel_s = append(rel_s,this.{{$v2.StructName}}){{end}}
    default:
      return &rel_s, &rel, errors.New(v + ",字段不存在")
    }
  }
  rel = strings.TrimRight(rel, ",")
  return &rel_s, &rel,  nil
}*/

func (this *{{$v.StructName}}) AllReflect() (*[]interface{},*string,*string,error) {
  rel := &{{$v.StructName}}_all_column
  rel_str := &{{$v.StructName}}_all_column_index
  rel_s :=make([]interface{}, 0, 10)

  {{range $k2, $v2 := $v.Columns}}
  if this.{{$v2.StructName}} == {{column_def $v2.Type}} {
    this.{{$v2.StructName}} = {{column_def_replace $v2.Type $v2.Default}}
  }
  rel_s = append(rel_s,&(this.{{$v2.StructName}})){{end}}
  return &rel_s, rel, rel_str, nil
}

func (this *{{$v.StructName}}) TableName() string {
  return "{{$v.Name}}"
}

func (this *{{$v.StructName}}) AppendSelf(all interface{})error{
   all_data,ok := all.(*[]{{$v.StructName}})
   if ok==false{
    return errors.New("传入结构和表名不符")
   }
  *all_data = append(*all_data, *this)
  return nil
}


func {{$v.StructName}}New() postgres.ReflectInterface{
  return new({{$v.StructName}})
}

{{end}}


func test() {
    fmt.Println("start sqlmap")
}
`
