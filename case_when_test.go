package rapier

import "testing"

func Test_Expr_CaseWhen(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "case when - one",
			expr:     CaseWhen(WhenThen(NewField("", "id1").Gt(100), NewField("", "value1"))),
			wantVars: []any{100},
			want:     "(CASE WHEN `id1` > ? THEN `value1` END)",
		},
		{
			name: "case when - multiple",
			expr: CaseWhen(
				WhenThen(
					NewField("", "id1").Gt(100),
					NewField("", "value1"),
				),
				WhenThen(
					NewField("", "id2").Gt(200),
					NewField("", "value2"),
				),
			),
			wantVars: []any{100, 200},
			want:     "(CASE WHEN `id1` > ? THEN `value1` WHEN `id2` > ? THEN `value2` END)",
		},
		{
			name: "case when - multiple with AND",
			expr: CaseWhen(
				WhenThen(
					And(NewField("", "id1").Gt(100), NewField("", "id1").Lt(1000)),
					NewField("", "value1"),
				),
				WhenThen(
					NewField("", "id2").Gt(200),
					NewField("", "value2"),
				),
			),
			wantVars: []any{100, 1000, 200},
			want:     "(CASE WHEN (`id1` > ? AND `id1` < ?) THEN `value1` WHEN `id2` > ? THEN `value2` END)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_Expr_CaseWhenElse(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name: "case when else - one",
			expr: CaseWhenElse(
				NewField("", "result"),
				WhenThen(NewField("", "id1").Gt(100), NewField("", "value1")),
			),
			wantVars: []any{100},
			want:     "(CASE WHEN `id1` > ? THEN `value1` ELSE `result` END)",
		},
		{
			name: "case when else - multiple",
			expr: CaseWhenElse(
				NewField("", "result"),
				WhenThen(NewField("", "id1").Gt(100),
					NewField("", "value1"),
				),
				WhenThen(NewField("", "id2").Gt(200), NewField("", "value2")),
			),
			wantVars: []any{100, 200},
			want:     "(CASE WHEN `id1` > ? THEN `value1` WHEN `id2` > ? THEN `value2` ELSE `result` END)",
		},
		{
			name: "case when else - multiple with AND",
			expr: CaseWhenElse(
				NewField("", "result"),
				WhenThen(
					And(NewField("", "id1").Gt(100), NewField("", "id1").Lt(1000)),
					NewField("", "value1"),
				),
				WhenThen(
					NewField("", "id2").Gt(200),
					NewField("", "value2"),
				),
			),
			wantVars: []any{100, 1000, 200},
			want:     "(CASE WHEN (`id1` > ? AND `id1` < ?) THEN `value1` WHEN `id2` > ? THEN `value2` ELSE `result` END)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
