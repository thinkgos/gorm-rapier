package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (x *Executor[T]) Clauses(conds ...clause.Expression) *Executor[T] {
	if len(conds) > 0 {
		x.funcs.Append(GormClauses(conds...))
	}
	return x
}

func (x *Executor[T]) Table(name string, args ...any) *Executor[T] {
	x.table = GormTable(name, args...)
	return x
}

func (x *Executor[T]) Distinct(args ...any) *Executor[T] {
	x.funcs.Append(GormDistinct(args...))
	return x
}

func (x *Executor[T]) Select(query any, args ...any) *Executor[T] {
	x.funcs.Append(GormSelect(query, args...))
	return x
}

func (x *Executor[T]) Omit(columns ...string) *Executor[T] {
	x.funcs.Append(GormOmit(columns...))
	return x
}

func (x *Executor[T]) Where(query any, args ...any) *Executor[T] {
	x.funcs.Append(GormWhere(query, args...))
	return x
}

func (x *Executor[T]) Not(query any, args ...any) *Executor[T] {
	x.funcs.Append(GormNot(query, args...))
	return x
}

func (x *Executor[T]) Or(query any, args ...any) *Executor[T] {
	x.funcs.Append(GormOr(query, args...))
	return x
}

func (x *Executor[T]) Joins(query string, args ...any) *Executor[T] {
	x.funcs.Append(GormJoins(query, args...))
	return x
}

func (x *Executor[T]) InnerJoins(query string, args ...any) *Executor[T] {
	x.funcs.Append(GormInnerJoins(query, args...))
	return x
}

func (x *Executor[T]) Group(name string) *Executor[T] {
	x.funcs.Append(GormGroup(name))
	return x
}

func (x *Executor[T]) Having(query any, args ...any) *Executor[T] {
	x.funcs.Append(GormHaving(query, args...))
	return x
}

func (x *Executor[T]) Order(value any) *Executor[T] {
	x.funcs.Append(GormOrder(value))
	return x
}

func (x *Executor[T]) Limit(limit int) *Executor[T] {
	x.funcs.Append(GormLimit(limit))
	return x
}

func (x *Executor[T]) Offset(offset int) *Executor[T] {
	x.funcs.Append(GormOffset(offset))
	return x
}

func (x *Executor[T]) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *Executor[T] {
	if len(funcs) > 0 {
		x.funcs.Append(funcs...)
	}
	return x
}

func (x *Executor[T]) Preload(query string, args ...any) *Executor[T] {
	x.funcs.Append(GormPreload(query, args...))
	return x
}

func (x *Executor[T]) Attrs(attrs ...any) *Executor[T] {
	x.funcs.Append(GormAttrs(attrs...))
	return x
}

func (x *Executor[T]) Assign(attrs ...any) *Executor[T] {
	x.funcs.Append(GormAssign(attrs...))
	return x
}

func (x *Executor[T]) Unscoped() *Executor[T] {
	x.funcs.Append(GormUnscoped())
	return x
}

func (x *Executor[T]) TableExpr(fromSubs ...From) *Executor[T] {
	x.funcs.Table(fromSubs...)
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
