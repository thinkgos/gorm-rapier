package model

import (
	rapier "github.com/thinkgos/gorm-rapier"
	"gorm.io/gorm"
)

var ref_Dict_Native = new_Dict("dict", "dict")

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

// Ref_Alias hold table name when call New_Dict or Dict_Active.As that you defined.
func (x *Dict_Native) Ref_Alias() string { return x.refAlias }

// TableName hold model `Dict` table name returns `dict`.
func (x *Dict_Native) TableName() string { return x.refTableName }

// New_Executor new entity executor which suggest use only once.
func (*Dict_Native) New_Executor(db *gorm.DB) *rapier.Executor[Dict] {
	return rapier.NewExecutor[Dict](db)
}

// Select_Expr select model fields
func (x *Dict_Native) Select_Expr() []rapier.Expr {
	return []rapier.Expr{
		x.Id,
		x.Key,
		x.Name,
		x.IsPin,
		x.Remark,
		x.CreatedAt,
		x.UpdatedAt,
	}
}

// Select_VariantExpr select model fields, but time.Time field convert to timestamp(int64).
func (x *Dict_Native) Select_VariantExpr(prefixes ...string) []rapier.Expr {
	if len(prefixes) > 0 && prefixes[0] != "" {
		return []rapier.Expr{
			x.Id.As(x.Id.FieldName(prefixes...)),
			x.Key.As(x.Key.FieldName(prefixes...)),
			x.Name.As(x.Name.FieldName(prefixes...)),
			x.IsPin.As(x.IsPin.FieldName(prefixes...)),
			x.Remark.As(x.Remark.FieldName(prefixes...)),
			x.CreatedAt.UnixTimestamp().As(x.CreatedAt.FieldName(prefixes...)),
			x.UpdatedAt.UnixTimestamp().As(x.UpdatedAt.FieldName(prefixes...)),
		}
	} else {
		return []rapier.Expr{
			x.Id,
			x.Key,
			x.Name,
			x.IsPin,
			x.Remark,
			x.CreatedAt.UnixTimestamp().As(x.CreatedAt.FieldName()),
			x.UpdatedAt.UnixTimestamp().As(x.UpdatedAt.FieldName()),
		}
	}
}
