package assist

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func Test_Expr_Integer(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		testExprInteger(t, NewInteger[int])
	})
	t.Run("int8", func(t *testing.T) {
		testExprInteger(t, NewInteger[int8])
	})
	t.Run("int16", func(t *testing.T) {
		testExprInteger(t, NewInteger[int16])
	})
	t.Run("int32", func(t *testing.T) {
		testExprInteger(t, NewInteger[int32])
	})
	t.Run("int64", func(t *testing.T) {
		testExprInteger(t, NewInteger[int64])
	})
	t.Run("uint", func(t *testing.T) {
		testExprInteger(t, NewInteger[uint])
	})
	t.Run("uint8", func(t *testing.T) {
		testExprInteger(t, NewInteger[uint8])
	})
	t.Run("uint16", func(t *testing.T) {
		testExprInteger(t, NewInteger[uint16])
	})
	t.Run("uint32", func(t *testing.T) {
		testExprInteger(t, NewInteger[uint32])
	})
	t.Run("uint64", func(t *testing.T) {
		testExprInteger(t, NewInteger[uint64])
	})
}

func Test_SetExpr_Integer(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[int])
	})
	t.Run("int8", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[int8])
	})
	t.Run("int16", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[int16])
	})
	t.Run("int32", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[int32])
	})
	t.Run("int64", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[int64])
	})
	t.Run("uint", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[uint])
	})
	t.Run("uint8", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[uint8])
	})
	t.Run("uint16", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[uint16])
	})
	t.Run("uint32", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[uint32])
	})
	t.Run("uint64", func(t *testing.T) {
		testSetExprInteger(t, NewInteger[uint64])
	})
}

