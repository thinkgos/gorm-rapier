package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Condition func(*gorm.DB) *gorm.DB

func Select(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db.Clauses(clause.Select{})
		}
		query, args := buildSelectValue(db.Statement, columns...)
		return db.Select(query, args...)
	}
}

// Order with field
func Order(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Order(buildColumnsValue(db, columns...))
	}
}

// Group with field
func Group(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Group(buildColumnsValue(db, columns...))
	}
}

