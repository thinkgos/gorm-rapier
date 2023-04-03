package assist

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

func testExprFloat[T constraints.Float | ~string](
	t *testing.T,
	newFloat func(table, column string, opts ...Option) Float[T],
	getTestValue func() (T, T, T),
) {
	value1, value2, value3 := getTestValue()

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
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
			name:     "not in",
			expr:     newFloat("t1", "score").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`t1`.`score` NOT IN (?,?,?)",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
