package assist

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

var _ Expr = (*Field)(nil)

// Field a standard field struct
type Field struct{ expr }
type Raw = Field

// NewField new field
func NewField(table, column string, opts ...Option) Field {
	return Field{expr: expr{col: intoColumn(table, column, opts...)}}
}

// NewRaw new raw
func NewRaw(sql string, vars ...any) Raw {
	return Raw{expr: expr{
		e: clause.Expr{
			SQL:                sql,
			Vars:               vars,
			WithoutParentheses: false,
		},
	}}
}

// IfNull use IFNULL(expr,?)
func (field Field) IfNull(value any) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field Field) Eq(value any) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to, use expr <> ?
func (field Field) Neq(value any) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than, use expr > ?
func (field Field) Gt(value any) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to, use expr >= ?
func (field Field) Gte(value any) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than, use expr < ?
func (field Field) Lt(value any) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to, use expr <= ?
func (field Field) Lte(value any) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between use expr BETWEEN ? AND ?
func (field Field) Between(left any, right any) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Field) NotBetween(left any, right any) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field Field) In(values ...any) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)}}
}

// NotIn use expr NOT IN (?)
func (field Field) NotIn(values ...any) Expr {
	return expr{e: clause.Not(clause.IN{Column: field.RawExpr(), Values: intoSlice(values...)})}
}

// Like use expr LIKE ?
func (field Field) Like(value any) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// NotLike use expr NOT LIKE ?
func (field Field) NotLike(value any) Expr {
	return expr{e: clause.Not(clause.Like{Column: field.RawExpr(), Value: value})}
}

// Sum use SUM(expr)
func (field Field) Sum() Field {
	return Field{field.sum()}
}

// Add use
// value type:
//
//	time.Duration: use DATE_ADD(expr, INTERVAL ? MICROSECOND)
//	other: use expr+?
func (field Field) Add(value any) Field {
	return Field{field.add(value)}
}

// Sub use below
// value type:
//
//	time.Duration: use DATE_SUB(expr, INTERVAL ? MICROSECOND)
//	other: use expr-?
func (field Field) Sub(value any) Field {
	return Field{field.sub(value)}
}

// Mul use expr*?
func (field Field) Mul(value any) Field {
	return Field{field.mul(value)}
}

// Div use expr/?
func (field Field) Div(value any) Field {
	return Field{field.div(value)}
}

// Mod use expr%?
func (field Field) Mod(value any) Field {
	return Field{field.mod(value)}
}

// FloorDiv use expr DIV ?
func (field Field) FloorDiv(value any) Field {
	return Field{field.floorDiv(value)}
}

// Floor se FLOOR(expr)
func (field Field) Floor() Field {
	return Field{field.floor()}
}

// Round use ROUND(expr, ?)
func (field Field) Round(decimals int) Field {
	return Field{field.round(decimals)}
}

// RightShift use expr>>?
func (field Field) RightShift(value any) Field {
	return Field{field.rightShift(value)}
}

// LeftShift use expr<<?
func (field Field) LeftShift(value any) Field {
	return Field{field.leftShift(value)}
}

// BitXor use expr expr^?
func (field Field) BitXor(value any) Field {
	return Field{field.bitXor(value)}
}

// BitAnd use expr expr&?
func (field Field) BitAnd(value any) Field {
	return Field{field.bitAnd(value)}
}

// BitOr use expr expr|?
func (field Field) BitOr(value any) Field {
	return Field{field.bitOr(value)}
}

// BitFlip use expr ~expr
func (field Field) BitFlip() Field {
	return Field{field.bitFlip()}
}

// Regexp use expr REGEXP ?
func (field Field) Regexp(value any) Expr {
	return field.regexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field Field) NotRegxp(value string) Expr {
	return field.notRegexp(value)
}

// FindInSet use FIND_IN_SET(field_name, input_string_list)
func (field Field) FindInSet(targetList string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{field.RawExpr(), targetList}}}
}

// FindInSetWith use FIND_IN_SET(?, expr)
func (field Field) FindInSetWith(target string) Expr {
	return expr{e: clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{target, field.RawExpr()}}}
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field Field) SubstringIndex(delim string, count int) String {
	return String{expr{e: clause.Expr{
		SQL:  fmt.Sprintf("SUBSTRING_INDEX(?,%q,%d)", delim, count),
		Vars: []any{field.RawExpr()},
	}}}
}

