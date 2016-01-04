package models

import (
  "errors"
  "github.com/asyoume/postgres"
  "strings"
  "fmt"
)

type Test1 struct { 
    Test1   string `json:"test1"`
    Test2   int64 `json:"test2"`
}

var Test1_all_column = "test1,test2"
var Test1_all_column_index = "$1,$2"

func Test1Reflect(s interface{}, column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""
  sr := s.(*Test1)
  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "test1":
      rel = append(rel, &(sr.Test1))
    case "test2":
      rel = append(rel, &(sr.Test2))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func Test1UpdateReflect(s interface{}, column []string) (*string, *[]interface{}, error) {
  rel := ""
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*Test1)
  for k, v := range column {
    rel = rel+v+"=$"+fmt.Sprintf("%d",k+1)+","
    switch v { 
    case "test1":
      rel_s = append(rel_s,sr.Test1)
    case "test2":
      rel_s = append(rel_s,sr.Test2)
    default:
      return &rel, &rel_s, errors.New(v + ",字段不存在")
    }
  }
  rel = strings.TrimRight(rel, ",")
  return &rel, &rel_s, nil
}

func Test1AllReflect(s interface{}) (*string,*string, *[]interface{}, error) {
  rel := &Test1_all_column
  rel_str := &Test1_all_column_index
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*Test1)
  
  if sr.Test1 != "" {
    rel_s = append(rel_s,sr.Test1)
  }else{
    rel_s = append(rel_s,"")
  }
  if sr.Test2 != 0 {
    rel_s = append(rel_s,sr.Test2)
  }else{
    rel_s = append(rel_s,0)
  }
  return rel, rel_str,&rel_s, nil
}

func Test1New() interface{}{
  return &Test1{}
}

func Test1Check(s interface{}) bool{
  _, ok := s.(*Test1)
  return ok
}

func Test1Check2(s interface{}) bool{
  _, ok := s.(*[]Test1)
  return ok
}

func Test1Add(all interface{},s interface{}){
   all_data := all.(*[]Test1)
   sr := *s.(*Test1)

   new_sr := Test1{}
   new_sr = sr
  *all_data = append(*all_data, new_sr)
}

func init() {
  postgres.SqlFuncMap["test1"] = Test1Reflect
  postgres.SqlNewMap["test1"] = Test1New
  postgres.SqlAddMap["test1"] = Test1Add
  postgres.SqlCheckMap["test1"] = Test1Check
  postgres.SqlCheck2Map["test1"] = Test1Check2
  postgres.AllReflectMap["test1"] = Test1AllReflect
  postgres.UpdateReflectMap["test1"] = Test1UpdateReflect
}
func test() {
    fmt.Println("start sqlmap")
}
