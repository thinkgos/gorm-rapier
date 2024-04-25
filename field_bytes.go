package rapier

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// Bytes []byte type field
type Bytes Field

// NewBytes new bytes field.
func NewBytes(table, column string, opts ...Option) Bytes {
	return Bytes{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IntoField convert into Field. then use Field abilities.
func (field Bytes) IntoField() Field {
	return Field(field)
}

// IfNull use IFNULL(expr,?)
func (field Bytes) IfNull(value []byte) Expr {
	return field.innerIfNull(value)
}

// NullIf use NULLIF(expr,?)
func (field Bytes) NullIf(value []byte) Expr {
	return field.innerNullIf(value)
}

// Eq equal to, use expr = ?
func (field Bytes) Eq(value []byte) Expr {
	return field.innerEq(value)
}

// Neq not equal to, use expr <> ?
func (field Bytes) Neq(value []byte) Expr {
	return field.innerNeq(value)
}

// Gt greater than, use expr > ?
func (field Bytes) Gt(value []byte) Expr {
	return field.innerGt(value)
}

// Gte greater or equal to, use expr >= ?
func (field Bytes) Gte(value []byte) Expr {
	return field.innerGte(value)
}

// Lt less than, use expr < ?
func (field Bytes) Lt(value []byte) Expr {
	return field.innerLt(value)
}

// Lte less or equal to, use expr <= ?
func (field Bytes) Lte(value []byte) Expr {
	return field.innerLte(value)
}

// Between use expr BETWEEN ? AND ?
func (field Bytes) Between(left []byte, right []byte) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Bytes) NotBetween(left []byte, right []byte) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field Bytes) In(values ...[]byte) Expr {
	return field.innerIn(intoAnySlice(values))
}

// InAny use expr IN (?)
// value must be a array/slice
func (field Bytes) InAny(value any) Expr {
	return field.innerInAny(value)
}

// NotIn use expr NOT IN (?)
func (field Bytes) NotIn(values ...[]byte) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field Bytes) NotInAny(value any) Expr {
	return field.innerNotInAny(value)
}

// Like use expr LIKE ?
func (field Bytes) Like(value string) Expr {
	return field.innerLike(value)
}

// FuzzyLike use expr LIKE ?, ? contain prefix % and suffix %
// e.g. expr LIKE %value%
func (field Bytes) FuzzyLike(value string) Expr {
	return field.innerLike("%" + value + "%")
}

// LeftLike use expr LIKE ?, ? contain suffix %.
// e.g. expr LIKE value%
func (field Bytes) LeftLike(value string) Expr {
	return field.innerLike(value + "%")
}

// NotLike use expr NOT LIKE ?
func (field Bytes) NotLike(value string) Expr {
	return field.innerNotLike(value)
}

// Regexp use expr REGEXP ?
func (field Bytes) Regexp(value string) Expr {
	return field.innerRegexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field Bytes) NotRegxp(value string) Expr {
	return field.innerNotRegexp(value)
}

// FindInSet FIND_IN_SET(expr, ?)
func (field Bytes) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// FindInSetWith FIND_IN_SET(?, expr)
func (field Bytes) FindInSetWith(target string) Expr {
	return field.innerFindInSetWith(target)
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field Bytes) SubstringIndex(delim string, count int) Bytes {
	return Bytes{
		expr{
			col: field.col,
			e: clause.Expr{
				SQL:  fmt.Sprintf("SUBSTRING_INDEX(?,%q,%d)", delim, count),
				Vars: []any{field.RawExpr()},
			},
			buildOpts: field.buildOpts,
		},
	}
}

// Value set value
func (field Bytes) Value(value []byte) AssignExpr {
	return field.value(value)
}

// ValueZero set value zero
func (field Bytes) ValueZero() AssignExpr {
	return field.value([]byte{})
}