// Replace use REPLACE(expr,?,?)
func (field Field) Replace(from, to string) String {
	return String{expr{e: clause.Expr{SQL: "REPLACE(?,?,?)", Vars: []any{field.RawExpr(), from, to}}}}
}

// Concat use CONCAT(?,?,?)
func (field Field) Concat(before, after string) String {
	switch {
	case before != "" && after != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?,?)", Vars: []any{before, field.RawExpr(), after}}}}
	case before != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{before, field.RawExpr()}}}}
	case after != "":
		return String{expr{e: clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{field.RawExpr(), after}}}}
	default:
		return String{
			expr: expr{
				col:       field.col,
				e:         field.e,
				buildOpts: field.buildOpts,
			},
		}
	}
}

// FromUnixTime use FromUnixTime(unix_timestamp[, format])
func (field Field) FromUnixTime(format ...string) String {
	if len(format) > 0 && strings.TrimSpace(format[0]) != "" {
		return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?, ?)", Vars: []any{field.RawExpr(), format[0]}}}}
	}
	return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?)", Vars: []any{field.RawExpr()}}}}
}

// FromDays use FROM_DAYS(value)
func (field Field) FromDays() Time {
	return Time{expr{e: clause.Expr{SQL: "FROM_DAYS(?)", Vars: []any{field.RawExpr()}}}}
}

// UnixTimestamp use UnixTimestamp(date)
func (field Field) UnixTimestamp() Int64 {
	return Int64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []any{field.RawExpr()}}}}
}

// Date use DATE(expr) return the date.
func (field Field) Date() Time {
	return Time{expr{e: clause.Expr{SQL: "DATE(?)", Vars: []any{field.RawExpr()}}}}
}

// Year use YEAR(date) return the year.
func (field Field) Year() Int {
	return Int{expr{e: clause.Expr{SQL: "YEAR(?)", Vars: []any{field.RawExpr()}}}}
}

// Month use MONTH(date) return the month.
func (field Field) Month() Int {
	return Int{expr{e: clause.Expr{SQL: "MONTH(?)", Vars: []any{field.RawExpr()}}}}
}

// Day use DAY(date) return the day.
func (field Field) Day() Int {
	return Int{expr{e: clause.Expr{SQL: "DAY(?)", Vars: []any{field.RawExpr()}}}}
}

// Hour use HOUR(date) return the hour.
func (field Field) Hour() Int {
	return Int{expr{e: clause.Expr{SQL: "HOUR(?)", Vars: []any{field.RawExpr()}}}}
}

// Minute use MINUTE(date) return the minute.
func (field Field) Minute() Int {
	return Int{expr{e: clause.Expr{SQL: "MINUTE(?)", Vars: []any{field.RawExpr()}}}}
}

// Second use SECOND(date) return the second.
func (field Field) Second() Int {
	return Int{expr{e: clause.Expr{SQL: "SECOND(?)", Vars: []any{field.RawExpr()}}}}
}

// MicroSecond use MICROSECOND(date) return the microsecond.
func (field Field) MicroSecond() Int {
	return Int{expr{e: clause.Expr{SQL: "MICROSECOND(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfWeek use DAYOFWEEK(date)
func (field Field) DayOfWeek() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFWEEK(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfMonth use DAYOFMONTH(date)
func (field Field) DayOfMonth() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFMONTH(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfYear use DAYOFYEAR(date)
func (field Field) DayOfYear() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFYEAR(?)", Vars: []any{field.RawExpr()}}}}
}

// DateDiff use DATEDIFF(expr1, expr2)
func (field Field) DateDiff(expr2 time.Time) Int {
	return Int{expr{e: clause.Expr{SQL: "DATEDIFF(?,?)", Vars: []any{field.RawExpr(), expr2}}}}
}

// DateFormat use DATE_FORMAT(date,format)
func (field Field) DateFormat(format string) String {
	return String{expr{e: clause.Expr{SQL: "DATE_FORMAT(?,?)", Vars: []any{field.RawExpr(), format}}}}
}

// DayName use DAYNAME(date) return the name of the day of the week.
func (field Field) DayName() String {
	return String{expr{e: clause.Expr{SQL: "DAYNAME(?)", Vars: []any{field.RawExpr()}}}}
}

// MonthName use MONTHNAME(date) return the name of the month of the year.
func (field Field) MonthName() String {
	return String{expr{e: clause.Expr{SQL: "MONTHNAME(?)", Vars: []any{field.RawExpr()}}}}
}
