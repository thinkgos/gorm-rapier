package model

import (
	assist "github.com/things-go/gorm-assist"
	"gorm.io/gorm"
)

const (
	// hold model `Dict` table name
	xx_Dict_TableName = "dict"
	// hold model `Dict` column name
	xx_Dict_Id        = "id"
	xx_Dict_Key       = "key"
	xx_Dict_Name      = "name"
	xx_Dict_IsPin     = "is_pin"
	xx_Dict_Remark    = "remark"
	xx_Dict_CreatedAt = "created_at"
	xx_Dict_UpdatedAt = "updated_at"
)

var xxx_Dict_Native_Model = new_Dict("")
var xxx_Dict_Model = new_Dict(xx_Dict_TableName)

type Dict_Native struct {
	xAlias    string
	ALL       assist.Asterisk
	Id        assist.Int64
	Key       assist.String
	Name      assist.String
	IsPin     assist.Bool
	Remark    assist.String
	CreatedAt assist.Time
	UpdatedAt assist.Time
}

// X_Dict model with TableName `dict`.
func X_Dict() Dict_Native {
	return xxx_Dict_Model
}

func new_Dict(xAlias string) Dict_Native {
	return Dict_Native{
		xAlias:    xAlias,
		ALL:       assist.NewAsterisk(xAlias),
		Id:        assist.NewInt64(xAlias, xx_Dict_Id),
		Key:       assist.NewString(xAlias, xx_Dict_Key),
		Name:      assist.NewString(xAlias, xx_Dict_Name),
		IsPin:     assist.NewBool(xAlias, xx_Dict_IsPin),
		Remark:    assist.NewString(xAlias, xx_Dict_Remark),
		CreatedAt: assist.NewTime(xAlias, xx_Dict_CreatedAt),
		UpdatedAt: assist.NewTime(xAlias, xx_Dict_UpdatedAt),
	}
}

// New_Dict new instance.
func New_Dict(xAlias string) Dict_Native {
	switch xAlias {
	case "":
		return xxx_Dict_Native_Model
	case xx_Dict_TableName:
		return xxx_Dict_Model
	default:
		return new_Dict(xAlias)
	}
}

// As alias
func (*Dict_Native) As(alias string) Dict_Native {
	return New_Dict(alias)
}

// X_Alias hold table name when call New_Dict or Dict_Active.As that you defined.
func (x *Dict_Native) X_Alias() string {
	return x.xAlias
}

// New_Executor new entity executor which suggest use only once.
func (*Dict_Native) New_Executor(db *gorm.DB) *assist.Executor[Dict] {
	return assist.NewExecutor[Dict](db)
}

// TableName hold model `Dict` table name returns `dict`.
func (*Dict_Native) TableName() string {
	return xx_Dict_TableName
}

// Select_Expr select model fields
func (x *Dict_Native) Select_Expr() []assist.Expr {
	return []assist.Expr{
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
func (x *Dict_Native) Select_VariantExpr(prefixes ...string) []assist.Expr {
	if len(prefixes) > 0 {
		return []assist.Expr{
			x.Id.As(x.Id.FieldName(prefixes...)),
			x.Key.As(x.Key.FieldName(prefixes...)),
			x.Name.As(x.Name.FieldName(prefixes...)),
			x.IsPin.As(x.IsPin.FieldName(prefixes...)),
			x.Remark.As(x.Remark.FieldName(prefixes...)),
			x.CreatedAt.UnixTimestamp().As(x.CreatedAt.FieldName(prefixes...)),
			x.UpdatedAt.UnixTimestamp().As(x.UpdatedAt.FieldName(prefixes...)),
		}
	} else {
		return []assist.Expr{
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
