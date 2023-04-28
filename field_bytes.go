package assist

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// Bytes []byte type field
type Bytes Field

// NewBytes new bytes field.
func NewBytes(table, column string, opts ...Option) Bytes {
	return Bytes{expr: expr{col: intoClauseColumn(table, column, opts...)}}
}

// IfNull use IFNULL(expr,?)
func (field Bytes) IfNull(value []byte) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field Bytes) Eq(value []byte) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to, use expr <> ?
func (field Bytes) Neq(value []byte) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than, use expr > ?
func (field Bytes) Gt(value []byte) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to, use expr >= ?
func (field Bytes) Gte(value []byte) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than, use expr < ?
func (field Bytes) Lt(value []byte) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to, use expr <= ?
func (field Bytes) Lte(value []byte) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between use expr BETWEEN ? AND ?
func (field Bytes) Between(left []byte, right []byte) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Bytes) NotBetween(left []byte, right []byte) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field Bytes) In(values ...[]byte) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)}}
}

// NotIn use expr NOT IN (?)
func (field Bytes) NotIn(values ...[]byte) Expr {
	return expr{e: clause.Not(clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)})}
}

// Like use expr LIKE ?
func (field Bytes) Like(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// FuzzyLike use expr LIKE ?, ? contain prefix % and suffix %
// e.g. expr LIKE %value%
func (field Bytes) FuzzyLike(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: "%" + value + "%"}}
}

// LeftLike use expr LIKE ?, ? contain suffix %.
// e.g. expr LIKE value%
func (field Bytes) LeftLike(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value + "%"}}
}

// NotLike use expr NOT LIKE ?
func (field Bytes) NotLike(value string) Expr {
	return expr{e: clause.Not(clause.Like{Column: field.RawExpr(), Value: value})}
}

// Regexp use expr REGEXP ?
func (field Bytes) Regexp(value string) Expr {
	return field.regexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field Bytes) NotRegxp(value string) Expr {
	return field.notRegexp(value)
}

// FindInSet FIND_IN_SET(field_name, input_string_list)
func (field Bytes) FindInSet(targetList string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{field.RawExpr(), targetList}}}
}

// FindInSetWith FIND_IN_SET(input_string, field_name)
func (field Bytes) FindInSetWith(target string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{target, field.RawExpr()}}}
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field Bytes) SubstringIndex(delim string, count int) Bytes {
	return Bytes{expr{e: clause.Expr{
		SQL:  fmt.Sprintf("SUBSTRING_INDEX(?,%q,%d)", delim, count),
		Vars: []any{field.RawExpr()},
	}}}
}

// IntoColumns columns array with sub method
func (field Bytes) IntoColumns() Columns {
	return NewColumns(field)
}
