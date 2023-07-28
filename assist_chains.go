package assist

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

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

// Configure the conditions in the scope
func (c *Conditions) Configure(funcs ...func(*Conditions) *Conditions) *Conditions {
	for _, f := range funcs {
		c = f(c)
	}
	return c
}

// Table with field
func (c *Conditions) Build() []Condition {
	return c.funcs
}

func (c *Conditions) Clauses(conds ...clause.Expression) *Conditions {
	if len(conds) > 0 {
		c.Scopes(GormClauses(conds...))
	}
	return c
}

// Distinct with field
func (c *Conditions) Distinct(args ...any) *Conditions {
	return c.Scopes(GormDistinct(args...))
}

// Select with field
func (c *Conditions) Select(query any, args ...any) *Conditions {
	return c.Scopes(GormSelect(query, args...))
}

func (c *Conditions) Omit(columns ...string) *Conditions {
	c.Scopes(GormOmit(columns...))
	return c
}

func (c *Conditions) Where(query any, args ...any) *Conditions {
	c.Scopes(GormWhere(query, args...))
	return c
}

func (c *Conditions) Not(query any, args ...any) *Conditions {
	c.Scopes(GormNot(query, args...))
	return c
}

func (c *Conditions) Or(query any, args ...any) *Conditions {
	c.Scopes(GormOr(query, args...))
	return c
}

func (c *Conditions) Joins(query string, args ...any) *Conditions {
	c.Scopes(GormJoins(query, args...))
	return c
}

func (c *Conditions) InnerJoins(query string, args ...any) *Conditions {
	c.Scopes(GormInnerJoins(query, args...))
	return c
}

func (c *Conditions) Group(name string) *Conditions {
	c.Scopes(GormGroup(name))
	return c
}

func (c *Conditions) Having(query any, args ...any) *Conditions {
	c.Scopes(GormHaving(query, args...))
	return c
}

func (c *Conditions) Order(value any) *Conditions {
	c.Scopes(GormOrder(value))
	return c
}

func (c *Conditions) Limit(limit int) *Conditions {
	c.Scopes(GormLimit(limit))
	return c
}

func (c *Conditions) Offset(offset int) *Conditions {
	c.Scopes(GormOffset(offset))
	return c
}

// Scopes more condition
func (c *Conditions) Scopes(cs ...Condition) *Conditions {
	c.funcs = append(c.funcs, cs...)
	return c
}

func (c *Conditions) Preload(query string, args ...any) *Conditions {
	c.Scopes(GormPreload(query, args...))
	return c
}

func (c *Conditions) Unscoped() *Conditions {
	c.Scopes(GormUnscoped())
	return c
}

// DistinctExpr with field
func (c *Conditions) DistinctExpr(columns ...Expr) *Conditions {
	return c.Scopes(DistinctExpr(columns...))
}

// SelectExpr with field
func (c *Conditions) SelectExpr(columns ...Expr) *Conditions {
	return c.Scopes(SelectExpr(columns...))
}

// OmitExpr with field
func (c *Conditions) OmitExpr(columns ...Expr) *Conditions {
	return c.Scopes(OmitExpr(columns...))
}

// OrderExpr with field
func (c *Conditions) OrderExpr(columns ...Expr) *Conditions {
	return c.Scopes(OrderExpr(columns...))
}

// GroupExpr with field
func (c *Conditions) GroupExpr(columns ...Expr) *Conditions {
	return c.Scopes(GroupExpr(columns...))
}

// CrossJoinsExpr cross joins condition
func (c *Conditions) CrossJoinsExpr(table schema.Tabler, conds ...Expr) *Conditions {
	return c.Scopes(CrossJoinsExpr(table, conds...))
}

// CrossJoinsXExpr cross joins condition
func (c *Conditions) CrossJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Conditions {
	return c.Scopes(CrossJoinsXExpr(table, alias, conds...))
}

// InnerJoinsExpr inner joins condition
func (c *Conditions) InnerJoinsExpr(table schema.Tabler, conds ...Expr) *Conditions {
	return c.Scopes(InnerJoinsExpr(table, conds...))
}

// InnerJoinsXExpr inner joins condition
func (c *Conditions) InnerJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Conditions {
	return c.Scopes(InnerJoinsXExpr(table, alias, conds...))
}

// LeftJoinsExpr left join condition
func (c *Conditions) LeftJoinsExpr(table schema.Tabler, conds ...Expr) *Conditions {
	return c.Scopes(LeftJoinsExpr(table, conds...))
}

// LeftJoinsXExpr left join condition
func (c *Conditions) LeftJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Conditions {
	return c.Scopes(LeftJoinsXExpr(table, alias, conds...))
}

// RightJoinsExpr right join condition
func (c *Conditions) RightJoinsExpr(table schema.Tabler, conds ...Expr) *Conditions {
	return c.Scopes(RightJoinsExpr(table, conds...))
}

// RightJoinsXExpr right join condition
func (c *Conditions) RightJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) *Conditions {
	return c.Scopes(RightJoinsXExpr(table, alias, conds...))
}

// LockingUpdate specify the lock strength to UPDATE
func (c *Conditions) LockingUpdate() *Conditions {
	return c.Scopes(LockingUpdate())
}

// LockingShare specify the lock strength to SHARE
func (c *Conditions) LockingShare() *Conditions {
	return c.Scopes(LockingShare())
}

// Pagination 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func (c *Conditions) Pagination(page, perPage int64, maxPages ...int64) *Conditions {
	return c.Scopes(Pagination(page, perPage, maxPages...))
}
