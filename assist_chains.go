package assist

import "gorm.io/gorm/clause"

// Conditions hold Condition slice
type Conditions struct {
	funcs []Condition
}

// NewConditions new condition instance.
func NewConditions(cs ...Condition) *Conditions {
	funcs := make([]Condition, 0, 16)
	return &Conditions{
		funcs: append(funcs, cs...),
	}
}

// Table with field
func (c *Conditions) Build() []Condition {
	return c.funcs
}

func (x *Conditions) Clauses(conds ...clause.Expression) *Conditions {
	if len(conds) > 0 {
		x.Scopes(GormClauses(conds...))
	}
	return x
}

// Distinct with field
func (c *Conditions) Distinct(args ...any) *Conditions {
	return c.Scopes(GormDistinct(args...))
}

// Select with field
func (c *Conditions) Select(query any, args ...any) *Conditions {
	return c.Scopes(GormSelect(query, args...))
}

func (x *Conditions) Omit(columns ...string) *Conditions {
	x.Scopes(GormOmit(columns...))
	return x
}

func (x *Conditions) Where(query any, args ...any) *Conditions {
	x.Scopes(GormWhere(query, args...))
	return x
}

func (x *Conditions) Not(query any, args ...any) *Conditions {
	x.Scopes(GormNot(query, args...))
	return x
}

func (x *Conditions) Or(query any, args ...any) *Conditions {
	x.Scopes(GormOr(query, args...))
	return x
}

func (x *Conditions) Joins(query string, args ...any) *Conditions {
	x.Scopes(GormJoins(query, args...))
	return x
}

func (x *Conditions) InnerJoins(query string, args ...any) *Conditions {
	x.Scopes(GormInnerJoins(query, args...))
	return x
}

func (x *Conditions) Group(name string) *Conditions {
	x.Scopes(GormGroup(name))
	return x
}

func (x *Conditions) Having(query any, args ...any) *Conditions {
	x.Scopes(GormHaving(query, args...))
	return x
}

func (x *Conditions) Order(value any) *Conditions {
	x.Scopes(GormOrder(value))
	return x
}

func (x *Conditions) Limit(limit int) *Conditions {
	x.Scopes(GormLimit(limit))
	return x
}

func (x *Conditions) Offset(offset int) *Conditions {
	x.Scopes(GormOffset(offset))
	return x
}

// Scopes more condition
func (c *Conditions) Scopes(cs ...Condition) *Conditions {
	c.funcs = append(c.funcs, cs...)
	return c
}

func (x *Conditions) Preload(query string, args ...any) *Conditions {
	x.Scopes(GormPreload(query, args...))
	return x
}

func (x *Conditions) Attrs(attrs ...any) *Conditions {
	x.Scopes(GormAttrs(attrs...))
	return x
}

func (x *Conditions) Assign(attrs ...any) *Conditions {
	x.Scopes(GormAssign(attrs...))
	return x
}

func (x *Conditions) Unscoped() *Conditions {
	x.Scopes(GormUnscoped())
	return x
}

// DistinctExpr with field
func (c *Conditions) DistinctExpr(columns ...Expr) *Conditions {
	return c.Scopes(DistinctExpr(columns...))
}

// SelectExpr with field
func (c *Conditions) SelectExpr(columns ...Expr) *Conditions {
	return c.Scopes(SelectExpr(columns...))
}

// OrderExpr with field
func (c *Conditions) OrderExpr(columns ...Expr) *Conditions {
	return c.Scopes(OrderExpr(columns...))
}

// GroupExpr with field
func (c *Conditions) GroupExpr(columns ...Expr) *Conditions {
	return c.Scopes(GroupExpr(columns...))
}

// LockingUpdate specify the lock strength to UPDATE
func (c *Conditions) LockingUpdate() *Conditions {
	return c.Scopes(LockingUpdate())
}

// LockingShare specify the lock strength to SHARE
func (c *Conditions) LockingShare() *Conditions {
	return c.Scopes(LockingShare())
}

// CrossJoinsExpr cross joins condition
func (c *Conditions) CrossJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Scopes(CrossJoinsExpr(tableName, conds...))
}

// CrossJoinsXExpr cross joins condition
func (c *Conditions) CrossJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Scopes(CrossJoinsXExpr(tableName, alias, conds...))
}

// InnerJoinsExpr inner joins condition
func (c *Conditions) InnerJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Scopes(InnerJoinsExpr(tableName, conds...))
}

// InnerJoinsXExpr inner joins condition
func (c *Conditions) InnerJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Scopes(InnerJoinsXExpr(tableName, alias, conds...))
}

// LeftJoinsExpr left join condition
func (c *Conditions) LeftJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Scopes(LeftJoinsExpr(tableName, conds...))
}

// LeftJoinsXExpr left join condition
func (c *Conditions) LeftJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Scopes(LeftJoinsXExpr(tableName, alias, conds...))
}

// RightJoinsExpr right join condition
func (c *Conditions) RightJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Scopes(RightJoinsExpr(tableName, conds...))
}

// RightJoinsXExpr right join condition
func (c *Conditions) RightJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Scopes(RightJoinsXExpr(tableName, alias, conds...))
}

// Pagination 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func (c *Conditions) Pagination(page, perPage int64, maxPages ...int64) *Conditions {
	return c.Scopes(Pagination(page, perPage, maxPages...))
}
