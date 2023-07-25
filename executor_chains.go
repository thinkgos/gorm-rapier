package assist

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
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
	x.conditions.Clauses(conds...)
	return x
}

func (x *Executor[T]) Distinct(args ...any) *Executor[T] {
	x.conditions.Distinct(args...)
	return x
}

func (x *Executor[T]) Select(query any, args ...any) *Executor[T] {
	x.conditions.Select(query, args...)
	return x
}

func (x *Executor[T]) Omit(columns ...string) *Executor[T] {
	x.conditions.Omit(columns...)
	return x
}

func (x *Executor[T]) Where(query any, args ...any) *Executor[T] {
	x.conditions.Where(query, args...)
	return x
}

func (x *Executor[T]) Not(query any, args ...any) *Executor[T] {
	x.conditions.Not(query, args...)
	return x
}

func (x *Executor[T]) Or(query any, args ...any) *Executor[T] {
	x.conditions.Or(query, args...)
	return x
}

func (x *Executor[T]) Joins(query string, args ...any) *Executor[T] {
	x.conditions.Joins(query, args...)
	return x
}

func (x *Executor[T]) InnerJoins(query string, args ...any) *Executor[T] {
	x.conditions.InnerJoins(query, args...)
	return x
}

func (x *Executor[T]) Group(name string) *Executor[T] {
	x.conditions.Group(name)
	return x
}

func (x *Executor[T]) Having(query any, args ...any) *Executor[T] {
	x.conditions.Having(query, args...)
	return x
}

func (x *Executor[T]) Order(value any) *Executor[T] {
	x.conditions.Order(value)
	return x
}

func (x *Executor[T]) Limit(limit int) *Executor[T] {
	x.conditions.Limit(limit)
	return x
}

func (x *Executor[T]) Offset(offset int) *Executor[T] {
	x.conditions.Offset(offset)
	return x
}

func (x *Executor[T]) Scopes(cs ...Condition) *Executor[T] {
	if len(cs) > 0 {
		x.conditions.Scopes(cs...)
	}
	return x
}

func (x *Executor[T]) Preload(query string, args ...any) *Executor[T] {
	x.conditions.Preload(query, args...)
	return x
}

func (x *Executor[T]) Unscoped() *Executor[T] {
	x.conditions.Unscoped()
	return x
}

func (x *Executor[T]) DistinctExpr(columns ...Expr) *Executor[T] {
	x.conditions.DistinctExpr(columns...)
	return x
}

func (x *Executor[T]) SelectExpr(columns ...Expr) *Executor[T] {
	x.conditions.SelectExpr(columns...)
	return x
}

func (x *Executor[T]) OmitExpr(columns ...Expr) *Executor[T] {
	x.conditions.OmitExpr(columns...)
	return x
}

func (x *Executor[T]) OrderExpr(columns ...Expr) *Executor[T] {
	x.conditions.OrderExpr(columns...)
	return x
}

func (x *Executor[T]) GroupExpr(columns ...Expr) *Executor[T] {
	x.conditions.GroupExpr(columns...)
	return x
}

func (x *Executor[T]) CrossJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	x.conditions.CrossJoinsExpr(table, conds...)
	return x
}

func (x *Executor[T]) CrossJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	x.conditions.CrossJoinsXExpr(table, alias, conds...)
	return x
}

func (x *Executor[T]) InnerJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	x.conditions.InnerJoinsExpr(table, conds...)
	return x
}

func (x *Executor[T]) InnerJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	x.conditions.InnerJoinsXExpr(table, alias, conds...)
	return x
}

func (x *Executor[T]) LeftJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	x.conditions.LeftJoinsExpr(table, conds...)
	return x
}

func (x *Executor[T]) LeftJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	x.conditions.LeftJoinsXExpr(table, alias, conds...)
	return x
}

func (x *Executor[T]) RightJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	x.conditions.RightJoinsExpr(table, conds...)
	return x
}

func (x *Executor[T]) RightJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	x.conditions.RightJoinsXExpr(table, alias, conds...)
	return x
}

func (x *Executor[T]) LockingUpdate() *Executor[T] {
	x.conditions.LockingUpdate()
	return x
}

func (x *Executor[T]) LockingShare() *Executor[T] {
	x.conditions.LockingShare()
	return x
}

func (x *Executor[T]) Pagination(page, perPage int64, maxPerPages ...int64) *Executor[T] {
	x.conditions.Pagination(page, perPage, maxPerPages...)
	return x
}
