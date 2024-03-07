package testdata

import (
	"time"
)

// Dict 字典
type Dict struct {
	Id        int64     `gorm:"column:id;autoIncrement:true;not null;primaryKey" json:"id,omitempty"`
	Key       string    `gorm:"column:key;type:varchar(64);not null;default:'';uniqueIndex:uk_key" json:"key,omitempty"`
	Name      string    `gorm:"column:name;type:varchar(64);not null;default:''" json:"name,omitempty"`
	IsPin     bool      `gorm:"column:is_pin;type:tinyint(1);not null;default:0" json:"is_pin,omitempty"`
	Remark    string    `gorm:"column:remark;type:varchar(128);not null;default:''" json:"remark,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty"`
}

// TableName implement schema.Tabler interface
func (*Dict) TableName() string {
	return "dict"
}
