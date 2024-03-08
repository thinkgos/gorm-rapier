package testdata

import rapier "github.com/thinkgos/gorm-rapier"

var ref_DictItem_Native = New_DictItem("dict_item")

type DictItem_Native struct {
	refAlias     string
	refTableName string
	ALL          rapier.Asterisk
	Id           rapier.Int64
	DictId       rapier.Int64
	Name         rapier.String
	Sort         rapier.Uint32
	IsPin        rapier.Bool
}

func new_DictItem(tableName, alias string) *DictItem_Native {
	return &DictItem_Native{
		refAlias:     alias,
		refTableName: tableName,
		ALL:          rapier.NewAsterisk(alias),
		Id:           rapier.NewInt64(alias, "id"),
		DictId:       rapier.NewInt64(alias, "dict_id"),
		Name:         rapier.NewString(alias, "name"),
		IsPin:        rapier.NewBool(alias, "is_pin"),
		Sort:         rapier.NewUint32(alias, "sort"),
	}
}

// Ref_Dict model with TableName `dict_item`.
func Ref_DictItem() *DictItem_Native { return ref_DictItem_Native }

// New_DictItem new instance.
func New_DictItem(tableName string) *DictItem_Native {
	return new_DictItem(tableName, tableName)
}

// As alias
func (x *DictItem_Native) As(alias string) *DictItem_Native {
	return new_DictItem(x.refTableName, alias)
}

// Ref_Alias hold alias name when call DictItem_Active.As that you defined.
func (x *DictItem_Native) Ref_Alias() string { return x.refAlias }

// TableName hold table name when call New_DictItem that you defined.
func (x *DictItem_Native) TableName() string { return x.refTableName }
