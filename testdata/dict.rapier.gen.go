package testdata

import (
	rapier "github.com/thinkgos/gorm-rapier"
)

var ref_Dict_Native = New_Dict("dict")

type Dict_Native struct {
	refAlias     string
	refTableName string
	ALL          rapier.Asterisk
	Id           rapier.Int64
	Key          rapier.String
	Name         rapier.String
	IsPin        rapier.Bool
	Remark       rapier.String
	CreatedAt    rapier.Time
	UpdatedAt    rapier.Time
}

func new_Dict(tableName, alias string) *Dict_Native {
	return &Dict_Native{
		refAlias:     alias,
		refTableName: tableName,
		ALL:          rapier.NewAsterisk(alias),
		Id:           rapier.NewInt64(alias, "id"),
		Key:          rapier.NewString(alias, "key"),
		Name:         rapier.NewString(alias, "name"),
		IsPin:        rapier.NewBool(alias, "is_pin"),
		Remark:       rapier.NewString(alias, "remark"),
		CreatedAt:    rapier.NewTime(alias, "created_at"),
		UpdatedAt:    rapier.NewTime(alias, "updated_at"),
	}
}

// Ref_Dict model with TableName `dict`.
func Ref_Dict() *Dict_Native { return ref_Dict_Native }

// New_Dict new instance.
func New_Dict(tableName string) *Dict_Native {
	return new_Dict(tableName, tableName)
}

// As alias
func (x *Dict_Native) As(alias string) *Dict_Native {
	return new_Dict(x.refTableName, alias)
}

// Ref_Alias hold alias name when call Dict_Active.As that you defined.
func (x *Dict_Native) Ref_Alias() string { return x.refAlias }

// TableName hold table name when call New_Dict that you defined.
func (x *Dict_Native) TableName() string { return x.refTableName }
