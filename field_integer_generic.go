package rapier

import (
	"strings"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm/clause"
)

// Integer type field
type Integer[T constraints.Integer] Field

// NewInt new Integer
func NewInteger[T constraints.Integer](table, column string, opts ...Option) Integer[T] {
	return Integer[T]{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IntoField convert into Field. then use Field abilities.
func (field Integer[T]) IntoField() Field {
	return Field(field)
}

// IfNull use IFNULL(expr,?)
func (field Integer[T]) IfNull(value T) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field Integer[T]) Eq(value T) Expr {
	return field.innerEq(value)
}

// Neq not equal to, use expr <> ?
func (field Integer[T]) Neq(value T) Expr {
	return field.innerNeq(value)
}

// Gt greater than, use expr > ?
func (field Integer[T]) Gt(value T) Expr {
	return field.innerGt(value)
}

// Gte greater or equal to, use expr >= ?
func (field Integer[T]) Gte(value T) Expr {
	return field.innerGte(value)
}

// Lt less than, use expr < ?
func (field Integer[T]) Lt(value T) Expr {
	return field.innerLt(value)
}

// Lte less or equal to, use expr <= ?
func (field Integer[T]) Lte(value T) Expr {
	return field.innerLte(value)
}

// Between use expr BETWEEN ? AND ?
func (field Integer[T]) Between(left T, right T) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Integer[T]) NotBetween(left T, right T) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field Integer[T]) In(values ...T) Expr {
	return field.innerIn(intoAnySlice(values))
}

// InAny use expr IN (?)
// value must be a array/slice
func (field Integer[T]) InAny(value any) Expr {
	return field.innerInAny(value)
}

// NotIn use expr NOT IN (?)
func (field Integer[T]) NotIn(values ...T) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field Integer[T]) NotInAny(value any) Expr {
	return field.innerNotInAny(value)
}

// Like use expr LIKE ?
func (field Integer[T]) Like(value T) Expr {
	return field.innerLike(value)
}

// NotLike use expr NOT LIKE ?
func (field Integer[T]) NotLike(value T) Expr {
	return field.innerNotLike(value)
}

// FindInSet equal to FIND_IN_SET(expr, ?)
func (field Integer[T]) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// Sum use SUM(expr)
func (field Integer[T]) Sum() Integer[T] {
	return Integer[T]{field.innerSum()}
}

// Add use expr+?
func (field Integer[T]) Add(value T) Integer[T] {
	return Integer[T]{field.innerAdd(value)}
}

// Add use expr-?
func (field Integer[T]) Sub(value T) Integer[T] {
	return Integer[T]{field.innerSub(value)}
}

// Mul use expr*?
func (field Integer[T]) Mul(value T) Integer[T] {
	return Integer[T]{field.innerMul(value)}
}

// Div use expr/?
func (field Integer[T]) Div(value T) Integer[T] {
	return Integer[T]{field.innerDiv(value)}
}

// Mod use expr%?
func (field Integer[T]) Mod(value T) Integer[T] {
	return Integer[T]{field.innerMod(value)}
}

// FloorDiv use expr DIV ?
func (field Integer[T]) FloorDiv(value T) Integer[T] {
	return Integer[T]{field.innerFloorDiv(value)}
}

// Round use ROUND(expr, ?)
func (field Integer[T]) Round(value int) Integer[T] {
	return Integer[T]{field.innerRound(value)}
}

// RightShift use expr>>?
func (field Integer[T]) RightShift(value T) Integer[T] {
	return Integer[T]{field.innerRightShift(value)}
}

// LeftShift use expr<<?
func (field Integer[T]) LeftShift(value T) Integer[T] {
	return Integer[T]{field.innerLeftShift(value)}
}

// BitXor use expr expr^?
func (field Integer[T]) BitXor(value T) Integer[T] {
	return Integer[T]{field.innerBitXor(value)}
}

// BitAnd use expr expr&?
func (field Integer[T]) BitAnd(value T) Integer[T] {
	return Integer[T]{field.innerBitAnd(value)}
}

// BitOr use expr expr|?
func (field Integer[T]) BitOr(value T) Integer[T] {
	return Integer[T]{field.innerBitOr(value)}
}

// BitFlip use expr ~expr
func (field Integer[T]) BitFlip() Integer[T] {
	return Integer[T]{field.innerBitFlip()}
}

// FromUnixTime use FromUnixTime(unix_timestamp[, format])
func (field Integer[T]) FromUnixTime(format ...string) String {
	var e clause.Expression

	if len(format) > 0 && strings.TrimSpace(format[0]) != "" {
		e = &clause.Expr{SQL: "FROM_UNIXTIME(?, ?)", Vars: []any{field.RawExpr(), format[0]}}
	} else {
		e = &clause.Expr{SQL: "FROM_UNIXTIME(?)", Vars: []any{field.RawExpr()}}
	}
	return String{
		expr{
			col:       field.col,
			e:         e,
			buildOpts: field.buildOpts,
		},
	}
}

// FromDays use FROM_DAYS(value)
func (field Integer[T]) FromDays() Time {
	return Time{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "FROM_DAYS(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// AddCol use expr1 + expr2
func (e Integer[T]) AddCol(col Expr) Integer[T] {
	return Integer[T]{e.innerAddCol(col)}
}

// SubCol use expr1 - expr2
func (e Integer[T]) SubCol(col Expr) Integer[T] {
	return Integer[T]{e.innerSubCol(col)}
}

// MulCol use (expr1) * (expr2)
func (e Integer[T]) MulCol(col Expr) Integer[T] {
	return Integer[T]{e.innerMulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (e Integer[T]) DivCol(col Expr) Integer[T] {
	return Integer[T]{e.innerDivCol(col)}
}

// Value set value
func (field Integer[T]) Value(value T) SetExpr {
	return field.value(value)
}

// Value set value use pointer
func (field Integer[T]) ValuePointer(value *T) SetExpr {
	return field.value(value)
}

// ValueZero set value zero
func (field Integer[T]) ValueZero() SetExpr {
	var zero T

	return field.value(zero)
}
