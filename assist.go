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

// Where with field
func Where(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Where(columns[0], intoAnySlice(columns[1:])...)
	}
}

// Having with field
func Having(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Having(columns[0], intoAnySlice(columns[1:])...)
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

// Conditions Condition slice
type Conditions []Condition

// NewConditions new condition instance.
func NewConditions(cs ...Condition) Conditions {
	c := make(Conditions, 0, 16)
	return append(c, cs...)
}

// Build into []Condition
func (c Conditions) Build() []Condition { return c }

// Table with field
func (c Conditions) Table(fromSubs ...From) Conditions {
	return append(c, Table(fromSubs...))
}

// Select with field
func (c Conditions) Select(columns ...Expr) Conditions {
	return append(c, Select(columns...))
}

// Distinct with field
func (c Conditions) Distinct(columns ...Expr) Conditions {
	return append(c, Distinct(columns...))
}

// Order with field
func (c Conditions) Order(columns ...Expr) Conditions {
	return append(c, Order(columns...))
}

// Group with field
func (c Conditions) Group(columns ...Expr) Conditions {
	return append(c, Group(columns...))
}

// Group with field
func (c Conditions) Where(columns ...Expr) Conditions {
	return append(c, Where(columns...))
}

// Having with field
func (c Conditions) Having(columns ...Expr) Conditions {
	return append(c, Having(columns...))
}

// LockingUpdate specify the lock strength to UPDATE
func (c Conditions) LockingUpdate() Conditions {
	return append(c, LockingUpdate())
}

// LockingShare specify the lock strength to SHARE
func (c Conditions) LockingShare() Conditions {
	return append(c, LockingShare())
}

// Pagination 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func (c Conditions) Pagination(page, perPage int64, maxPages ...int64) Conditions {
	return append(c, Pagination(page, perPage, maxPages...))
}

// CrossJoins cross joins condition
func (c Conditions) CrossJoins(tableName string, conds ...Expr) Conditions {
	return append(c, CrossJoins(tableName, conds...))
}

// CrossJoinsX cross joins condition
func (c Conditions) CrossJoinsX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, CrossJoinsX(tableName, alias, conds...))
}

// InnerJoins inner joins condition
func (c Conditions) InnerJoins(tableName string, conds ...Expr) Conditions {
	return append(c, InnerJoins(tableName, conds...))
}

// InnerJoinsX inner joins condition
func (c Conditions) InnerJoinsX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, InnerJoinsX(tableName, alias, conds...))
}

// LeftJoins left join condition
func (c Conditions) LeftJoins(tableName string, conds ...Expr) Conditions {
	return append(c, LeftJoins(tableName, conds...))
}

// LeftJoinsX left join condition
func (c Conditions) LeftJoinsX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, LeftJoinsX(tableName, alias, conds...))
}

// RightJoins right join condition
func (c Conditions) RightJoins(tableName string, conds ...Expr) Conditions {
	return append(c, RightJoins(tableName, conds...))
}

// RightJoinsX right join condition
func (c Conditions) RightJoinsX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, RightJoinsX(tableName, alias, conds...))
}

// Append more condition
func (c Conditions) Append(cs ...Condition) Conditions {
	return append(c, cs...)
}
