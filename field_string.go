package assist

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// String string type field
type String Field

// NewString new string field.
func NewString(table, column string, opts ...Option) String {
	return String{expr: expr{col: intoColumn(table, column, opts...)}}
}

// IfNull use IFNULL(expr,?)
func (field String) IfNull(value string) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field String) Eq(value string) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to, use expr <> ?
func (field String) Neq(value string) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than, use expr > ?
func (field String) Gt(value string) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to, use expr >= ?
func (field String) Gte(value string) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than, use expr < ?
func (field String) Lt(value string) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to, use expr <= ?
func (field String) Lte(value string) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between use expr BETWEEN ? AND ?
func (field String) Between(left, right string) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field String) NotBetween(left, right string) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field String) In(values ...string) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)}}
}

// NotIn use expr NOT IN (?)
func (field String) NotIn(values ...string) Expr {
	return expr{e: clause.Not(clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)})}
}

// Like use expr LIKE ?
func (field String) Like(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// FuzzyLike use expr LIKE ?, ? contain prefix % and suffix %
// e.g. expr LIKE %value%
func (field String) FuzzyLike(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: "%" + value + "%"}}
}

// LeftLike use expr LIKE ?, ? contain suffix %.
// e.g. expr LIKE value%
func (field String) LeftLike(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value + "%"}}
}

// NotLike use expr NOT LIKE ?
func (field String) NotLike(value string) Expr {
	return expr{e: clause.Not(clause.Like{Column: field.RawExpr(), Value: value})}
}

// Regexp use expr REGEXP ?
func (field String) Regexp(value string) Expr {
	return field.regexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field String) NotRegxp(value string) Expr {
	return field.notRegexp(value)
}

// FindInSet equal to FIND_IN_SET(field_name, input_string_list)
func (field String) FindInSet(targetList string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{field.RawExpr(), targetList}}}
}

// FindInSetWith equal to FIND_IN_SET(input_string, field_name)
func (field String) FindInSetWith(target string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{target, field.RawExpr()}}}
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field String) SubstringIndex(delim string, count int) String {
	return String{expr{e: clause.Expr{
		SQL:  fmt.Sprintf("SUBSTRING_INDEX(?,%q,%d)", delim, count),
		Vars: []any{field.RawExpr()},
	}}}
}

// Replace use REPLACE(expr,?,?)
func (field String) Replace(from, to string) String {
	return String{expr{e: clause.Expr{SQL: "REPLACE(?,?,?)", Vars: []any{field.RawExpr(), from, to}}}}
}

// Concat use CONCAT(?,?,?)
func (field String) Concat(before, after string) String {
	switch {
	case before != "" && after != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?,?)", Vars: []any{before, field.RawExpr(), after}}}}
	case before != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{before, field.RawExpr()}}}}
	case after != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{field.RawExpr(), after}}}}
	default:
		return field
	}
}
