package postgres

// ReflectTable 数据库表接口
type ReflectTable interface {
	Reflect(column []string) (*[]interface{}, *string, error)
	AllReflect() (*[]interface{}, *string, *string, error)
	TableName() string
	AppendSelf(interface{}) error
}
