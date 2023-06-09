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
	return field.ifNull(value)
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
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?).
func (field Float[T]) NotBetween(left T, right T) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field Float[T]) In(values ...T) Expr {
	return expr{
		col:       field.col,
		e:         clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)},
		buildOpts: field.buildOpts,
	}
}

// InAny use expr IN (?)
// value must be a array/slice
func (field Float[T]) InAny(value any) Expr {
	return expr{
		col:       field.col,
		e:         intoInExpr(field.RawExpr(), value),
		buildOpts: field.buildOpts,
	}
}

// NotIn use expr NOT IN (?)
func (field Float[T]) NotIn(values ...T) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)}),
		buildOpts: field.buildOpts,
	}
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field Float[T]) NotInAny(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(intoInExpr(field.RawExpr(), value)),
		buildOpts: field.buildOpts,
	}
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
	return field.findInSet(targetList)
}

// Sum use SUM(expr)
func (field Float[T]) Sum() Float[T] {
	return Float[T]{field.sum()}
}

// Add use expr+?
func (field Float[T]) Add(value T) Float[T] {
	return Float[T]{field.add(value)}
}

// Sub use expr-?
func (field Float[T]) Sub(value T) Float[T] {
	return Float[T]{field.sub(value)}
}

// Mul use expr*?
func (field Float[T]) Mul(value T) Float[T] {
	return Float[T]{field.mul(value)}
}

// Div use expr/?
func (field Float[T]) Div(value T) Float[T] {
	return Float[T]{field.div(value)}
}

// FloorDiv use expr DIV ?
func (field Float[T]) FloorDiv(value T) Int {
	return Int{field.floorDiv(value)}
}

// Floor se FLOOR(expr)
func (field Float[T]) Floor() Int {
	return Int{field.floor()}
}

// Round use ROUND(expr, ?)
func (field Float[T]) Round(decimals int) Float[T] {
	return Float[T]{field.round(decimals)}
}

// IntoColumns columns array with sub method
func (field Float[T]) IntoColumns() Columns {
	return NewColumns(field)
}
