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
	return Field{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// NewRaw new raw
func NewRaw(sql string, vars ...any) Raw {
	return Raw{
		expr{
			e: clause.Expr{
				SQL:                sql,
				Vars:               vars,
				WithoutParentheses: false,
			},
		},
	}
}

// IfNull use IFNULL(expr,?)
func (field Field) IfNull(value any) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field Field) Eq(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Eq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Neq not equal to, use expr <> ?
func (field Field) Neq(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Neq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gt greater than, use expr > ?
func (field Field) Gt(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gte greater or equal to, use expr >= ?
func (field Field) Gte(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lt less than, use expr < ?
func (field Field) Lt(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lte less or equal to, use expr <= ?
func (field Field) Lte(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Between use expr BETWEEN ? AND ?
func (field Field) Between(left any, right any) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Field) NotBetween(left any, right any) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field Field) In(values ...any) Expr {
	return field.innerIn(intoAnySlice(values))
}

// NotIn use expr NOT IN (?)
func (field Field) NotIn(values ...any) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// Like use expr LIKE ?
func (field Field) Like(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Like{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// NotLike use expr NOT LIKE ?
func (field Field) NotLike(value any) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(clause.Like{Column: field.RawExpr(), Value: value}),
		buildOpts: field.buildOpts,
	}
}

// Sum use SUM(expr)
func (field Field) Sum() Field {
	return Field{field.innerSum()}
}

// Add use
// value type:
//
//	time.Duration: use DATE_ADD(expr, INTERVAL ? MICROSECOND)
//	other: use expr+?
func (field Field) Add(value any) Field {
	return Field{field.innerAdd(value)}
}

// Sub use below
// value type:
//
//	time.Duration: use DATE_SUB(expr, INTERVAL ? MICROSECOND)
//	other: use expr-?
func (field Field) Sub(value any) Field {
	return Field{field.innerSub(value)}
}

// Mul use expr*?
func (field Field) Mul(value any) Field {
	return Field{field.innerMul(value)}
}

// Div use expr/?
func (field Field) Div(value any) Field {
	return Field{field.innerDiv(value)}
}

// Mod use expr%?
func (field Field) Mod(value any) Field {
	return Field{field.innerMod(value)}
}

// FloorDiv use expr DIV ?
func (field Field) FloorDiv(value any) Field {
	return Field{field.innerFloorDiv(value)}
}

// Floor se FLOOR(expr)
func (field Field) Floor() Field {
	return Field{field.innerFloor()}
}

// Round use ROUND(expr, ?)
func (field Field) Round(decimals int) Field {
	return Field{field.innerRound(decimals)}
}

// RightShift use expr>>?
func (field Field) RightShift(value any) Field {
	return Field{field.innerRightShift(value)}
}

// LeftShift use expr<<?
func (field Field) LeftShift(value any) Field {
	return Field{field.innerLeftShift(value)}
}

// BitXor use expr expr^?
func (field Field) BitXor(value any) Field {
	return Field{field.innerBitXor(value)}
}

// BitAnd use expr expr&?
func (field Field) BitAnd(value any) Field {
	return Field{field.innerBitAnd(value)}
}

// BitOr use expr expr|?
func (field Field) BitOr(value any) Field {
	return Field{field.innerBitOr(value)}
}

// BitFlip use expr ~expr
func (field Field) BitFlip() Field {
	return Field{field.innerBitFlip()}
}

// Regexp use expr REGEXP ?
func (field Field) Regexp(value any) Expr {
	return field.innerRegexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field Field) NotRegxp(value string) Expr {
	return field.innerNotRegexp(value)
}

// FindInSet use FIND_IN_SET(expr, ?)
func (field Field) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// FindInSetWith use FIND_IN_SET(?, expr)
func (field Field) FindInSetWith(target string) Expr {
	return field.innerFindInSetWith(target)
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field Field) SubstringIndex(delim string, count int) String {
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
func (field Field) Replace(from, to string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "REPLACE(?,?,?)", Vars: []any{field.RawExpr(), from, to}},
			buildOpts: field.buildOpts,
		},
	}
}

// Concat use CONCAT(?,?,?)
func (field Field) Concat(before, after string) String {
	var e clause.Expression

	switch {
	case before != "" && after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?,?)", Vars: []any{before, field.RawExpr(), after}}
	case before != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{before, field.RawExpr()}}
	case after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{field.RawExpr(), after}}
	default:
		e = field.e
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
func (field Field) Trim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(BOTH ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// LTrim use TRIM(LEADING ? FROM ?)
func (field Field) LTrim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(LEADING ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// RTrim use TRIM(TRAILING ? FROM ?)
func (field Field) RTrim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(TRAILING ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// TrimSpace use TRIM(?)
func (field Field) TrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// LTrimSpace use LTRIM(?)
func (field Field) LTrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "LTRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// RTrimSpace use RTRIM(?)
func (field Field) RTrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "RTRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// FromUnixTime use FromUnixTime(unix_timestamp[, format])
func (field Field) FromUnixTime(format ...string) String {
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
func (field Field) FromDays() Time {
	return Time{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "FROM_DAYS(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// UnixTimestamp use UnixTimestamp(date)
func (field Field) UnixTimestamp() Int64 {
	return Int64{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Date use DATE(expr) return the date.
func (field Field) Date() Time {
	return Time{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATE(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Year use YEAR(date) return the year.
func (field Field) Year() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "YEAR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Month use MONTH(date) return the month.
func (field Field) Month() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MONTH(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Day use DAY(date) return the day.
func (field Field) Day() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAY(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Hour use HOUR(date) return the hour.
func (field Field) Hour() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "HOUR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Minute use MINUTE(date) return the minute.
func (field Field) Minute() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MINUTE(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Second use SECOND(date) return the second.
func (field Field) Second() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "SECOND(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// MicroSecond use MICROSECOND(date) return the microsecond.
func (field Field) MicroSecond() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MICROSECOND(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfWeek use DAYOFWEEK(date)
func (field Field) DayOfWeek() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFWEEK(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfMonth use DAYOFMONTH(date)
func (field Field) DayOfMonth() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFMONTH(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfYear use DAYOFYEAR(date)
func (field Field) DayOfYear() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFYEAR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DateDiff use DATEDIFF(expr1, expr2)
func (field Field) DateDiff(expr2 time.Time) Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATEDIFF(?,?)", Vars: []any{field.RawExpr(), expr2}},
			buildOpts: field.buildOpts,
		},
	}
}

// DateFormat use DATE_FORMAT(date,format)
func (field Field) DateFormat(format string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATE_FORMAT(?,?)", Vars: []any{field.RawExpr(), format}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayName use DAYNAME(date) return the name of the day of the week.
func (field Field) DayName() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYNAME(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// MonthName use MONTHNAME(date) return the name of the month of the year.
func (field Field) MonthName() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MONTHNAME(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// AddCol use expr1 + expr2
func (e Field) AddCol(col Expr) Field {
	return Field{e.addCol(col)}
}

// SubCol use expr1 - expr2
func (e Field) SubCol(col Expr) Field {
	return Field{e.subCol(col)}
}

// MulCol use (expr1) * (expr2)
func (e Field) MulCol(col Expr) Field {
	return Field{e.mulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (e Field) DivCol(col Expr) Field {
	return Field{e.divCol(col)}
}

// ConcatCol use CONCAT(expr1,exp2...exprN)
func (e Field) ConcatCol(cols ...Expr) Field {
	return Field{e.concatCol(cols...)}
}

// IntoColumns columns array with sub method
func (field Field) IntoColumns() Columns {
	return NewColumns(field)
}
