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
	return x.getInstance(x.db.Clauses(conds...))
}

func (x *Executor[T]) Distinct(args ...any) *Executor[T] {
	return x.getInstance(x.db.Distinct(args...))
}

func (x *Executor[T]) Select(query any, args ...any) *Executor[T] {
	return x.getInstance(x.db.Select(query, args...))
}

func (x *Executor[T]) Omit(columns ...string) *Executor[T] {
	return x.getInstance(x.db.Omit(columns...))
}

func (x *Executor[T]) Where(query any, args ...any) *Executor[T] {
	return x.getInstance(x.db.Where(query, args...))
}

func (x *Executor[T]) Not(query any, args ...any) *Executor[T] {
	return x.getInstance(x.db.Not(query, args...))
}

func (x *Executor[T]) Or(query any, args ...any) *Executor[T] {
	return x.getInstance(x.db.Or(query, args...))
}

func (x *Executor[T]) Joins(query string, args ...any) *Executor[T] {
	return x.getInstance(x.db.Joins(query, args...))
}

func (x *Executor[T]) InnerJoins(query string, args ...any) *Executor[T] {
	return x.getInstance(x.db.InnerJoins(query, args...))
}

func (x *Executor[T]) Group(name string) *Executor[T] {
	return x.getInstance(x.db.Group(name))
}

func (x *Executor[T]) Having(query any, args ...any) *Executor[T] {
	return x.getInstance(x.db.Having(query, args...))
}

func (x *Executor[T]) Order(value any) *Executor[T] {
	return x.getInstance(x.db.Order(value))
}

func (x *Executor[T]) Limit(limit int) *Executor[T] {
	return x.getInstance(x.db.Limit(limit))
}

func (x *Executor[T]) Offset(offset int) *Executor[T] {
	return x.getInstance(x.db.Offset(offset))
}

func (x *Executor[T]) Scopes(cs ...Condition) *Executor[T] {
	x.scopes = append(x.scopes, cs...)
	return x
}

func (x *Executor[T]) Preload(query string, args ...any) *Executor[T] {
	return x.getInstance(x.db.Preload(query, args...))
}

func (x *Executor[T]) Unscoped() *Executor[T] {
	return x.getInstance(x.db.Unscoped())
}

func (x *Executor[T]) DistinctExpr(columns ...Expr) *Executor[T] {
	return x.execute(DistinctExpr(columns...))
}

func (x *Executor[T]) SelectExpr(columns ...Expr) *Executor[T] {
	return x.execute(SelectExpr(columns...))
}

func (x *Executor[T]) OmitExpr(columns ...Expr) *Executor[T] {
	return x.execute(OmitExpr(columns...))
}

func (x *Executor[T]) OrderExpr(columns ...Expr) *Executor[T] {
	return x.execute(OrderExpr(columns...))
}

func (x *Executor[T]) GroupExpr(columns ...Expr) *Executor[T] {
	return x.execute(GroupExpr(columns...))
}

func (x *Executor[T]) CrossJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	return x.execute(CrossJoinsExpr(table, conds...))
}

func (x *Executor[T]) CrossJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	return x.execute(CrossJoinsXExpr(table, alias, conds...))
}

func (x *Executor[T]) InnerJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	return x.execute(InnerJoinsExpr(table, conds...))
}

func (x *Executor[T]) InnerJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	return x.execute(InnerJoinsXExpr(table, alias, conds...))
}

func (x *Executor[T]) LeftJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	return x.execute(LeftJoinsExpr(table, conds...))
}

func (x *Executor[T]) LeftJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	return x.execute(LeftJoinsXExpr(table, alias, conds...))
}

func (x *Executor[T]) RightJoinsExpr(table schema.Tabler, conds ...Expr) *Executor[T] {
	return x.execute(RightJoinsExpr(table, conds...))
}

func (x *Executor[T]) RightJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Executor[T] {
	return x.execute(RightJoinsXExpr(table, alias, conds...))
}

func (x *Executor[T]) LockingUpdate() *Executor[T] {
	return x.execute(LockingUpdate())
}

func (x *Executor[T]) LockingShare() *Executor[T] {
	return x.execute(LockingShare())
}

func (x *Executor[T]) Pagination(page, perPage int64, maxPerPages ...int64) *Executor[T] {
	return x.execute(Pagination(page, perPage, maxPerPages...))
}

func (x *Executor[T]) Returning(columns ...string) *Executor[T] {
	clauseColumn := make([]clause.Column, 0, len(columns))
	for _, column := range columns {
		clauseColumn = append(clauseColumn, clause.Column{Name: column})
	}
	return x.getInstance(x.db.Clauses(clause.Returning{Columns: clauseColumn}))
}
