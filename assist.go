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
		tableExprs := make([]interface{}, len(fromSubs))
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

// Where with field
func Where(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Where(columns[0], intoSlice(columns[1:])...)
	}
}

// Having with field
func Having(columns ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) == 0 {
			return db
		}
		return db.Having(columns[0], intoSlice(columns[1:])...)
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
func NewConditions() Conditions {
	return make(Conditions, 0, 16)
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
func (c Conditions) Pagination(page, perPage int64) Conditions {
	return append(c, Pagination(page, perPage))
}

// CrossJoin cross joins condition
func (c Conditions) CrossJoin(tableName string, conds ...Expr) Conditions {
	return append(c, CrossJoin(tableName, conds...))
}

// CrossJoinX cross joins condition
func (c Conditions) CrossJoinX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, CrossJoinX(tableName, alias, conds...))
}

// Join same as InnerJoin.
func (c Conditions) Join(tableName string, conds ...Expr) Conditions {
	return append(c, Join(tableName, conds...))
}

// JoinX same as InnerJoinX.
func (c Conditions) JoinX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, JoinX(tableName, alias, conds...))
}

// InnerJoin inner joins condition
func (c Conditions) InnerJoin(tableName string, conds ...Expr) Conditions {
	return append(c, InnerJoin(tableName, conds...))
}

// InnerJoinX inner joins condition
func (c Conditions) InnerJoinX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, InnerJoinX(tableName, alias, conds...))
}

// LeftJoin left join condition
func (c Conditions) LeftJoin(tableName string, conds ...Expr) Conditions {
	return append(c, LeftJoin(tableName, conds...))
}

// LeftJoinX left join condition
func (c Conditions) LeftJoinX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, LeftJoinX(tableName, alias, conds...))
}

// RightJoin right join condition
func (c Conditions) RightJoin(tableName string, conds ...Expr) Conditions {
	return append(c, RightJoin(tableName, conds...))
}

// RightJoinX right join condition
func (c Conditions) RightJoinX(tableName, alias string, conds ...Expr) Conditions {
	return append(c, RightJoinX(tableName, alias, conds...))
}
