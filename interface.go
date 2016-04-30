package postgres

type ReflectInterface interface {
	Reflect(column []string) (*[]interface{}, *string, error)
	AllReflect() (*[]interface{}, *string, *string, error)
	TableName() string
	AppendSelf(interface{}) error
}
