package assist

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// String string type field
type String Field

// NewString new string field.
func NewString(table, column string, opts ...Option) String {
	return String{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IfNull use IFNULL(expr,?)
func (field String) IfNull(value string) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field String) Eq(value string) Expr {
	return field.innerEq(value)
}

// Neq not equal to, use expr <> ?
func (field String) Neq(value string) Expr {
	return field.innerNeq(value)
}

// Gt greater than, use expr > ?
func (field String) Gt(value string) Expr {
	return field.innerGt(value)
}

// Gte greater or equal to, use expr >= ?
func (field String) Gte(value string) Expr {
	return field.innerGte(value)
}

// Lt less than, use expr < ?
func (field String) Lt(value string) Expr {
	return field.innerLt(value)
}

// Lte less or equal to, use expr <= ?
func (field String) Lte(value string) Expr {
	return field.innerLte(value)
}

// Between use expr BETWEEN ? AND ?
func (field String) Between(left, right string) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field String) NotBetween(left, right string) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field String) In(values ...string) Expr {
	return field.innerIn(intoAnySlice(values))
}

// InAny use expr IN (?)
// value must be a array/slice
func (field String) InAny(value any) Expr {
	return field.innerInAny(value)
}

// NotIn use expr NOT IN (?)
func (field String) NotIn(values ...string) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field String) NotInAny(value any) Expr {
	return field.innerNotInAny(value)
}

// Like use expr LIKE ?
func (field String) Like(value string) Expr {
	return field.innerLike(value)
}

// FuzzyLike use expr LIKE ?, ? contain prefix % and suffix %
// e.g. expr LIKE %value%
func (field String) FuzzyLike(value string) Expr {
	return field.innerLike("%" + value + "%")
}

// LeftLike use expr LIKE ?, ? contain suffix %.
// e.g. expr LIKE value%
func (field String) LeftLike(value string) Expr {
	return field.innerLike(value + "%")
}

// NotLike use expr NOT LIKE ?
func (field String) NotLike(value string) Expr {
	return field.innerNotLike(value)
}

// Regexp use expr REGEXP ?
func (field String) Regexp(value string) Expr {
	return field.innerRegexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field String) NotRegxp(value string) Expr {
	return field.innerNotRegexp(value)
}

// FindInSet equal to FIND_IN_SET(expr, ?)
func (field String) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// FindInSetWith equal to FIND_IN_SET(?, expr)
func (field String) FindInSetWith(target string) Expr {
	return field.innerFindInSetWith(target)
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field String) SubstringIndex(delim string, count int) String {
	return String{
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

// Replace use REPLACE(expr,?,?)
func (field String) Replace(from, to string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "REPLACE(?,?,?)", Vars: []any{field.RawExpr(), from, to}},
			buildOpts: field.buildOpts,
		},
	}
}

// Hidden use CONCAT(LEFT(expr,left),pad,RIGHT(expr,right))
// detail see innerHidden
func (field String) Hidden(left, right int, pad string) String {
	return String{field.innerHidden(left, right, pad)}
}

// HiddenSuffix use CONCAT(LEFT(expr,left),pad)
func (field String) HiddenSuffix(left int, pad string) String {
	return String{field.innerHiddenSuffix(left, pad)}
}

// HiddenPrefix use CONCAT(pad,RIGHT(expr,right))
func (field String) HiddenPrefix(right int, pad string) String {
	return String{field.innerHiddenPrefix(right, pad)}
}

// Concat use CONCAT(?,?,?)
func (field String) Concat(before, after string) String {
	var e clause.Expression

	switch {
	case before != "" && after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?,?)", Vars: []any{before, field.RawExpr(), after}}
	case before != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{before, field.RawExpr()}}
	case after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{field.RawExpr(), after}}
	default:
		return field
	}
	return String{
		expr{
			col:       field.col,
			e:         e,
			buildOpts: field.buildOpts,
		},
	}
}

// Trim use TRIM(BOTH ? FROM ?)
func (field String) Trim(remStr string) String {
	return String{field.innerTrim(remStr)}
}

// LTrim use TRIM(LEADING ? FROM ?)
func (field String) LTrim(remStr string) String {
	return String{field.innerLTrim(remStr)}
}

// RTrim use TRIM(TRAILING ? FROM ?)
func (field String) RTrim(remStr string) String {
	return String{field.innerRTrim(remStr)}
}

// TrimSpace use TRIM(?)
func (field String) TrimSpace() String {
	return String{field.innerTrimSpace()}
}

// LTrimSpace use LTRIM(?)
func (field String) LTrimSpace() String {
	return String{field.innerLTrimSpace()}
}

// RTrimSpace use RTRIM(?)
func (field String) RTrimSpace() String {
	return String{field.innerRTrimSpace()}
}

// AddCol use expr1 + expr2
func (field String) AddCol(col Expr) String {
	return String{field.innerAddCol(col)}
}

// SubCol use expr1 - expr2
func (field String) SubCol(col Expr) String {
	return String{field.innerSubCol(col)}
}

// MulCol use (expr1) * (expr2)
func (field String) MulCol(col Expr) String {
	return String{field.innerMulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (field String) DivCol(col Expr) String {
	return String{field.innerDivCol(col)}
}

// ConcatCol use CONCAT(expr1,exp2...exprN)
func (field String) ConcatCol(cols ...Expr) String {
	return String{field.innerConcatCol(cols...)}
}

// IntoColumns columns array with sub method
func (field String) IntoColumns() Columns {
	return NewColumns(field)
}
