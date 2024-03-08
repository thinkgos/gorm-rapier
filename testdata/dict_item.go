package testdata

// Dict 字典
type DictItem struct {
	Id        int64  `gorm:"column:id;autoIncrement:true;not null;primaryKey" json:"id,omitempty"`
	DictId    int64  `gorm:"column:dict_id;not null;default:0" json:"dict_id,omitempty"`
	Name      string `gorm:"column:name;type:varchar(64);not null;default:''" json:"name,omitempty"`
	Sort      uint32 `gorm:"column:sort;not null;default:0" json:"sort,omitempty"`
	IsEnabled bool   `gorm:"column:is_enabled;type:tinyint(1);not null;default:0" json:"is_enabled,omitempty"`
}

// TableName implement schema.Tabler interface
func (*DictItem) TableName() string {
	return "dict_item"
}
