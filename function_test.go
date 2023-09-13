package rapier

import "testing"

func Test_Expr_Function(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name: "or",
			expr: Or(NewField("", "id1"), NewField("", "id2")),
			want: "(`id1` OR `id2`)",
		},
		{
			name: "and",
			expr: And(NewField("", "id1"), NewField("", "id2")),
			want: "(`id1` AND `id2`)",
		},
		{
			name: "not",
			expr: Not(NewField("", "id1"), NewField("", "id2")),
			want: "(NOT `id1` AND NOT `id2`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
