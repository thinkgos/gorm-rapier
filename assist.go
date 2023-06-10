package assist

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Condition alias func(*gorm.DB) *gorm.DB
type Condition = func(*gorm.DB) *gorm.DB

// From hold subQuery
type From struct {
	Alias    string
	SubQuery *gorm.DB
}

// Table return a table produced by SubQuery.
func Table(fromSubs ...From) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(fromSubs) == 0 {
			return db
		}
		tablePlaceholder := make([]string, len(fromSubs))
		tableExprs := make([]any, len(fromSubs))
		for i, query := range fromSubs {
			tablePlaceholder[i] = "(?)"
			tableExprs[i] = query.SubQuery
			if query.Alias != "" {
				tablePlaceholder[i] += " AS " + db.Statement.Quote(query.Alias)
			}
		}
		return db.Table(strings.Join(tablePlaceholder, ", "), tableExprs...)
	}
}

// Select with field
func Select(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db.Clauses(clause.Select{})
		}
		query, args := buildSelectValue(db.Statement, columns...)
		return db.Select(query, args...)
	}
}

// Distinct with field
func Distinct(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Distinct()
		if len(columns) > 0 {
			query, args := buildSelectValue(db.Statement, columns...)
			db = db.Select(query, args...)
		}
		return db
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

// LockingUpdate specify the lock strength to UPDATE
func LockingUpdate() Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
}

// LockingShare specify the lock strength to SHARE
func LockingShare() Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "SHARE"})
	}
}