func testExprInteger[T constraints.Integer](
	t *testing.T,
	newInteger func(table, column string, opts ...Option) Integer[T],
) {
	var value1 T = 0
	var value2 T = 5
	var value3 T = 8
	var value4 []T = []T{value1, value2, value3}
	var value5 []TestInteger = []TestInteger{1, 2, 3}
	var value6 []string = []string{"1", "2", "3"}

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     newInteger("t1", "age").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`t1`.`age`,?)",
		},
		{
			name:     "eq",
			expr:     newInteger("t1", "age").Eq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` = ?",
		},
		{
			name:     "neq",
			expr:     newInteger("t1", "age").Neq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` <> ?",
		},
		{
			name:     "gt",
			expr:     newInteger("t1", "age").Gt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` > ?",
		},
		{
			name:     "gte",
			expr:     newInteger("t1", "age").Gte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` >= ?",
		},
		{
			name:     "lt",
			expr:     newInteger("t1", "age").Lt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` < ?",
		},
		{
			name:     "lte",
			expr:     newInteger("t1", "age").Lte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` <= ?",
		},
		{
			name:     "between",
			expr:     newInteger("t1", "age").Between(value1, value2),
			wantVars: []any{value1, value2},
			want:     "`t1`.`age` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     newInteger("t1", "age").NotBetween(value1, value2),
			wantVars: []any{value1, value2},
			want:     "NOT (`t1`.`age` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     newInteger("t1", "age").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` IN (?,?,?)",
		},
		{
			name:     "in any current type",
			expr:     newInteger("t1", "age").InAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` IN (?,?,?)",
		},
		{
			name:     "in any under new type",
			expr:     newInteger("t1", "age").InAny(value5),
			wantVars: []any{TestInteger(1), TestInteger(2), TestInteger(3)},
			want:     "`t1`.`age` IN (?,?,?)",
		},
		{
			name:     "in any under type string",
			expr:     newInteger("t1", "age").InAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`t1`.`age` IN (?,?,?)",
		},
		{
			name:     "in any but not a array/slice",
			expr:     newInteger("t1", "age").InAny(1),
			wantVars: nil,
			want:     "",
		},
		{
			name:     "not in",
			expr:     newInteger("t1", "age").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` NOT IN (?,?,?)",
		},
		{
			name:     "not in any current type",
			expr:     newInteger("t1", "age").NotInAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`age` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under new type",
			expr:     newInteger("t1", "age").NotInAny(value5),
			wantVars: []any{TestInteger(1), TestInteger(2), TestInteger(3)},
			want:     "`t1`.`age` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under type string",
			expr:     newInteger("t1", "age").NotInAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`t1`.`age` NOT IN (?,?,?)",
		},
		{
			name:     "not in any but not a array/slice",
			expr:     newInteger("t1", "age").NotInAny(1),
			wantVars: nil,
			want:     "NOT",
		},
		{
			name:     "like",
			expr:     newInteger("t1", "age").Like(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` LIKE ?",
		},
		{
			name:     "not like",
			expr:     newInteger("t1", "age").NotLike(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` NOT LIKE ?",
		},
		{
			name:     "find_in_set",
			expr:     newInteger("t1", "age").FindInSet("1,2,3"),
			wantVars: []any{"1,2,3"},
			want:     "FIND_IN_SET(`t1`.`age`, ?)",
		},
		{
			name:     "Sum",
			expr:     newInteger("t1", "age").Sum(),
			wantVars: nil,
			want:     "SUM(`t1`.`age`)",
		},
		{
			name:     "add",
			expr:     newInteger("t1", "age").Add(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`+?",
		},
		{
			name:     "sub",
			expr:     newInteger("t1", "age").Sub(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`-?",
		},
		{
			name:     "mul",
			expr:     newInteger("t1", "age").Mul(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`*?",
		},
		{
			name:     "div",
			expr:     newInteger("t1", "age").Div(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`/?",
		},
		{
			name:     "mod",
			expr:     newInteger("t1", "age").Mod(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age`%?",
		},
		{
			name:     "floor div",
			expr:     newInteger("t1", "age").FloorDiv(value1),
			wantVars: []any{value1},
			want:     "`t1`.`age` DIV ?",
		},
		{
			name:     "round",
			expr:     newInteger("t1", "age").Round(2),
			wantVars: []any{2},
			want:     "ROUND(`t1`.`age`, ?)",
		},
		{
			name:     "complex use +,-,*,/",
			expr:     newInteger("t1", "age").Add(value1).Mul(value2).Div(value3),
			wantVars: []any{value1, value2, value3},
			want:     "((`t1`.`age`+?)*?)/?",
		},
		{
			name:     "right shift",
			expr:     newInteger("t1", "age").RightShift(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`>>?",
		},
		{
			name:     "left shift",
			expr:     newInteger("t1", "age").LeftShift(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`<<?",
		},
		{
			name:     "bit xor",
			expr:     newInteger("t1", "age").BitXor(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`^?",
		},
		{
			name:     "bit and",
			expr:     newInteger("t1", "age").BitAnd(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`&?",
		},
		{
			name:     "bit or",
			expr:     newInteger("t1", "age").BitOr(value3),
			wantVars: []any{value3},
			want:     "`t1`.`age`|?",
		},
		{
			name:     "bit flip",
			expr:     newInteger("t1", "age").BitFlip(),
			wantVars: nil,
			want:     "~`t1`.`age`",
		},
		{
			name:     "FromUnixTime use FROM_UNIXTIME(date)",
			expr:     newInteger("t1", "age").FromUnixTime(),
			wantVars: nil,
			want:     "FROM_UNIXTIME(`t1`.`age`)",
		},
		{
			name:     "FromUnixTime use FROM_UNIXTIME(date,format)",
			expr:     newInteger("t1", "age").FromUnixTime("%Y%m%d"),
			wantVars: []any{"%Y%m%d"},
			want:     "FROM_UNIXTIME(`t1`.`age`, ?)",
		},
		{
			name:     "FROM_DAYS",
			expr:     newInteger("t1", "age").FromDays(),
			wantVars: nil,
			want:     "FROM_DAYS(`t1`.`age`)",
		},
		{
			name: "add",
			expr: newInteger("", "id").AddCol(newInteger("", "new_id")),
			want: "`id` + `new_id`",
		},
		{
			name: "add with table",
			expr: newInteger("user", "id").AddCol(newInteger("userB", "new_id")),
			want: "`user`.`id` + `userB`.`new_id`",
		},
		{
			name: "sub",
			expr: newInteger("", "id").SubCol(newInteger("", "new_id")),
			want: "`id` - `new_id`",
		},
		{
			name: "sub with table",
			expr: newInteger("user", "id").SubCol(newInteger("userB", "new_id")),
			want: "`user`.`id` - `userB`.`new_id`",
		},
		{
			name: "mul",
			expr: newInteger("", "id").MulCol(newInteger("", "new_id")),
			want: "(`id`) * (`new_id`)",
		},
		{
			name: "mul with table",
			expr: newInteger("user", "id").MulCol(newInteger("userB", "new_id")),
			want: "(`user`.`id`) * (`userB`.`new_id`)",
		},
		{
			name: "mul",
			expr: newInteger("", "id").DivCol(newInteger("", "new_id")),
			want: "(`id`) / (`new_id`)",
		},
		{
			name: "mul with table",
			expr: newInteger("user", "id").DivCol(newInteger("userB", "new_id")),
			want: "(`user`.`id`) / (`userB`.`new_id`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func testSetExprInteger[T constraints.Integer](
	t *testing.T,
	newInteger func(table, column string, opts ...Option) Integer[T],
) {
	var zeroValue T
	var value T = 5

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     newInteger("user", "address").Value(value),
			wantVars: []any{value},
			want:     "`address` = ?",
		},
		{
			name:     "Value",
			expr:     newInteger("user", "address").ValueZero(),
			wantVars: []any{zeroValue},
			want:     "`address` = ?",
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
