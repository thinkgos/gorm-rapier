package rapier

import (
	"time"

	"gorm.io/gorm/clause"
)

// Time time type field
type Time Field

// NewTime bew time field
func NewTime(table, column string, opts ...Option) Time {
	return Time{expr: expr{col: intoClauseColumn(table, column, opts...)}}
}

// IfNull use IFNULL(expr,?)
func (field Time) IfNull(value time.Time) Expr {
	return field.innerIfNull(value)
}

// Eq equal to, use expr = ?
func (field Time) Eq(value time.Time) Expr {
	return field.innerEq(value)
}

// Neq not equal to, use expr <> ?
func (field Time) Neq(value time.Time) Expr {
	return field.innerNeq(value)
}

// Gt greater than, use expr > ?
func (field Time) Gt(value time.Time) Expr {
	return field.innerGt(value)
}

// Gte greater or equal to, use expr >= ?
func (field Time) Gte(value time.Time) Expr {
	return field.innerGte(value)
}

// Lt less than, use expr < ?
func (field Time) Lt(value time.Time) Expr {
	return field.innerLt(value)
}

// Lte less or equal to, use expr <= ?
func (field Time) Lte(value time.Time) Expr {
	return field.innerLte(value)
}

// Between use expr BETWEEN ? AND ?
func (field Time) Between(left time.Time, right time.Time) Expr {
	return field.innerBetween(left, right)
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field Time) NotBetween(left time.Time, right time.Time) Expr {
	return field.innerNotBetween(left, right)
}

// In use expr IN (?)
func (field Time) In(values ...time.Time) Expr {
	return field.innerIn(intoAnySlice(values))
}

// InAny use expr IN (?)
// value must be a array/slice
func (field Time) InAny(value any) Expr {
	return field.innerInAny(value)
}

// NotIn use expr NOT IN (?)
func (field Time) NotIn(values ...time.Time) Expr {
	return field.innerNotIn(intoAnySlice(values))
}

// NotInAny use expr NOT IN (?)
// value must be a array/slice
func (field Time) NotInAny(value any) Expr {
	return field.innerNotInAny(value)
}

// Sum use SUM(expr)
func (field Time) Sum() Time {
	return Time{field.innerSum()}
}

// Add use DATE_ADD(date, INTERVAL ? MICROSECOND)
func (field Time) Add(value time.Duration) Time {
	return Time{field.innerAdd(value)}
}

// Sub use DATE_SUB(date, INTERVAL ? MICROSECOND)
func (field Time) Sub(value time.Duration) Time {
	return Time{field.innerSub(value)}
}

// FindInSet use FIND_IN_SET(expr, ?)
func (field Time) FindInSet(targetList string) Expr {
	return field.innerFindInSet(targetList)
}

// UnixTimestamp use UnixTimestamp(date)
func (field Time) UnixTimestamp() Int64 {
	return Int64{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Date use DATE(expr) return the date.
func (field Time) Date() Time {
	return Time{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATE(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Year use YEAR(date) return the year.
func (field Time) Year() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "YEAR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Month use MONTH(date) return the month.
func (field Time) Month() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MONTH(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Day use DAY(date) return the day.
func (field Time) Day() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAY(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Hour use HOUR(date) return the hour.
func (field Time) Hour() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "HOUR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Minute use MINUTE(date) return the minute.
func (field Time) Minute() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MINUTE(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// Second use SECOND(date) return the second.
func (field Time) Second() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "SECOND(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// MicroSecond use MICROSECOND(date) return the microsecond.
func (field Time) MicroSecond() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MICROSECOND(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfWeek use DAYOFWEEK(date)
func (field Time) DayOfWeek() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFWEEK(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfMonth use DAYOFMONTH(date)
func (field Time) DayOfMonth() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFMONTH(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayOfYear use DAYOFYEAR(date)
func (field Time) DayOfYear() Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYOFYEAR(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// DateDiff use DATEDIFF(expr1, expr2)
func (field Time) DateDiff(expr2 time.Time) Int {
	return Int{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATEDIFF(?,?)", Vars: []any{field.RawExpr(), expr2}},
			buildOpts: field.buildOpts,
		},
	}
}

// DateFormat use DATE_FORMAT(date,format)
func (field Time) DateFormat(format string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DATE_FORMAT(?,?)", Vars: []any{field.RawExpr(), format}},
			buildOpts: field.buildOpts,
		},
	}
}

// DayName use DAYNAME(date) return the name of the day of the week.
func (field Time) DayName() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "DAYNAME(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// MonthName use MONTHNAME(date) return the name of the month of the year.
func (field Time) MonthName() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "MONTHNAME(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// AddCol use expr1 + expr2
func (e Time) AddCol(col Expr) Time {
	return Time{e.innerAddCol(col)}
}

// SubCol use expr1 - expr2
func (e Time) SubCol(col Expr) Time {
	return Time{e.innerSubCol(col)}
}

// MulCol use (expr1) * (expr2)
func (e Time) MulCol(col Expr) Time {
	return Time{e.innerMulCol(col)}
}

// DivCol use (expr1) / (expr2)
func (e Time) DivCol(col Expr) Time {
	return Time{e.innerDivCol(col)}
}

// Value set value
func (field Time) Value(value time.Time) SetExpr {
	return field.value(value)
}

// ValuePointer set value use pointer
func (field Time) ValuePointer(value *time.Time) SetExpr {
	return field.value(value)
}

// ValueZero set value zero
func (field Time) ValueZero() SetExpr {
	return field.value(time.Time{})
}
