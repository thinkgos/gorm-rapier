package rapier

import (
	"gorm.io/gorm"
)

func GromModel[T any]() Condition {
	return func(db *gorm.DB) *gorm.DB {
		var t T

		db = db.Model(&t)
		err := db.Statement.Parse(t)
		if err != nil {
			_ = db.AddError(err)
		}
		return db
	}
}

func GormTable(name string, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(name, args...)
	}
}
