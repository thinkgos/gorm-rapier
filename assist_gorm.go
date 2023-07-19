package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GormClauses(conds ...clause.Expression) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(conds...)
	}
}

func GormTable(name string, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(name, args...)
	}
}

func GormDistinct(args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Distinct(args...)
	}
}

func GormSelect(query any, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	}
}

func GormOmit(columns ...string) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	}
}

func GormWhere(query any, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func GormNot(query any, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Not(query, args...)
	}
}

func GormOr(query any, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(query, args...)
	}
}

func GormJoins(query string, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	}
}

func GormInnerJoins(query string, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.InnerJoins(query, args...)
	}
}

func GormGroup(name string) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Group(name)
	}
}

func GormHaving(query any, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	}
}

func GormOrder(value any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(value)
	}
}

func GormLimit(limit int) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func GormOffset(offset int) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

func GormPreload(query string, args ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(query, args...)
	}
}

func GormUnscoped() Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}
