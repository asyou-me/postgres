package models

import (
  "errors"
  "github.com/asyoume/postgres"
  "strings"
  "fmt"
)

type User struct { 
    Id   string `json:"id"`
    Nick   string `json:"nick"`
    Email   string `json:"email"`
    Passwd   string `json:"passwd"`
    created   int64 `json:"created"`
    Updated   int64 `json:"updated"`
    Deleted   int64 `json:"deleted"`
    Active   bool `json:"active"`
    Attr   string `json:"attr"`
}

var User_all_column = "id,nick,email,passwd,created,updated,deleted,active,attr"
var User_all_column_index = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

func UserReflect(s interface{}, column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""
  sr := s.(*User)
  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "id":
      rel = append(rel, &(sr.Id))
    case "nick":
      rel = append(rel, &(sr.Nick))
    case "email":
      rel = append(rel, &(sr.Email))
    case "passwd":
      rel = append(rel, &(sr.Passwd))
    case "created":
      rel = append(rel, &(sr.created))
    case "updated":
      rel = append(rel, &(sr.Updated))
    case "deleted":
      rel = append(rel, &(sr.Deleted))
    case "active":
      rel = append(rel, &(sr.Active))
    case "attr":
      rel = append(rel, &(sr.Attr))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func UserUpdateReflect(s interface{}, column []string) (*string, *[]interface{}, error) {
  rel := ""
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*User)
  for k, v := range column {
    rel = rel+v+"=$"+fmt.Sprintf("%d",k+1)+","
    switch v { 
    case "id":
      rel_s = append(rel_s,sr.Id)
    case "nick":
      rel_s = append(rel_s,sr.Nick)
    case "email":
      rel_s = append(rel_s,sr.Email)
    case "passwd":
      rel_s = append(rel_s,sr.Passwd)
    case "created":
      rel_s = append(rel_s,sr.created)
    case "updated":
      rel_s = append(rel_s,sr.Updated)
    case "deleted":
      rel_s = append(rel_s,sr.Deleted)
    case "active":
      rel_s = append(rel_s,sr.Active)
    case "attr":
      rel_s = append(rel_s,sr.Attr)
    default:
      return &rel, &rel_s, errors.New(v + ",字段不存在")
    }
  }
  rel = strings.TrimRight(rel, ",")
  return &rel, &rel_s, nil
}

func UserAllReflect(s interface{}) (*string,*string, *[]interface{}, error) {
  rel := &User_all_column
  rel_str := &User_all_column_index
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*User)
  
  if sr.Id != "" {
    rel_s = append(rel_s,sr.Id)
  }else{
    rel_s = append(rel_s,"default")
  }
  if sr.Nick != "" {
    rel_s = append(rel_s,sr.Nick)
  }else{
    rel_s = append(rel_s,"")
  }
  if sr.Email != "" {
    rel_s = append(rel_s,sr.Email)
  }else{
    rel_s = append(rel_s,"")
  }
  if sr.Passwd != "" {
    rel_s = append(rel_s,sr.Passwd)
  }else{
    rel_s = append(rel_s,"")
  }
  if sr.created != 0 {
    rel_s = append(rel_s,sr.created)
  }else{
    rel_s = append(rel_s,0)
  }
  if sr.Updated != 0 {
    rel_s = append(rel_s,sr.Updated)
  }else{
    rel_s = append(rel_s,0)
  }
  if sr.Deleted != 0 {
    rel_s = append(rel_s,sr.Deleted)
  }else{
    rel_s = append(rel_s,0)
  }
  if sr.Active != false {
    rel_s = append(rel_s,sr.Active)
  }else{
    rel_s = append(rel_s,true)
  }
  if sr.Attr != "" {
    rel_s = append(rel_s,sr.Attr)
  }else{
    rel_s = append(rel_s,"{}")
  }
  return rel, rel_str,&rel_s, nil
}

func UserNewReflect() interface{}{
  return &User{}
}

func UserCheckReflect(s interface{}) bool{
  _, ok := s.(*User)
  return ok
}

func UserCheck2Reflect(s interface{}) bool{
  _, ok := s.(*[]User)
  return ok
}

func UserAddReflect(all interface{},s interface{}){
   all_data := all.(*[]User)
   sr := *s.(*User)

   new_sr := User{}
   new_sr = sr
  *all_data = append(*all_data, new_sr)
}
type Task struct { 
    D   string `json:"d"`
}

var Task_all_column = "d"
var Task_all_column_index = "$1"

func TaskReflect(s interface{}, column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""
  sr := s.(*Task)
  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "d":
      rel = append(rel, &(sr.D))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func TaskUpdateReflect(s interface{}, column []string) (*string, *[]interface{}, error) {
  rel := ""
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*Task)
  for k, v := range column {
    rel = rel+v+"=$"+fmt.Sprintf("%d",k+1)+","
    switch v { 
    case "d":
      rel_s = append(rel_s,sr.D)
    default:
      return &rel, &rel_s, errors.New(v + ",字段不存在")
    }
  }
  rel = strings.TrimRight(rel, ",")
  return &rel, &rel_s, nil
}

func TaskAllReflect(s interface{}) (*string,*string, *[]interface{}, error) {
  rel := &Task_all_column
  rel_str := &Task_all_column_index
  rel_s :=make([]interface{}, 0, 10)
  sr := s.(*Task)
  
  if sr.D != "" {
    rel_s = append(rel_s,sr.D)
  }else{
    rel_s = append(rel_s,"{}")
  }
  return rel, rel_str,&rel_s, nil
}

func TaskNewReflect() interface{}{
  return &Task{}
}

func TaskCheckReflect(s interface{}) bool{
  _, ok := s.(*Task)
  return ok
}

func TaskCheck2Reflect(s interface{}) bool{
  _, ok := s.(*[]Task)
  return ok
}

func TaskAddReflect(all interface{},s interface{}){
   all_data := all.(*[]Task)
   sr := *s.(*Task)

   new_sr := Task{}
   new_sr = sr
  *all_data = append(*all_data, new_sr)
}


func NewDB() *postgres.DB {
  db := postgres.NewDB()
  
  db.SqlFuncMap["user"] = UserReflect
  db.SqlFuncMap["task"] = TaskReflect
  
  db.SqlNewMap["user"] = UserNewReflect
  db.SqlNewMap["task"] = TaskNewReflect
  
  db.SqlAddMap["user"] = UserAddReflect
  db.SqlAddMap["task"] = TaskAddReflect
  
  db.SqlCheckMap["user"] = UserCheckReflect
  db.SqlCheckMap["task"] = TaskCheckReflect
  
  db.SqlCheck2Map["user"] = UserCheck2Reflect
  db.SqlCheck2Map["task"] = TaskCheck2Reflect
  
  db.AllReflectMap["user"] = UserAllReflect
  db.AllReflectMap["task"] = TaskAllReflect
  
  db.UpdateReflectMap["user"] = UserUpdateReflect
  db.UpdateReflectMap["task"] = TaskUpdateReflect
  return db
}

func test() {
    fmt.Println("start sqlmap")
}
