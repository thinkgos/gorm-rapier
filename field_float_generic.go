package assist

import (
	"golang.org/x/exp/constraints"
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

// IntoField convert into Field. then use Field abilities.
func (field Float[T]) IntoField() Field {
	return Field(field)
}

// IfNull use IFNULL(expr,?)
func (field Float[T]) IfNull(value T) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field Float[T]) Eq(value T) Expr {
	return field.innerEq(value)
}

// Neq not equal to, use expr <> ?
func (field Float[T]) Neq(value T) Expr {
	return field.innerNeq(value)
}

// Gt greater than, use expr > ?
func (field Float[T]) Gt(value T) Expr {
	return field.innerGt(value)
}

// Gte greater or equal to, use expr >= ?
func (field Float[T]) Gte(value T) Expr {
	return field.innerGte(value)
}

// Lt less than, use expr < ?
func (field Float[T]) Lt(value T) Expr {
	return field.innerLt(value)
}

// Lte less or equal to, use expr <= ?
func (field Float[T]) Lte(value T) Expr {
	return field.innerLte(value)
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
	return field.innerLike(value)
}

// NotLike use expr NOT LIKE ?
func (field Float[T]) NotLike(value T) Expr {
	return field.innerNotLike(value)
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
	return Float[T]{e.innerAddCol(col)}
}

// SubCol use expr1 - expr2
func (e Float[T]) SubCol(col Expr) Float[T] {
	return Float[T]{e.innerSubCol(col)}
}

// MulCol use (expr1) * (expr2)
func (e Float[T]) MulCol(col Expr) Float[T] {
	return Float[T]{e.innerMulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (e Float[T]) DivCol(col Expr) Float[T] {
	return Float[T]{e.innerDivCol(col)}
}

// Value set value
func (field Float[T]) Value(value T) SetExpr {
	return field.value(value)
}

// Value set value use pointer
func (field Float[T]) ValuePointer(value *T) SetExpr {
	return field.value(value)
}

// ValueZero set value zero
func (field Float[T]) ValueZero() SetExpr {
	var zero T

	return field.value(zero)
}
