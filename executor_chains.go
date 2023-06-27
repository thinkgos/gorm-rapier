package assist

import (
	"gorm.io/gorm/clause"
)

func (x *Executor[T]) Table(name string, args ...any) *Executor[T] {
	x.table = GormTable(name, args...)
	return x
}

func (x *Executor[T]) TableExpr(fromSubs ...From) *Executor[T] {
	x.table = TableExpr(fromSubs...)
	return x
}

func (x *Executor[T]) Clauses(conds ...clause.Expression) *Executor[T] {
	x.funcs.Clauses(conds...)
	return x
}

func (x *Executor[T]) Distinct(args ...any) *Executor[T] {
	x.funcs.Distinct(args...)
	return x
}

func (x *Executor[T]) Select(query any, args ...any) *Executor[T] {
	x.funcs.Select(query, args...)
	return x
}

func (x *Executor[T]) Omit(columns ...string) *Executor[T] {
	x.funcs.Omit(columns...)
	return x
}

func (x *Executor[T]) Where(query any, args ...any) *Executor[T] {
	x.funcs.Where(query, args...)
	return x
}

func (x *Executor[T]) Not(query any, args ...any) *Executor[T] {
	x.funcs.Not(query, args...)
	return x
}

func (x *Executor[T]) Or(query any, args ...any) *Executor[T] {
	x.funcs.Or(query, args...)
	return x
}

func (x *Executor[T]) Joins(query string, args ...any) *Executor[T] {
	x.funcs.Joins(query, args...)
	return x
}

func (x *Executor[T]) InnerJoins(query string, args ...any) *Executor[T] {
	x.funcs.InnerJoins(query, args...)
	return x
}

func (x *Executor[T]) Group(name string) *Executor[T] {
	x.funcs.Group(name)
	return x
}

func (x *Executor[T]) Having(query any, args ...any) *Executor[T] {
	x.funcs.Having(query, args...)
	return x
}

func (x *Executor[T]) Order(value any) *Executor[T] {
	x.funcs.Order(value)
	return x
}

func (x *Executor[T]) Limit(limit int) *Executor[T] {
	x.funcs.Limit(limit)
	return x
}

func (x *Executor[T]) Offset(offset int) *Executor[T] {
	x.funcs.Offset(offset)
	return x
}

func (x *Executor[T]) Scopes(cs ...Condition) *Executor[T] {
	if len(cs) > 0 {
		x.funcs.Scopes(cs...)
	}
	return x
}

func (x *Executor[T]) Preload(query string, args ...any) *Executor[T] {
	x.funcs.Preload(query, args...)
	return x
}

func (x *Executor[T]) Attrs(attrs ...any) *Executor[T] {
	x.funcs.Attrs(attrs...)
	return x
}

func (x *Executor[T]) Assign(attrs ...any) *Executor[T] {
	x.funcs.Assign(attrs...)
	return x
}

func (x *Executor[T]) Unscoped() *Executor[T] {
	x.funcs.Unscoped()
	return x
}

func (x *Executor[T]) DistinctExpr(columns ...Expr) *Executor[T] {
	x.funcs.DistinctExpr(columns...)
	return x
}

func (x *Executor[T]) SelectExpr(columns ...Expr) *Executor[T] {
	x.funcs.SelectExpr(columns...)
	return x
}

func (x *Executor[T]) OrderExpr(columns ...Expr) *Executor[T] {
	x.funcs.OrderExpr(columns...)
	return x
}

func (x *Executor[T]) GroupExpr(columns ...Expr) *Executor[T] {
	x.funcs.GroupExpr(columns...)
	return x
}

func (x *Executor[T]) CrossJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs.CrossJoinsExpr(tableName, conds...)
	return x
}

func (x *Executor[T]) CrossJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs.CrossJoinsXExpr(tableName, alias, conds...)
	return x
}

func (x *Executor[T]) InnerJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs.InnerJoinsExpr(tableName, conds...)
	return x
}

func (x *Executor[T]) InnerJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs.InnerJoinsXExpr(tableName, alias, conds...)
	return x
}

func (x *Executor[T]) LeftJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs.LeftJoinsExpr(tableName, conds...)
	return x
}

func (x *Executor[T]) LeftJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs.LeftJoinsXExpr(tableName, alias, conds...)
	return x
}

func (x *Executor[T]) RightJoinsExpr(tableName string, conds ...Expr) *Executor[T] {
	x.funcs.RightJoinsExpr(tableName, conds...)
	return x
}

func (x *Executor[T]) RightJoinsXExpr(tableName, alias string, conds ...Expr) *Executor[T] {
	x.funcs.RightJoinsXExpr(tableName, alias, conds...)
	return x
}

func (x *Executor[T]) LockingUpdate() *Executor[T] {
	x.funcs.LockingUpdate()
	return x
}

func (x *Executor[T]) LockingShare() *Executor[T] {
	x.funcs.LockingShare()
	return x
}

func (x *Executor[T]) Pagination(page, perPage int64, maxPerPages ...int64) *Executor[T] {
	x.funcs.Pagination(page, perPage, maxPerPages...)
	return x
}
