package rapier

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func Test_Expr_Float(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		testExprFloat(
			t,
			NewFloat[float32],
			func() (float32, float32, float32) {
				return 1.0, 2.0, 3.0
			},
		)
	})
	t.Run("float64", func(t *testing.T) {
		testExprFloat(
			t,
			NewFloat[float64],
			func() (float64, float64, float64) {
				return 1.0, 2.0, 3.0
			},
		)
	})
	t.Run("decimal", func(t *testing.T) {
		testExprFloat(
			t,
			NewFloat[string],
			func() (string, string, string) {
				return "1.0", "2.0", "3.0"
			},
		)
	})
}

func Test_SetExpr_Float(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		testSetExprFloat(
			t,
			NewFloat[float32],
			func() (float32, float32) {
				return 0, 2.0
			},
		)
	})
	t.Run("float64", func(t *testing.T) {
		testSetExprFloat(
			t,
			NewFloat[float64],
			func() (float64, float64) {
				return 0, 2.0
			},
		)
	})
	t.Run("decimal", func(t *testing.T) {
		testSetExprFloat(
			t,
			NewFloat[string],
			func() (string, string) {
				return "", "2.0"
			},
		)
	})
}

func testExprFloat[T constraints.Float | ~string](
	t *testing.T,
	newFloat func(table, column string, opts ...Option) Float[T],
	getTestValue func() (T, T, T),
) {
	value1, value2, value3 := getTestValue()
	var value4 []T = []T{value1, value2, value3}
	var value5 []TestFloat = []TestFloat{1.1, 2.2, 3.3}
	var value6 []string = []string{"1", "2", "3"}

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IntoField",
			expr:     newFloat("t1", "score").IntoField().IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`t1`.`score`,?)",
		},
		{
			name:     "IfNull",
			expr:     newFloat("t1", "score").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`t1`.`score`,?)",
		},
		{
			name:     "eq",
			expr:     newFloat("t1", "score").Eq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` = ?",
		},
		{
			name:     "neq",
			expr:     newFloat("t1", "score").Neq(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` <> ?",
		},
		{
			name:     "gt",
			expr:     newFloat("t1", "score").Gt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` > ?",
		},
		{
			name:     "gte",
			expr:     newFloat("t1", "score").Gte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` >= ?",
		},
		{
			name:     "lt",
			expr:     newFloat("t1", "score").Lt(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` < ?",
		},
		{
			name:     "lte",
			expr:     newFloat("t1", "score").Lte(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` <= ?",
		},
		{
			name:     "between",
			expr:     newFloat("t1", "score").Between(value1, value2),
			wantVars: []any{value1, value2},
			want:     "`t1`.`score` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     newFloat("t1", "score").NotBetween(value1, value2),
			wantVars: []any{value1, value2},
			want:     "NOT (`t1`.`score` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     newFloat("t1", "score").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`score` IN (?,?,?)",
		},
		{
			name:     "in any current type",
			expr:     newFloat("t1", "score").InAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`score` IN (?,?,?)",
		},
		{
			name:     "in any under new type",
			expr:     newFloat("t1", "score").InAny(value5),
			wantVars: []any{TestFloat(1.1), TestFloat(2.2), TestFloat(3.3)},
			want:     "`t1`.`score` IN (?,?,?)",
		},
		{
			name:     "in any under type string",
			expr:     newFloat("t1", "score").InAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`t1`.`score` IN (?,?,?)",
		},
		{
			name:     "in any but not a array/slice",
			expr:     newFloat("t1", "score").InAny(1),
			wantVars: nil,
			want:     "",
		},
		{
			name:     "not in",
			expr:     newFloat("t1", "score").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`score` NOT IN (?,?,?)",
		},

		{
			name:     "not in any current type",
			expr:     newFloat("t1", "score").NotInAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`score` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under new type",
			expr:     newFloat("t1", "score").NotInAny(value5),
			wantVars: []any{TestFloat(1.1), TestFloat(2.2), TestFloat(3.3)},
			want:     "`t1`.`score` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under type string",
			expr:     newFloat("t1", "score").NotInAny(value6),
			wantVars: []any{"1", "2", "3"},
			want:     "`t1`.`score` NOT IN (?,?,?)",
		},
		{
			name:     "not in any but not a array/slice",
			expr:     newFloat("t1", "score").NotInAny(1),
			wantVars: nil,
			want:     "NOT",
		},

		{
			name:     "like",
			expr:     newFloat("t1", "score").Like(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` LIKE ?",
		},
		{
			name:     "not like",
			expr:     newFloat("t1", "score").NotLike(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` NOT LIKE ?",
		},
		{
			name:     "find_in_set",
			expr:     newFloat("t1", "score").FindInSet("1,2,3"),
			wantVars: []any{"1,2,3"},
			want:     "FIND_IN_SET(`t1`.`score`, ?)",
		},
		{
			name:     "Sum",
			expr:     newFloat("t1", "score").Sum(),
			wantVars: nil,
			want:     "SUM(`t1`.`score`)",
		},
		{
			name:     "add",
			expr:     newFloat("t1", "score").Add(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score`+?",
		},
		{
			name:     "sub",
			expr:     newFloat("t1", "score").Sub(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score`-?",
		},
		{
			name:     "mul",
			expr:     newFloat("t1", "score").Mul(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score`*?",
		},
		{
			name:     "div",
			expr:     newFloat("t1", "score").Div(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score`/?",
		},
		{
			name:     "floor div",
			expr:     newFloat("t1", "score").FloorDiv(value1),
			wantVars: []any{value1},
			want:     "`t1`.`score` DIV ?",
		},
		{
			name:     "floor",
			expr:     newFloat("t1", "score").Floor(),
			wantVars: nil,
			want:     "FLOOR(`t1`.`score`)",
		},
		{
			name:     "round",
			expr:     newFloat("t1", "score").Round(2),
			wantVars: []any{2},
			want:     "ROUND(`t1`.`score`, ?)",
		},
		{
			name:     "complex use +,-,*,/",
			expr:     newFloat("t1", "score").Add(value1).Mul(value2).Div(value3),
			wantVars: []any{value1, value2, value3},
			want:     "((`t1`.`score`+?)*?)/?",
		},
		{
			name: "add",
			expr: newFloat("", "id").AddCol(newFloat("", "new_id")),
			want: "`id` + `new_id`",
		},
		{
			name: "add with table",
			expr: newFloat("user", "id").AddCol(newFloat("userB", "new_id")),
			want: "`user`.`id` + `userB`.`new_id`",
		},
		{
			name: "sub",
			expr: newFloat("", "id").SubCol(newFloat("", "new_id")),
			want: "`id` - `new_id`",
		},
		{
			name: "sub with table",
			expr: newFloat("user", "id").SubCol(newFloat("userB", "new_id")),
			want: "`user`.`id` - `userB`.`new_id`",
		},
		{
			name: "mul",
			expr: newFloat("", "id").MulCol(newFloat("", "new_id")),
			want: "(`id`) * (`new_id`)",
		},
		{
			name: "mul with table",
			expr: newFloat("user", "id").MulCol(newFloat("userB", "new_id")),
			want: "(`user`.`id`) * (`userB`.`new_id`)",
		},
		{
			name: "mul",
			expr: newFloat("", "id").DivCol(newFloat("", "new_id")),
			want: "(`id`) / (`new_id`)",
		},
		{
			name: "mul with table",
			expr: newFloat("user", "id").DivCol(newFloat("userB", "new_id")),
			want: "(`user`.`id`) / (`userB`.`new_id`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func testSetExprFloat[T constraints.Float | ~string](
	t *testing.T,
	newFloat func(table, column string, opts ...Option) Float[T],
	getTestValue func() (T, T),
) {
	zeroValue, value := getTestValue()

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     newFloat("user", "address").Value(value),
			wantVars: []any{value},
			want:     "`address`=?",
		},
		{
			name:     "ValuePointer: null",
			expr:     newFloat("user", "address").ValuePointer(nil),
			wantVars: []any{(*T)(nil)},
			want:     "`address`=?",
		},
		{
			name:     "ValuePointer: pointer",
			expr:     newFloat("user", "address").ValuePointer(&value),
			wantVars: []any{&value},
			want:     "`address`=?",
		},
		{
			name:     "Value",
			expr:     newFloat("user", "address").ValueZero(),
			wantVars: []any{zeroValue},
			want:     "`address`=?",
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
