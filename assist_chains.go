package assist

// Conditions Condition slice
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

// Table with field
func (c *Conditions) Table(fromSubs ...From) *Conditions {
	return c.Append(Table(fromSubs...))
}

// SelectExpr with field
func (c *Conditions) SelectExpr(columns ...Expr) *Conditions {
	return c.Append(SelectExpr(columns...))
}

// DistinctExpr with field
func (c *Conditions) DistinctExpr(columns ...Expr) *Conditions {
	return c.Append(DistinctExpr(columns...))
}

// OrderExpr with field
func (c *Conditions) OrderExpr(columns ...Expr) *Conditions {
	return c.Append(OrderExpr(columns...))
}

// GroupExpr with field
func (c *Conditions) GroupExpr(columns ...Expr) *Conditions {
	return c.Append(GroupExpr(columns...))
}

// LockingUpdate specify the lock strength to UPDATE
func (c *Conditions) LockingUpdate() *Conditions {
	return c.Append(LockingUpdate())
}

// LockingShare specify the lock strength to SHARE
func (c *Conditions) LockingShare() *Conditions {
	return c.Append(LockingShare())
}

// Pagination 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func (c *Conditions) Pagination(page, perPage int64, maxPages ...int64) *Conditions {
	return c.Append(Pagination(page, perPage, maxPages...))
}

// CrossJoinsExpr cross joins condition
func (c *Conditions) CrossJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Append(CrossJoinsExpr(tableName, conds...))
}

// CrossJoinsXExpr cross joins condition
func (c *Conditions) CrossJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(CrossJoinsXExpr(tableName, alias, conds...))
}

// InnerJoinsExpr inner joins condition
func (c *Conditions) InnerJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Append(InnerJoinsExpr(tableName, conds...))
}

// InnerJoinsXExpr inner joins condition
func (c *Conditions) InnerJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(InnerJoinsXExpr(tableName, alias, conds...))
}

// LeftJoinsExpr left join condition
func (c *Conditions) LeftJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Append(LeftJoinsExpr(tableName, conds...))
}

// LeftJoinsXExpr left join condition
func (c *Conditions) LeftJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(LeftJoinsXExpr(tableName, alias, conds...))
}

// RightJoinsExpr right join condition
func (c *Conditions) RightJoinsExpr(tableName string, conds ...Expr) *Conditions {
	return c.Append(RightJoinsExpr(tableName, conds...))
}

// RightJoinsXExpr right join condition
func (c *Conditions) RightJoinsXExpr(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(RightJoinsXExpr(tableName, alias, conds...))
}

// Append more condition
func (c *Conditions) Append(cs ...Condition) *Conditions {
	c.funcs = append(c.funcs, cs...)
	return c
}
