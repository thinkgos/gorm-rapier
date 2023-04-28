package assist

import (
	"time"

	"gorm.io/gorm/clause"
)

// Time time type field
type Time Field

// NewTime bew time field
func NewTime(table, column string, opts ...Option) Time {
	return Time{expr: expr{col: intoColumn(table, column, opts...)}}
}

// IfNull use IFNULL(expr,?)
func (field Time) IfNull(value time.Time) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field Time) Eq(value time.Time) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to, use expr <> ?
func (field Time) Neq(value time.Time) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than, use expr > ?
func (field Time) Gt(value time.Time) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to, use expr >= ?
func (field Time) Gte(value time.Time) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than, use expr < ?
func (field Time) Lt(value time.Time) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to, use expr <= ?
func (field Time) Lte(value time.Time) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between use expr BETWEEN ? AND ?
func (field Time) Between(left time.Time, right time.Time) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Time) NotBetween(left time.Time, right time.Time) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field Time) In(values ...time.Time) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)}}
}

// NotIn use expr NOT IN (?)
func (field Time) NotIn(values ...time.Time) Expr {
	return expr{e: clause.Not(clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)})}
}

// Sum use SUM(expr)
func (field Time) Sum() Time {
	return Time{field.sum()}
}

// Add use DATE_ADD(date, INTERVAL ? MICROSECOND)
func (field Time) Add(value time.Duration) Time {
	return Time{field.add(value)}
}

// Sub use DATE_SUB(date, INTERVAL ? MICROSECOND)
func (field Time) Sub(value time.Duration) Time {
	return Time{field.sub(value)}
}

// UnixTimestamp use UnixTimestamp(date)
func (field Time) UnixTimestamp() Int64 {
	return Int64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []any{field.RawExpr()}}}}
}

// Date use DATE(expr) return the date.
func (field Time) Date() Time {
	return Time{expr{e: clause.Expr{SQL: "DATE(?)", Vars: []any{field.RawExpr()}}}}
}

// Year use YEAR(date) return the year.
func (field Time) Year() Int {
	return Int{expr{e: clause.Expr{SQL: "YEAR(?)", Vars: []any{field.RawExpr()}}}}
}

// Month use MONTH(date) return the month.
func (field Time) Month() Int {
	return Int{expr{e: clause.Expr{SQL: "MONTH(?)", Vars: []any{field.RawExpr()}}}}
}

// Day use DAY(date) return the day.
func (field Time) Day() Int {
	return Int{expr{e: clause.Expr{SQL: "DAY(?)", Vars: []any{field.RawExpr()}}}}
}

// Hour use HOUR(date) return the hour.
func (field Time) Hour() Int {
	return Int{expr{e: clause.Expr{SQL: "HOUR(?)", Vars: []any{field.RawExpr()}}}}
}

// Minute use MINUTE(date) return the minute.
func (field Time) Minute() Int {
	return Int{expr{e: clause.Expr{SQL: "MINUTE(?)", Vars: []any{field.RawExpr()}}}}
}

// Second use SECOND(date) return the second.
func (field Time) Second() Int {
	return Int{expr{e: clause.Expr{SQL: "SECOND(?)", Vars: []any{field.RawExpr()}}}}
}

// MicroSecond use MICROSECOND(date) return the microsecond.
func (field Time) MicroSecond() Int {
	return Int{expr{e: clause.Expr{SQL: "MICROSECOND(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfWeek use DAYOFWEEK(date)
func (field Time) DayOfWeek() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFWEEK(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfMonth use DAYOFMONTH(date)
func (field Time) DayOfMonth() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFMONTH(?)", Vars: []any{field.RawExpr()}}}}
}

// DayOfYear use DAYOFYEAR(date)
func (field Time) DayOfYear() Int {
	return Int{expr{e: clause.Expr{SQL: "DAYOFYEAR(?)", Vars: []any{field.RawExpr()}}}}
}

// DateDiff use DATEDIFF(expr1, expr2)
func (field Time) DateDiff(expr2 time.Time) Int {
	return Int{expr{e: clause.Expr{SQL: "DATEDIFF(?,?)", Vars: []any{field.RawExpr(), expr2}}}}
}

// DateFormat use DATE_FORMAT(date,format)
func (field Time) DateFormat(format string) String {
	return String{expr{e: clause.Expr{SQL: "DATE_FORMAT(?,?)", Vars: []any{field.RawExpr(), format}}}}
}

// DayName use DAYNAME(date) return the name of the day of the week.
func (field Time) DayName() String {
	return String{expr{e: clause.Expr{SQL: "DAYNAME(?)", Vars: []any{field.RawExpr()}}}}
}

// MonthName use MONTHNAME(date) return the name of the month of the year.
func (field Time) MonthName() String {
	return String{expr{e: clause.Expr{SQL: "MONTHNAME(?)", Vars: []any{field.RawExpr()}}}}
}

// IntoColumns columns array with sub method
func (field Time) IntoColumns() Columns {
	return NewColumns(field)
}
