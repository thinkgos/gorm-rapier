package model

import (
	assist "github.com/things-go/gorm-assist"
	"gorm.io/gorm"
)

type Dict_Entity struct {
	db *gorm.DB
}

func New_Dict(db *gorm.DB) *Dict_Entity {
	return &Dict_Entity{
		db: db,
	}
}

// Executor new executor
func (x *Dict_Entity) Executor() *assist.Executor[Dict] {
	return assist.NewExecutor[Dict](x.db)
}
