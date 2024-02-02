package rapier

import "testing"

func Test_Expr_IF(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IF with field",
			expr:     IF(NewField("", "id1").Gt(100), NewField("", "t"), NewField("", "f")),
			wantVars: []any{100},
			want:     "IF(`id1` > ?,`t`,`f`)",
		},
		{
			name:     "IF with raw value",
			expr:     IF(NewField("", "id1").Gt(100), NewRaw("t"), NewRaw("f")),
			wantVars: []any{100},
			want:     "IF(`id1` > ?,t,f)",
		},
		{
			name:     "IF with raw value",
			expr:     IF(NewField("", "id1").Gt(100), NewRaw("t"), NewField("", "f").Sub(1)),
			wantVars: []any{100, 1},
			want:     "IF(`id1` > ?,t,`f`-?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
