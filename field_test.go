package rapier

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"
	"time"
)

var _ sql.Scanner = (*password)(nil)
var _ driver.Valuer = (*password)(nil)

type password string

func (p *password) Scan(src any) error {
	*p = password(fmt.Sprintf("this is password {%q}", src))
	return nil
}
func (p password) Value() (driver.Value, error) {
	return strings.TrimPrefix(strings.TrimSuffix(string(p), "}"), "this is password {"), nil
}

func Test_Field_Expr_Keyword(t *testing.T) {
	var value1 = 1
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "as",
			expr:     NewField("", "id").As("user_id"),
			wantVars: nil,
			want:     "`id` AS `user_id`",
		},
		{
			name:     "as expression",
			expr:     NewField("", "id").Add(value1).As("user_id"),
			wantVars: []any{value1},
			want:     "`id`+? AS `user_id`",
		},
		{
			name:     "as with table",
			expr:     NewField("user", "id").As("user_id"),
			wantVars: nil,
			want:     "`user`.`id` AS `user_id`",
		},
		{
			name:     "as expression with table",
			expr:     NewField("user", "id").Add(value1).As("user_id"),
			wantVars: []any{value1},
			want:     "`user`.`id`+? AS `user_id`",
		},
		{
			name:     "desc",
			expr:     NewField("", "id").Desc(),
			wantVars: nil,
			want:     "`id` DESC",
		},
		{
			name:     "desc with table",
			expr:     NewField("user", "id").Desc(),
			wantVars: nil,
			want:     "`user`.`id` DESC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_Field_Expr_Basic(t *testing.T) {
	var p = password("i am password")

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "value support sql.Scanner and driver.Valuer",
			expr:     NewField("user", "password").Eq(p),
			wantVars: []any{p},
			want:     "`user`.`password` = ?",
		},
		{
			name: "is null",
			expr: NewField("", "id").IsNull(),
			want: "`id` IS NULL",
		},
		{
			name: "is not null",
			expr: NewField("", "id").IsNotNull(),
			want: "`id` IS NOT NULL",
		},
		{
			name: "count",
			expr: NewField("", "id").Count(),
			want: "COUNT(`id`)",
		},
		{
			name: "count",
			expr: NewField("", "id").Distinct(),
			want: "DISTINCT `id`",
		},
		{
			name: "length",
			expr: NewField("", "id").Length(),
			want: "LENGTH(`id`)",
		},
		{
			name: "max",
			expr: NewField("", "id").Max(),
			want: "MAX(`id`)",
		},
		{
			name: "min",
			expr: NewField("", "id").Min(),
			want: "MIN(`id`)",
		},
		{
			name: "avg",
			expr: NewField("", "id").Avg(),
			want: "AVG(`id`)",
		},
		{
			name: "GROUP_CONCAT",
			expr: NewField("", "id").GroupConcat(),
			want: "GROUP_CONCAT(`id`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_Field_Expr_Col(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name: "equal",
			expr: NewField("", "id").EqCol(NewField("", "new_id")),
			want: "`id` = `new_id`",
		},
		{
			name: "equal with table",
			expr: NewField("user", "id").EqCol(NewField("userB", "new_id")),
			want: "`user`.`id` = `userB`.`new_id`",
		},
		{
			name: "equal, field with table function",
			expr: NewField("", "id").EqCol(NewField("", "new_id").WithTable("userB")),
			want: "`id` = `userB`.`new_id`",
		},
		{
			name: "equal | AVG",
			expr: NewField("", "id").EqCol(NewField("", "new_id").Avg()),
			want: "`id` = AVG(`new_id`)",
		},
		{
			name: "equal | AVG with table",
			expr: NewField("user", "id").EqCol(NewField("userB", "new_id").Avg()),
			want: "`user`.`id` = AVG(`userB`.`new_id`)",
		},
		{
			name: "not equal",
			expr: NewField("", "id").NeqCol(NewField("", "new_id")),
			want: "`id` <> `new_id`",
		},
		{
			name: "not equal with table",
			expr: NewField("user", "id").NeqCol(NewField("userB", "new_id")),
			want: "`user`.`id` <> `userB`.`new_id`",
		},
		{
			name: "less than",
			expr: NewField("", "id").LtCol(NewField("", "new_id")),
			want: "`id` < `new_id`",
		},
		{
			name: "less than with table",
			expr: NewField("user", "id").LtCol(NewField("userB", "new_id")),
			want: "`user`.`id` < `userB`.`new_id`",
		},
		{
			name: "less than or equal",
			expr: NewField("", "id").LteCol(NewField("", "new_id")),
			want: "`id` <= `new_id`",
		},
		{
			name: "less than or equal with table",
			expr: NewField("user", "id").LteCol(NewField("userB", "new_id")),
			want: "`user`.`id` <= `userB`.`new_id`",
		},
		{
			name: "greater than",
			expr: NewField("", "id").GtCol(NewField("", "new_id")),
			want: "`id` > `new_id`",
		},
		{
			name: "greater than with table",
			expr: NewField("user", "id").GtCol(NewField("userB", "new_id")),
			want: "`user`.`id` > `userB`.`new_id`",
		},
		{
			name: "greater than or equal",
			expr: NewField("", "id").GteCol(NewField("", "new_id")),
			want: "`id` >= `new_id`",
		},
		{
			name: "greater than or equal with table",
			expr: NewField("user", "id").GteCol(NewField("userB", "new_id")),
			want: "`user`.`id` >= `userB`.`new_id`",
		},

		{
			name:     "find_in_set",
			expr:     NewField("", "address").FindInSetCol(NewField("userB", "new_id")),
			wantVars: nil,
			want:     "FIND_IN_SET(`address`, `userB`.`new_id`)",
		},
		{
			name:     "find_in_set with",
			expr:     NewField("", "address").FindInSetColWith(NewField("userB", "new_id")),
			wantVars: nil,
			want:     "FIND_IN_SET(`userB`.`new_id`, `address`)",
		},

		{
			name: "add",
			expr: NewField("", "id").AddCol(NewField("", "new_id")),
			want: "`id` + `new_id`",
		},
		{
			name: "add with table",
			expr: NewField("user", "id").AddCol(NewField("userB", "new_id")),
			want: "`user`.`id` + `userB`.`new_id`",
		},
		{
			name: "sub",
			expr: NewField("", "id").SubCol(NewField("", "new_id")),
			want: "`id` - `new_id`",
		},
		{
			name: "sub with table",
			expr: NewField("user", "id").SubCol(NewField("userB", "new_id")),
			want: "`user`.`id` - `userB`.`new_id`",
		},
		{
			name: "mul",
			expr: NewField("", "id").MulCol(NewField("", "new_id")),
			want: "(`id`) * (`new_id`)",
		},
		{
			name: "mul with table",
			expr: NewField("user", "id").MulCol(NewField("userB", "new_id")),
			want: "(`user`.`id`) * (`userB`.`new_id`)",
		},
		{
			name: "mul",
			expr: NewField("", "id").DivCol(NewField("", "new_id")),
			want: "(`id`) / (`new_id`)",
		},
		{
			name: "mul with table",
			expr: NewField("user", "id").DivCol(NewField("userB", "new_id")),
			want: "(`user`.`id`) / (`userB`.`new_id`)",
		},
		{
			name: "concat",
			expr: NewField("", "id").ConcatCol(NewField("", "new_id"), NewField("", "new_id2")),
			want: "Concat(`id`,`new_id`,`new_id2`)",
		},
		{
			name:     "concat with raw",
			expr:     NewField("", "id").ConcatCol(NewField("", "new_id"), NewRaw("'/'")),
			wantVars: nil,
			want:     "Concat(`id`,`new_id`,'/')",
		},
		{
			name: "concat with table",
			expr: NewField("user", "id").ConcatCol(NewField("userB", "new_id"), NewField("userC", "new_id2")),
			want: "Concat(`user`.`id`,`userB`.`new_id`,`userC`.`new_id2`)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_Expr_Field(t *testing.T) {
	var value1 int = 1
	var value2 int = 2
	var value3 int = 3
	timeTime1, _ := time.Parse("2006-01-02 15:04:05", "2021-06-29 15:11:49")
	timeDuration1 := time.Second
	timeDurationToMicroSecond1 := time.Duration(timeDuration1).Microseconds()

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewField("t1", "age").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`t1`.`age`,?)",
		},
		{
			name:     "eq",
			expr:     NewField("t1", "age").Eq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` = ?",
		},
		{
			name:     "neq",
			expr:     NewField("t1", "age").Neq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` <> ?",
		},
		{
			name:     "gt",
			expr:     NewField("t1", "age").Gt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` > ?",
		},
		{
			name:     "gte",
			expr:     NewField("t1", "age").Gte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` >= ?",
		},
		{
			name:     "lt",
			expr:     NewField("t1", "age").Lt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` < ?",
		},
		{
			name:     "lte",
			expr:     NewField("t1", "age").Lte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` <= ?",
		},
		{
			name:     "between",
			expr:     NewField("t1", "age").Between(value1, value2),
			wantVars: []any{value1, value2},
			want:     "`t1`.`age` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     NewField("t1", "age").NotBetween(value1, value2),
			wantVars: []any{value1, value2},
			want:     "NOT (`t1`.`age` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     NewField("t1", "age").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` IN (?,?,?)",
		},
		{
			name:     "not in",
			expr:     NewField("t1", "age").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` NOT IN (?,?,?)",
		},
		{
			name:     "like",
			expr:     NewField("t1", "age").Like(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` LIKE ?",
		},
		{
			name:     "not like",
			expr:     NewField("t1", "age").NotLike(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` NOT LIKE ?",
		},
		{
			name:     "Sum",
			expr:     NewField("t1", "age").Sum(),
			wantVars: nil,
			want:     "SUM(`t1`.`age`)",
		},

		{
			name:     "add",
			expr:     NewField("t1", "age").Add(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`+?",
		},
		{
			name:     "add time.Duration",
			expr:     NewField("t1", "age").Add(timeDuration1),
			wantVars: []any{timeDurationToMicroSecond1},
			want:     "DATE_ADD(`t1`.`age`, INTERVAL ? MICROSECOND)",
		},
		{
			name:     "sub",
			expr:     NewField("t1", "age").Sub(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`-?",
		},
		{
			name:     "sub time.Duration",
			expr:     NewField("t1", "age").Sub(timeDuration1),
			wantVars: []any{timeDurationToMicroSecond1},
			want:     "DATE_SUB(`t1`.`age`, INTERVAL ? MICROSECOND)",
		},
		{
			name:     "mul",
			expr:     NewField("t1", "age").Mul(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`*?",
		},
		{
			name:     "mul expression",
			expr:     NewField("t1", "age").Add(value1).Mul(value1),
			wantVars: []any{value1, value1},
			want:     "(`t1`.`age`+?)*?",
		},
		{
			name:     "div",
			expr:     NewField("t1", "age").Div(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`/?",
		},
		{
			name:     "div expression",
			expr:     NewField("t1", "age").Add(value1).Div(value1),
			wantVars: []any{value1, value1},
			want:     "(`t1`.`age`+?)/?",
		},
		{
			name:     "mod",
			expr:     NewField("t1", "age").Mod(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`%?",
		},
		{
			name:     "mul expression",
			expr:     NewField("t1", "age").Add(value1).Mod(value1),
			wantVars: []any{value1, value1},
			want:     "(`t1`.`age`+?)%?",
		},
		{
			name:     "floor div",
			expr:     NewField("t1", "age").FloorDiv(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` DIV ?",
		},
		{
			name:     "floor div expression",
			expr:     NewField("t1", "age").Add(value1).FloorDiv(value1),
			wantVars: []any{value1, value1},
			want:     "(`t1`.`age`+?) DIV ?",
		},
		{
			name:     "floor",
			expr:     NewField("t1", "age").Floor(),
			wantVars: nil,
			want:     "FLOOR(`t1`.`age`)",
		},
		{
			name:     "round",
			expr:     NewField("t1", "age").Round(2),
			wantVars: []any{2},
			want:     "ROUND(`t1`.`age`, ?)",
		},
		{
			name:     "complex use +,-,*,/",
			expr:     NewField("t1", "age").Add(value1).Mul(value2).Div(value3),
			wantVars: []any{value1, value2, value3},
			want:     "((`t1`.`age`+?)*?)/?",
		},
		{
			name:     "right shift",
			expr:     NewField("t1", "age").RightShift(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`>>?",
		},
		{
			name:     "right shift expression",
			expr:     NewField("t1", "age").Add(value1).RightShift(value3),
			wantVars: []any{value1, value3},
			want:     "(`t1`.`age`+?)>>?",
		},
		{
			name:     "left shift",
			expr:     NewField("t1", "age").LeftShift(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`<<?",
		},
		{
			name:     "left shift expression",
			expr:     NewField("t1", "age").Add(value1).LeftShift(value3),
			wantVars: []any{value1, value3},
			want:     "(`t1`.`age`+?)<<?",
		},
		{
			name:     "bit xor",
			expr:     NewField("t1", "age").BitXor(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`^?",
		},
		{
			name:     "bit xor expression",
			expr:     NewField("t1", "age").Add(value1).BitXor(value3),
			wantVars: []any{value1, value3},
			want:     "(`t1`.`age`+?)^?",
		},
		{
			name:     "bit and",
			expr:     NewField("t1", "age").BitAnd(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`&?",
		},
		{
			name:     "bit and expression",
			expr:     NewField("t1", "age").Add(value1).BitAnd(value3),
			wantVars: []any{value1, value3},
			want:     "(`t1`.`age`+?)&?",
		},
		{
			name:     "bit or",
			expr:     NewField("t1", "age").BitOr(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`|?",
		},
		{
			name:     "bit or expression",
			expr:     NewField("t1", "age").Add(value1).BitOr(value3),
			wantVars: []any{value1, value3},
			want:     "(`t1`.`age`+?)|?",
		},
		{
			name:     "bit flip",
			expr:     NewField("t1", "age").BitFlip(),
			wantVars: nil,
			want:     "~`t1`.`age`",
		},
		{
			name:     "bit flip expression",
			expr:     NewField("t1", "age").Add(value1).BitFlip(),
			wantVars: []any{value1},
			want:     "~(`t1`.`age`+?)",
		},
		{
			name:     "regexp",
			expr:     NewField("", "name").Regexp(".*"),
			wantVars: []any{".*"},
			want:     "`name` REGEXP ?",
		},
		{
			name:     "not regexp",
			expr:     NewField("", "name").NotRegxp(".*"),
			wantVars: []any{".*"},
			want:     "NOT `name` REGEXP ?",
		},
		{
			name:     "find_in_set",
			expr:     NewField("", "address").FindInSet("1,2,3"),
			wantVars: []any{"1,2,3"},
			want:     "FIND_IN_SET(`address`, ?)",
		},
		{
			name:     "find_in_set with",
			expr:     NewField("", "address").FindInSetWith("a"),
			wantVars: []any{"a"},
			want:     "FIND_IN_SET(?, `address`)",
		},
		{
			name:     "SUBSTRING_INDEX",
			expr:     NewField("", "address").SubstringIndex(",", 2),
			wantVars: nil,
			want:     "SUBSTRING_INDEX(`address`,\",\",2)",
		},
		{
			name:     "replace",
			expr:     NewField("", "address").Replace("address", "path"),
			wantVars: []any{"address", "path"},
			want:     "REPLACE(`address`,?,?)",
		},
		{
			name:     "concat with '',''",
			expr:     NewField("", "address").Concat("", ""),
			wantVars: nil,
			want:     "`address`",
		},
		{
			name:     "concat with '[',']'",
			expr:     NewField("", "address").Concat("[", "]"),
			wantVars: []any{"[", "]"},
			want:     "CONCAT(?,`address`,?)",
		},
		{
			name:     "concat with '','_'",
			expr:     NewField("", "address").Concat("", "_"),
			wantVars: []any{"_"},
			want:     "CONCAT(`address`,?)",
		},
		{
			name:     "concat with '_',''",
			expr:     NewField("", "address").Concat("_", ""),
			wantVars: []any{"_"},
			want:     "CONCAT(?,`address`)",
		},
		{
			name:     "replace then concat with '[',']'",
			expr:     NewField("", "address").Replace("address", "path").Concat("[", "]"),
			wantVars: []any{"[", "address", "path", "]"},
			want:     "CONCAT(?,REPLACE(`address`,?,?),?)",
		},
		{
			name:     "trim",
			expr:     NewField("", "address").Trim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(BOTH ? FROM `address`)",
		},
		{
			name:     "leading trim",
			expr:     NewField("", "address").LTrim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(LEADING ? FROM `address`)",
		},
		{
			name:     "trailing trim",
			expr:     NewField("", "address").RTrim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(TRAILING ? FROM `address`)",
		},
		{
			name:     "trim space",
			expr:     NewField("", "address").TrimSpace(),
			wantVars: nil,
			want:     "TRIM(`address`)",
		},
		{
			name:     "leading trim space",
			expr:     NewField("", "address").LTrimSpace(),
			wantVars: nil,
			want:     "LTRIM(`address`)",
		},
		{
			name:     "trailing trim space",
			expr:     NewField("", "address").RTrimSpace(),
			wantVars: nil,
			want:     "RTRIM(`address`)",
		},

		{
			name:     "FromUnixTime use FROM_UNIXTIME(date)",
			expr:     NewField("t1", "age").FromUnixTime(),
			wantVars: nil,
			want:     "FROM_UNIXTIME(`t1`.`age`)",
		},
		{
			name:     "FromUnixTime use FROM_UNIXTIME(date,format)",
			expr:     NewField("t1", "age").FromUnixTime("%Y%m%d"),
			wantVars: []any{"%Y%m%d"},
			want:     "FROM_UNIXTIME(`t1`.`age`, ?)",
		},
		{
			name:     "FROM_DAYS",
			expr:     NewField("t1", "age").FromDays(),
			wantVars: nil,
			want:     "FROM_DAYS(`t1`.`age`)",
		},
		{
			name:     "UNIX_TIMESTAMP use UNIX_TIMESTAMP(date)",
			expr:     NewField("", "created_at").UnixTimestamp(),
			wantVars: nil,
			want:     "UNIX_TIMESTAMP(`created_at`)",
		},
		{
			name:     "Date use DATE(date)",
			expr:     NewField("", "created_at").Date(),
			wantVars: nil,
			want:     "DATE(`created_at`)",
		},
		{
			name:     "Year use YEAR(date)",
			expr:     NewField("", "created_at").Year(),
			wantVars: nil,
			want:     "YEAR(`created_at`)",
		},
		{
			name:     "Month use MONTH(date)",
			expr:     NewField("", "created_at").Month(),
			wantVars: nil,
			want:     "MONTH(`created_at`)",
		},
		{
			name:     "Day use DAY(date)",
			expr:     NewField("", "created_at").Day(),
			wantVars: nil,
			want:     "DAY(`created_at`)",
		},
		{
			name:     "Hour use HOUR(date)",
			expr:     NewField("", "created_at").Hour(),
			wantVars: nil,
			want:     "HOUR(`created_at`)",
		},
		{
			name:     "Minute use MINUTE(date)",
			expr:     NewField("", "created_at").Minute(),
			wantVars: nil,
			want:     "MINUTE(`created_at`)",
		},
		{
			name:     "Second use SECOND(date)",
			expr:     NewField("", "created_at").Second(),
			wantVars: nil,
			want:     "SECOND(`created_at`)",
		},
		{
			name:     "Second use SECOND(date)",
			expr:     NewField("", "created_at").Second(),
			wantVars: nil,
			want:     "SECOND(`created_at`)",
		},
		{
			name:     "MicroSecond use MICROSECOND(date)",
			expr:     NewField("", "created_at").MicroSecond(),
			wantVars: nil,
			want:     "MICROSECOND(`created_at`)",
		},
		{
			name:     "DayOfWeek use DAYOFWEEK(date)",
			expr:     NewField("", "created_at").DayOfWeek(),
			wantVars: nil,
			want:     "DAYOFWEEK(`created_at`)",
		},
		{
			name:     "DayOfMonth use DAYOFMONTH(date)",
			expr:     NewField("", "created_at").DayOfMonth(),
			wantVars: nil,
			want:     "DAYOFMONTH(`created_at`)",
		},
		{
			name:     "DayOfYear use DAYOFYEAR(date)",
			expr:     NewField("", "created_at").DayOfYear(),
			wantVars: nil,
			want:     "DAYOFYEAR(`created_at`)",
		},
		{
			name:     "Date use DATEDIFF(self, value)",
			expr:     NewField("", "created_at").DateDiff(timeTime1),
			wantVars: []any{timeTime1},
			want:     "DATEDIFF(`created_at`,?)",
		},
		{
			name:     "DateFormat use DATE_FORMAT(date,format)",
			expr:     NewField("", "created_at").DateFormat("%W %M %Y"),
			wantVars: []any{"%W %M %Y"},
			want:     "DATE_FORMAT(`created_at`,?)",
		},
		{
			name:     "DayName use DAYNAME(date)",
			expr:     NewField("", "created_at").DayName(),
			wantVars: nil,
			want:     "DAYNAME(`created_at`)",
		},
		{
			name:     "MonthName use MONTHNAME(date)",
			expr:     NewField("", "created_at").MonthName(),
			wantVars: nil,
			want:     "MONTHNAME(`created_at`)",
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

func Test_SetExpr_Field(t *testing.T) {
	var nullTime *time.Time

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     NewField("user", "address").Value("abc"),
			wantVars: []any{"abc"},
			want:     "`address`=?",
		},
		{
			name:     "Value: any null",
			expr:     NewField("user", "address").Value(nil),
			wantVars: []any{nil},
			want:     "`address`=?",
		},
		{
			name:     "Value: null",
			expr:     NewField("user", "address").ValueNull(),
			wantVars: []any{nil},
			want:     "`address`=?",
		},
		{
			name:     "Value: any",
			expr:     NewField("user", "sex").ValueAny(1),
			wantVars: []any{1},
			want:     "`sex`=?",
		},
		{
			name:     "Value: any null",
			expr:     NewField("user", "sex").ValueAny(nil),
			wantVars: []any{nil},
			want:     "`sex`=?",
		},
		{
			name:     "Value: any null value",
			expr:     NewField("user", "sex").ValueAny(nullTime),
			wantVars: []any{nullTime},
			want:     "`sex`=?",
		},
		{
			name: "SetSubQuery",
			expr: NewField("user", "address").SetSubQuery(
				newDb().Table("`user`").Select("`address`").Where("`id` = ?", 100),
			),
			wantVars: []any{100},
			want:     "`address`=(SELECT `address` FROM `user` WHERE `id` = ?)",
		},
		// Col
		{
			name:     "SetCol",
			expr:     NewField("user", "id").SetCol(NewField("user", "new_id")),
			wantVars: nil,
			want:     "`id`=`user`.`new_id`",
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
