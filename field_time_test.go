package assist

import (
	"testing"
	"time"
)

func Test_Expr_Time(t *testing.T) {
	value1, _ := time.Parse("2006-01-02 15:04:05", "2021-06-29 15:11:49")
	value2 := value1.Add(1 * time.Hour)
	value3 := value1.Add(2 * time.Hour)
	value0 := value1.Add(24 * time.Hour)
	value4 := []time.Time{value1, value2, value3}
	value5 := []TestTime{TestTime(value1), TestTime(value2), TestTime(value3)}
	value6 := []string{"1", "2", "3"}

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewTime("", "created_at").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`created_at`,?)",
		},
		{
			name:     "eq",
			expr:     NewTime("", "created_at").Eq(value1),
			wantVars: []any{value1},
			want:     "`created_at` = ?",
		},
		{
			name:     "neq",
			expr:     NewTime("", "created_at").Neq(value1),
			wantVars: []any{value1},
			want:     "`created_at` <> ?",
		},
		{
			name:     "gt",
			expr:     NewTime("", "created_at").Gt(value1),
			wantVars: []any{value1},
			want:     "`created_at` > ?",
		},
		{
			name:     "gte",
			expr:     NewTime("", "created_at").Gte(value1),
			wantVars: []any{value1},
			want:     "`created_at` >= ?",
		},
		{
			name:     "lt",
			expr:     NewTime("", "created_at").Lt(value1),
			wantVars: []any{value1},
			want:     "`created_at` < ?",
		},
		{
			name:     "lte",
			expr:     NewTime("", "created_at").Lte(value1),
			wantVars: []any{value1},
			want:     "`created_at` <= ?",
		},
		{
			name:     "between",
			expr:     NewTime("", "created_at").Between(value1, value0),
			wantVars: []any{value1, value0},
			want:     "`created_at` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     NewTime("", "created_at").NotBetween(value1, value0),
			wantVars: []any{value1, value0},
			want:     "NOT (`created_at` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     NewTime("", "created_at").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`created_at` IN (?,?,?)",
		},
		{
			name:     "in any current type",
			expr:     NewTime("", "created_at").InAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`created_at` IN (?,?,?)",
		},
		{
			name:     "in any under new type",
			expr:     NewTime("", "created_at").InAny(value5),
			wantVars: []any{TestTime(value1), TestTime(value2), TestTime(value3)},
			want:     "`created_at` IN (?,?,?)",
		},
		{
			name:     "in any under type string",
			expr:     NewTime("", "created_at").InAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`created_at` IN (?,?,?)",
		},
		{
			name:     "in any but not a array/slice",
			expr:     NewTime("", "created_at").InAny(1),
			wantVars: nil,
			want:     "",
		},
		{
			name:     "not in",
			expr:     NewTime("", "created_at").NotIn(value1, value1.Add(1*time.Hour), value1.Add(2*time.Hour)),
			wantVars: []any{value1, value1.Add(1 * time.Hour), value1.Add(2 * time.Hour)},
			want:     "`created_at` NOT IN (?,?,?)",
		},
		{
			name:     "not in any current type",
			expr:     NewTime("", "created_at").NotInAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`created_at` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under new type",
			expr:     NewTime("", "created_at").NotInAny(value5),
			wantVars: []any{TestTime(value1), TestTime(value2), TestTime(value3)},
			want:     "`created_at` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under type string",
			expr:     NewTime("", "created_at").NotInAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`created_at` NOT IN (?,?,?)",
		},
		{
			name:     "not in any but not a array/slice",
			expr:     NewTime("", "created_at").NotInAny(1),
			wantVars: nil,
			want:     "NOT",
		},
		{
			name:     "Sum",
			expr:     NewTime("", "created_at").Sum(),
			wantVars: nil,
			want:     "SUM(`created_at`)",
		},
		{
			name:     "Add use DATE_ADD()",
			expr:     NewTime("", "created_at").Add(24 * time.Hour),
			wantVars: []any{time.Duration(24 * time.Hour).Microseconds()},
			want:     "DATE_ADD(`created_at`, INTERVAL ? MICROSECOND)",
		},
		{
			name:     "Sub use DATE_SUB(date, INTERVAL value unit)",
			expr:     NewTime("", "created_at").Sub(24 * time.Hour),
			wantVars: []any{time.Duration(24 * time.Hour).Microseconds()},
			want:     "DATE_SUB(`created_at`, INTERVAL ? MICROSECOND)",
		},
		{
			name:     "find_in_set",
			expr:     NewTime("", "created_at").FindInSet("1,2,3"),
			wantVars: []any{"1,2,3"},
			want:     "FIND_IN_SET(`created_at`, ?)",
		},
		{
			name:     "UNIX_TIMESTAMP use UNIX_TIMESTAMP(date)",
			expr:     NewTime("", "created_at").UnixTimestamp(),
			wantVars: nil,
			want:     "UNIX_TIMESTAMP(`created_at`)",
		},
		{
			name:     "Date use DATE(date)",
			expr:     NewTime("", "created_at").Date(),
			wantVars: nil,
			want:     "DATE(`created_at`)",
		},
		{
			name:     "Year use YEAR(date)",
			expr:     NewTime("", "created_at").Year(),
			wantVars: nil,
			want:     "YEAR(`created_at`)",
		},
		{
			name:     "Month use MONTH(date)",
			expr:     NewTime("", "created_at").Month(),
			wantVars: nil,
			want:     "MONTH(`created_at`)",
		},
		{
			name:     "Day use DAY(date)",
			expr:     NewTime("", "created_at").Day(),
			wantVars: nil,
			want:     "DAY(`created_at`)",
		},
		{
			name:     "Hour use HOUR(date)",
			expr:     NewTime("", "created_at").Hour(),
			wantVars: nil,
			want:     "HOUR(`created_at`)",
		},
		{
			name:     "Minute use MINUTE(date)",
			expr:     NewTime("", "created_at").Minute(),
			wantVars: nil,
			want:     "MINUTE(`created_at`)",
		},
		{
			name:     "Second use SECOND(date)",
			expr:     NewTime("", "created_at").Second(),
			wantVars: nil,
			want:     "SECOND(`created_at`)",
		},
		{
			name:     "Second use SECOND(date)",
			expr:     NewTime("", "created_at").Second(),
			wantVars: nil,
			want:     "SECOND(`created_at`)",
		},
		{
			name:     "MicroSecond use MICROSECOND(date)",
			expr:     NewTime("", "created_at").MicroSecond(),
			wantVars: nil,
			want:     "MICROSECOND(`created_at`)",
		},
		{
			name:     "DayOfWeek use DAYOFWEEK(date)",
			expr:     NewTime("", "created_at").DayOfWeek(),
			wantVars: nil,
			want:     "DAYOFWEEK(`created_at`)",
		},
		{
			name:     "DayOfMonth use DAYOFMONTH(date)",
			expr:     NewTime("", "created_at").DayOfMonth(),
			wantVars: nil,
			want:     "DAYOFMONTH(`created_at`)",
		},
		{
			name:     "DayOfYear use DAYOFYEAR(date)",
			expr:     NewTime("", "created_at").DayOfYear(),
			wantVars: nil,
			want:     "DAYOFYEAR(`created_at`)",
		},
		{
			name:     "Date use DATEDIFF(self, value)",
			expr:     NewTime("", "created_at").DateDiff(value1),
			wantVars: []any{value1},
			want:     "DATEDIFF(`created_at`,?)",
		},
		{
			name:     "DateFormat use DATE_FORMAT(date,format)",
			expr:     NewTime("", "created_at").DateFormat("%W %M %Y"),
			wantVars: []any{"%W %M %Y"},
			want:     "DATE_FORMAT(`created_at`,?)",
		},
		{
			name:     "DayName use DAYNAME(date)",
			expr:     NewTime("", "created_at").DayName(),
			wantVars: nil,
			want:     "DAYNAME(`created_at`)",
		},
		{
			name:     "MonthName use MONTHNAME(date)",
			expr:     NewTime("", "created_at").MonthName(),
			wantVars: nil,
			want:     "MONTHNAME(`created_at`)",
		},
		{
			name: "add",
			expr: NewTime("", "created_at").AddCol(NewTime("", "created_at1")),
			want: "`created_at` + `created_at1`",
		},
		{
			name: "add with table",
			expr: NewTime("user", "created_at").AddCol(NewTime("userB", "created_at1")),
			want: "`user`.`created_at` + `userB`.`created_at1`",
		},
		{
			name: "sub",
			expr: NewTime("", "created_at").SubCol(NewTime("", "created_at1")),
			want: "`created_at` - `created_at1`",
		},
		{
			name: "sub with table",
			expr: NewTime("user", "created_at").SubCol(NewTime("userB", "created_at1")),
			want: "`user`.`created_at` - `userB`.`created_at1`",
		},
		{
			name: "mul",
			expr: NewTime("", "created_at").MulCol(NewTime("", "created_at1")),
			want: "(`created_at`) * (`created_at1`)",
		},
		{
			name: "mul with table",
			expr: NewTime("user", "created_at").MulCol(NewTime("userB", "created_at1")),
			want: "(`user`.`created_at`) * (`userB`.`created_at1`)",
		},
		{
			name: "mul",
			expr: NewTime("", "created_at").DivCol(NewTime("", "created_at1")),
			want: "(`created_at`) / (`created_at1`)",
		},
		{
			name: "mul with table",
			expr: NewTime("user", "created_at").DivCol(NewTime("userB", "created_at1")),
			want: "(`user`.`created_at`) / (`userB`.`created_at1`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_SetExpr_Time(t *testing.T) {
	var zeroValue time.Time
	value1, _ := time.Parse("2006-01-02 15:04:05", "2021-06-29 15:11:49")

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     NewTime("user", "created_at").Value(value1),
			wantVars: []any{value1},
			want:     "`created_at` = ?",
		},
		{
			name:     "Value",
			expr:     NewTime("user", "created_at").ValueZero(),
			wantVars: []any{zeroValue},
			want:     "`created_at` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
			})
		})
	}
}
