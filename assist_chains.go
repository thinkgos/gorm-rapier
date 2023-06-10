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

// Select with field
func (c *Conditions) Select(columns ...Expr) *Conditions {
	return c.Append(Select(columns...))
}

// Distinct with field
func (c *Conditions) Distinct(columns ...Expr) *Conditions {
	return c.Append(Distinct(columns...))
}

// Order with field
func (c *Conditions) Order(columns ...Expr) *Conditions {
	return c.Append(Order(columns...))
}

// Group with field
func (c *Conditions) Group(columns ...Expr) *Conditions {
	return c.Append(Group(columns...))
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

// CrossJoins cross joins condition
func (c *Conditions) CrossJoins(tableName string, conds ...Expr) *Conditions {
	return c.Append(CrossJoins(tableName, conds...))
}

// CrossJoinsX cross joins condition
func (c *Conditions) CrossJoinsX(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(CrossJoinsX(tableName, alias, conds...))
}

// InnerJoins inner joins condition
func (c *Conditions) InnerJoins(tableName string, conds ...Expr) *Conditions {
	return c.Append(InnerJoins(tableName, conds...))
}

// InnerJoinsX inner joins condition
func (c *Conditions) InnerJoinsX(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(InnerJoinsX(tableName, alias, conds...))
}

// LeftJoins left join condition
func (c *Conditions) LeftJoins(tableName string, conds ...Expr) *Conditions {
	return c.Append(LeftJoins(tableName, conds...))
}

// LeftJoinsX left join condition
func (c *Conditions) LeftJoinsX(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(LeftJoinsX(tableName, alias, conds...))
}

// RightJoins right join condition
func (c *Conditions) RightJoins(tableName string, conds ...Expr) *Conditions {
	return c.Append(RightJoins(tableName, conds...))
}

// RightJoinsX right join condition
func (c *Conditions) RightJoinsX(tableName, alias string, conds ...Expr) *Conditions {
	return c.Append(RightJoinsX(tableName, alias, conds...))
}

// Append more condition
func (c *Conditions) Append(cs ...Condition) *Conditions {
	c.funcs = append(c.funcs, cs...)
	return c
}
