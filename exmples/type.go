package exmples

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

var User_all_column = "'id','nick','email','passwd','created','updated','deleted','active','attr'"
var User_all_column_index = "$1,$2,$3,$4,$5,$6,$7,$8,$9"

func (this *User) Reflect(column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""

  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "id":
      rel = append(rel, &(this.Id))
    case "nick":
      rel = append(rel, &(this.Nick))
    case "email":
      rel = append(rel, &(this.Email))
    case "passwd":
      rel = append(rel, &(this.Passwd))
    case "created":
      rel = append(rel, &(this.created))
    case "updated":
      rel = append(rel, &(this.Updated))
    case "deleted":
      rel = append(rel, &(this.Deleted))
    case "active":
      rel = append(rel, &(this.Active))
    case "attr":
      rel = append(rel, &(this.Attr))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func (this *User) AllReflect() (*[]interface{},*string,*string,error) {
  rel := &User_all_column
  rel_str := &User_all_column_index
  rel_s :=make([]interface{}, 0, 10)

  
  if this.Id == "" {
    this.Id = "default"
  }
  rel_s = append(rel_s,&(this.Id))
  if this.Nick == "" {
    this.Nick = ""
  }
  rel_s = append(rel_s,&(this.Nick))
  if this.Email == "" {
    this.Email = ""
  }
  rel_s = append(rel_s,&(this.Email))
  if this.Passwd == "" {
    this.Passwd = ""
  }
  rel_s = append(rel_s,&(this.Passwd))
  if this.created == 0 {
    this.created = 0
  }
  rel_s = append(rel_s,&(this.created))
  if this.Updated == 0 {
    this.Updated = 0
  }
  rel_s = append(rel_s,&(this.Updated))
  if this.Deleted == 0 {
    this.Deleted = 0
  }
  rel_s = append(rel_s,&(this.Deleted))
  if this.Active == false {
    this.Active = true
  }
  rel_s = append(rel_s,&(this.Active))
  if this.Attr == "" {
    this.Attr = "{}"
  }
  rel_s = append(rel_s,&(this.Attr))
  return &rel_s, rel, rel_str, nil
}

func (this *User) TableName() string {
  return "user"
}

func (this *User) AppendSelf(all interface{})error{
   all_data,ok := all.(*[]User)
   if ok==false{
    all_data2,ok2 := all.(*[]*User)
    if ok2==false {
      return errors.New("传入结构和表名不符")
    }
    *all_data2 = append(*all_data2, this)
    return nil
   }
  *all_data = append(*all_data, *this)
  return nil
}

func NewUser() postgres.ReflectTableInterface{
  return new(User)
}

func UserTest() {
    fmt.Println("start sqlmap")
}


var Test_all_column = "'d'"
var Test_all_column_index = "$1"

func (this *Test) Reflect(column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""

  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "d":
      rel = append(rel, &(this.D))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func (this *Test) AllReflect() (*[]interface{},*string,*string,error) {
  rel := &Test_all_column
  rel_str := &Test_all_column_index
  rel_s :=make([]interface{}, 0, 10)

  
  if this.D == nil {
    this.D = &map[string]string{}
  }
  rel_s = append(rel_s,&(this.D))
  return &rel_s, rel, rel_str, nil
}

func (this *Test) TableName() string {
  return "test"
}

func (this *Test) AppendSelf(all interface{})error{
   all_data,ok := all.(*[]Test)
   if ok==false{
    all_data2,ok2 := all.(*[]*Test)
    if ok2==false {
      return errors.New("传入结构和表名不符")
    }
    *all_data2 = append(*all_data2, this)
    return nil
   }
  *all_data = append(*all_data, *this)
  return nil
}

func NewTest() postgres.ReflectTableInterface{
  return new(Test)
}

func TestTest() {
    fmt.Println("start sqlmap")
}




