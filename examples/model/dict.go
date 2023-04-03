package model

import "time"

// Dict 字典
type Dict struct {
	Id        int64     `gorm:"column:id;autoIncrement:true;not null;primaryKey;comment:字典ID" json:"id,omitempty"`                   // 字典ID
	Key       string    `gorm:"column:key;type:varchar(64);not null;default:'';uniqueIndex:uk_key;comment:关键字" json:"key,omitempty"` // 关键字
	Name      string    `gorm:"column:name;type:varchar(64);not null;default:'';comment:名称" json:"name,omitempty"`                   // 名称
	IsPin     bool      `gorm:"column:is_pin;type:tinyint(1);not null;comment:是否锁定, 一旦锁定将不可删除" json:"is_pin,omitempty"`              // 是否锁定, 一旦锁定将不可删除
	Sort      uint16    `gorm:"column:sort;type:smallint unsigned;not null;default:0;comment:序号" json:"sort,omitempty"`              // 备注
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at,omitempty"`                   // 创建时间
	// UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at,omitempty"`                   // 更新时间
}

// TableName implement schema.Tabler interface
func (*Dict) TableName() string {
	return "dict"
}
