package assist

import (
	"golang.org/x/exp/constraints"
	"gorm.io/gorm/clause"
)

// Float type field
type Float[T constraints.Float | ~string] Field

// NewFloat new float field.
func NewFloat[T constraints.Float | ~string](table, column string, opts ...Option) Float[T] {
	return Float[T]{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IfNull use IFNULL(expr,?)
func (field Float[T]) IfNull(value T) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field Float[T]) Eq(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Eq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Neq not equal to, use expr <> ?
func (field Float[T]) Neq(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Neq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gt greater than, use expr > ?
func (field Float[T]) Gt(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gte greater or equal to, use expr >= ?
func (field Float[T]) Gte(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lt less than, use expr < ?
func (field Float[T]) Lt(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lte less or equal to, use expr <= ?
func (field Float[T]) Lte(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Between use expr BETWEEN ? AND ?
func (field Float[T]) Between(left T, right T) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?).
func (field Float[T]) NotBetween(left T, right T) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field Float[T]) In(values ...T) Expr {
	return field.innerIn(intoAnySlice(values))
}

// InAny use expr IN (?)
// value must be a array/slice
func (field Float[T]) InAny(value any) Expr {
	return field.innerInAny(value)
}

// NotIn use expr NOT IN (?)
func (field Float[T]) NotIn(values ...T) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field Float[T]) NotInAny(value any) Expr {
	return field.innerNotInAny(value)
}

// Like use expr LIKE ?
func (field Float[T]) Like(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Like{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// NotLike use expr NOT LIKE ?
func (field Float[T]) NotLike(value T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(clause.Like{Column: field.RawExpr(), Value: value}),
		buildOpts: field.buildOpts,
	}
}

// FindInSet use FIND_IN_SET(expr, ?)
func (field Float[T]) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// Sum use SUM(expr)
func (field Float[T]) Sum() Float[T] {
	return Float[T]{field.innerSum()}
}

// Add use expr+?
func (field Float[T]) Add(value T) Float[T] {
	return Float[T]{field.innerAdd(value)}
}

// Sub use expr-?
func (field Float[T]) Sub(value T) Float[T] {
	return Float[T]{field.innerSub(value)}
}

// Mul use expr*?
func (field Float[T]) Mul(value T) Float[T] {
	return Float[T]{field.innerMul(value)}
}

// Div use expr/?
func (field Float[T]) Div(value T) Float[T] {
	return Float[T]{field.innerDiv(value)}
}

// FloorDiv use expr DIV ?
func (field Float[T]) FloorDiv(value T) Int {
	return Int{field.innerFloorDiv(value)}
}

// Floor se FLOOR(expr)
func (field Float[T]) Floor() Int {
	return Int{field.innerFloor()}
}

// Round use ROUND(expr, ?)
func (field Float[T]) Round(decimals int) Float[T] {
	return Float[T]{field.innerRound(decimals)}
}

// AddCol use expr1 + expr2
func (e Float[T]) AddCol(col Expr) Float[T] {
	return Float[T]{e.addCol(col)}
}

// SubCol use expr1 - expr2
func (e Float[T]) SubCol(col Expr) Float[T] {
	return Float[T]{e.subCol(col)}
}

// MulCol use (expr1) * (expr2)
func (e Float[T]) MulCol(col Expr) Float[T] {
	return Float[T]{e.mulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (e Float[T]) DivCol(col Expr) Float[T] {
	return Float[T]{e.divCol(col)}
}

// IntoColumns columns array with sub method
func (field Float[T]) IntoColumns() Columns {
	return NewColumns(field)
}
