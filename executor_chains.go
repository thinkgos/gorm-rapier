package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (x *Executor[T]) Clauses(conds ...clause.Expression) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Clauses(conds...)
	})
	return x
}

func (x *Executor[T]) Table(name string, args ...any) *Executor[T] {
	x.table = func(db *gorm.DB) *gorm.DB {
		return db.Table(name, args...)
	}
	return x
}

func (x *Executor[T]) Distinct(args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Distinct(args...)
	})
	return x
}

func (x *Executor[T]) Select(query any, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	})
	return x
}

func (x *Executor[T]) Omit(columns ...string) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	})
	return x
}

func (x *Executor[T]) Where(query any, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return x
}

func (x *Executor[T]) Not(query any, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Not(query, args...)
	})
	return x
}

func (x *Executor[T]) Or(query any, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Or(query, args...)
	})
	return x
}

func (x *Executor[T]) Joins(query string, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	})
	return x
}

func (x *Executor[T]) InnerJoins(query string, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.InnerJoins(query, args...)
	})
	return x
}

func (x *Executor[T]) Group(name string) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Group(name)
	})
	return x
}

func (x *Executor[T]) Having(query any, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	})
	return x
}

func (x *Executor[T]) Order(value any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Order(value)
	})
	return x
}

func (x *Executor[T]) Limit(limit int) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	})
	return x
}

func (x *Executor[T]) Offset(offset int) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	})
	return x
}

func (x *Executor[T]) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *Executor[T] {
	x.funcs = append(x.funcs, funcs...)
	return x
}

func (x *Executor[T]) Preload(query string, args ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Preload(query, args...)
	})
	return x
}

func (x *Executor[T]) Attrs(attrs ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Attrs(attrs...)
	})
	return x
}

func (x *Executor[T]) Assign(attrs ...any) *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Assign(attrs...)
	})
	return x
}

func (x *Executor[T]) Unscoped() *Executor[T] {
	x.funcs = append(x.funcs, func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	})
	return x
}

func (x *Executor[T]) TableExpr(fromSubs ...From) *Executor[T] {
	x.funcs = append(x.funcs, Table(fromSubs...))
	return x
}

func (x *Executor[T]) DistinctExpr(columns ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, Distinct(columns...))
	return x
}

func (x *Executor[T]) SelectExpr(columns ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, Select(columns...))
	return x
}

func (x *Executor[T]) OrderExpr(columns ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, Order(columns...))
	return x
}

func (x *Executor[T]) GroupExpr(columns ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, Group(columns...))
	return x
}

func (x *Executor[T]) CrossJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, CrossJoins(tableName, conds...))
	return x
}

func (x *Executor[T]) CrossJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, CrossJoinsX(tableName, alias, conds...))
	return x
}

func (x *Executor[T]) InnerJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, InnerJoins(tableName, conds...))
	return x
}

func (x *Executor[T]) InnerJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, InnerJoinsX(tableName, alias, conds...))
	return x
}

func (x *Executor[T]) LeftJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, LeftJoins(tableName, conds...))
	return x
}

func (x *Executor[T]) LeftJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, LeftJoinsX(tableName, alias, conds...))
	return x
}

func (x *Executor[T]) RightJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, RightJoins(tableName, conds...))
	return x
}

func (x *Executor[T]) RightJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs = append(x.funcs, RightJoinsX(tableName, alias, conds...))
	return x
}

func (x *Executor[T]) LockingUpdate() *Executor[T] {
	x.funcs = append(x.funcs, LockingUpdate())
	return x
}

func (x *Executor[T]) LockingShare() *Executor[T] {
	x.funcs = append(x.funcs, LockingShare())
	return x
}

func (x *Executor[T]) Pagination(page, perPage int64, maxPerPages ...int64) *Executor[T] {
	x.funcs = append(x.funcs, Pagination(page, perPage, maxPerPages...))
	return x
}
