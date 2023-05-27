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
	// hold model `Dict` column name with table name(`dict`) prefix
	xx_Dict_Id_WithTableName        = xx_Dict_TableName + "_" + xx_Dict_Id
	xx_Dict_Key_WithTableName       = xx_Dict_TableName + "_" + xx_Dict_Key
	xx_Dict_Name_WithTableName      = xx_Dict_TableName + "_" + xx_Dict_Name
	xx_Dict_IsPin_WithTableName     = xx_Dict_TableName + "_" + xx_Dict_IsPin
	xx_Dict_Remark_WithTableName    = xx_Dict_TableName + "_" + xx_Dict_Remark
	xx_Dict_CreatedAt_WithTableName = xx_Dict_TableName + "_" + xx_Dict_CreatedAt
	xx_Dict_UpdatedAt_WithTableName = xx_Dict_TableName + "_" + xx_Dict_UpdatedAt
)

var xxx_Dict_Native_Model = new_X_Dict("")
var xxx_Dict_Model = new_X_Dict(xx_Dict_TableName)

type Dict_Active struct {
	// private fields
	xTableName string

	ALL       assist.Asterisk
	Id        assist.Int64
	Key       assist.String
	Name      assist.String
	IsPin     assist.Bool
	Remark    assist.String
	CreatedAt assist.Time
	UpdatedAt assist.Time
}

// X_Native_Dict native model without TableName.
func X_Native_Dict() Dict_Active {
	return xxx_Dict_Native_Model
}

// X_Dict model with TableName `dict`.
func X_Dict() Dict_Active {
	return xxx_Dict_Model
}

func new_X_Dict(xTableName string) Dict_Active {
	return Dict_Active{
		xTableName: xTableName,

		ALL: assist.NewAsterisk(xTableName),

		Id:        assist.NewInt64(xTableName, xx_Dict_Id),
		Key:       assist.NewString(xTableName, xx_Dict_Key),
		Name:      assist.NewString(xTableName, xx_Dict_Name),
		IsPin:     assist.NewBool(xTableName, xx_Dict_IsPin),
		Remark:    assist.NewString(xTableName, xx_Dict_Remark),
		CreatedAt: assist.NewTime(xTableName, xx_Dict_CreatedAt),
		UpdatedAt: assist.NewTime(xTableName, xx_Dict_UpdatedAt),
	}
}

// New_X_Dict new instance.
func New_X_Dict(xTableName string) Dict_Active {
	switch xTableName {
	case "":
		return xxx_Dict_Native_Model
	case xx_Dict_TableName:
		return xxx_Dict_Model
	default:
		return new_X_Dict(xTableName)
	}
}

// As alias
func (*Dict_Active) As(alias string) Dict_Active {
	return New_X_Dict(alias)
}

// X_TableName hold table name when call New_X_Dict or Dict_Active.As that you defined.
func (x *Dict_Active) X_TableName() string {
	return x.xTableName
}

// X_Model model
func (*Dict_Active) X_Model() *Dict {
	return &Dict{}
}

// X_TableName hold table name when call New_X_Dict or Dict_Active.As that you defined.
func (x *Dict_Active) X_Executor(db *gorm.DB) *assist.Executor[Dict] {
	return assist.NewExecutor[Dict](db)
}

// TableName hold model `Dict` table name returns `dict`.
func (*Dict_Active) TableName() string {
	return xx_Dict_TableName
}

// Field_Id hold model `Dict` column name.
// if prefixes not exist returns `id`, others `{prefixes[0]}_id`
func (*Dict_Active) Field_Id(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_Id
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_Id_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_Id
}

// Field_Key hold model `Dict` column name.
// if prefixes not exist returns `key`, others `{prefixes[0]}_key`
func (*Dict_Active) Field_Key(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_Key
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_Key_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_Key
}

// Field_Name hold model `Dict` column name.
// if prefixes not exist returns `name`, others `{prefixes[0]}_name`
func (*Dict_Active) Field_Name(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_Name
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_Name_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_Name
}

// Field_IsPin hold model `Dict` column name.
// if prefixes not exist returns `is_pin`, others `{prefixes[0]}_is_pin`
func (*Dict_Active) Field_IsPin(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_IsPin
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_IsPin_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_IsPin
}

// Field_Remark hold model `Dict` column name.
// if prefixes not exist returns `remark`, others `{prefixes[0]}_remark`
func (*Dict_Active) Field_Remark(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_Remark
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_Remark_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_Remark
}

// Field_CreatedAt hold model `Dict` column name.
// if prefixes not exist returns `created_at`, others `{prefixes[0]}_created_at`
func (*Dict_Active) Field_CreatedAt(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_CreatedAt
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_CreatedAt_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_CreatedAt
}

// Field_UpdatedAt hold model `Dict` column name.
// if prefixes not exist returns `updated_at`, others `{prefixes[0]}_updated_at`
func (*Dict_Active) Field_UpdatedAt(prefixes ...string) string {
	if len(prefixes) == 0 {
		return xx_Dict_UpdatedAt
	}
	if prefixes[0] == xx_Dict_TableName {
		return xx_Dict_UpdatedAt_WithTableName
	}
	return prefixes[0] + "_" + xx_Dict_UpdatedAt
}

func x_SelectDict(x *Dict_Active, prefixes ...string) []assist.Expr {
	if len(prefixes) > 0 {
		return []assist.Expr{
			x.Id.As(x.Field_Id(prefixes...)),
			x.Key.As(x.Field_Key(prefixes...)),
			x.Name.As(x.Field_Name(prefixes...)),
			x.IsPin.As(x.Field_IsPin(prefixes...)),
			x.Remark.As(x.Field_Remark(prefixes...)),
			x.CreatedAt.UnixTimestamp().As(x.Field_CreatedAt(prefixes...)),
			x.UpdatedAt.UnixTimestamp().As(x.Field_UpdatedAt(prefixes...)),
		}
	} else {
		return []assist.Expr{
			x.Id,
			x.Key,
			x.Name,
			x.IsPin,
			x.Remark,
			x.CreatedAt.UnixTimestamp().As(xx_Dict_CreatedAt),
			x.UpdatedAt.UnixTimestamp().As(xx_Dict_UpdatedAt),
		}
	}
}

// X_Native_SelectDict select field use use X_Native_Dict().
func X_Native_SelectDict() []assist.Expr {
	return x_SelectDict(&xxx_Dict_Native_Model)
}

// X_SelectDict select fields use X_Dict().
func X_SelectDict(prefixes ...string) []assist.Expr {
	return x_SelectDict(&xxx_Dict_Model, prefixes...)
}
