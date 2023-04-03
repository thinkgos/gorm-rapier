package dict

import (
	assist "github.com/things-go/gorm-assist"
)

type do struct {
	tablename string
}

func (d do) TableName() string {
	return d.tablename
}

type Dict struct {
	do
	Id        assist.Int64
	Key       assist.String
	Name      assist.String
	IsPin     assist.Bool
	Sort      assist.Uint16
	CreatedAt assist.Time

	TableName assist.String
}

func NewDict() Dict {
	tableName := "dict"
	return Dict{
		Id:        assist.NewInt64(tableName, "id"),
		Key:       assist.NewString(tableName, "key"),
		Name:      assist.NewString(tableName, "name"),
		IsPin:     assist.NewBool(tableName, "is_pin"),
		Sort:      assist.NewUint16(tableName, "sort"),
		CreatedAt: assist.NewTime(tableName, "created_at"),
	}
}

func (Dict) As(alias string) *Dict {
	return &Dict{
		Id:        assist.NewInt64(alias, "id"),
		Key:       assist.NewString(alias, "key"),
		Name:      assist.NewString(alias, "name"),
		IsPin:     assist.NewBool(alias, "is_pin"),
		Sort:      assist.NewUint16(alias, "sort"),
		CreatedAt: assist.NewTime(alias, "created_at"),
	}
}

func (d Dict) UseTable(tb string) {
	d.tablename = tb
}
