package rapier

import "testing"

func Test_Expr_Concat(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name: "concat",
			expr: ConcatCol(NewString("", "id"), NewString("", "new_id"), NewString("", "new_id2")),
			want: "CONCAT(`id`,`new_id`,`new_id2`)",
		},
		{
			name:     "concat with raw",
			expr:     ConcatCol(NewString("", "id"), NewString("", "new_id"), NewRaw("'/'")),
			wantVars: nil,
			want:     "CONCAT(`id`,`new_id`,'/')",
		},
		{
			name: "concat with table",
			expr: ConcatCol(NewString("user", "id"), NewString("userB", "new_id"), NewString("userC", "new_id2")),
			want: "CONCAT(`user`.`id`,`userB`.`new_id`,`userC`.`new_id2`)",
		},
		{
			name: "concat_ws",
			expr: ConcatWsCol(NewRaw(`'-'`), NewString("", "id"), NewString("", "new_id"), NewString("", "new_id2")),
			want: "CONCAT_WS('-',`id`,`new_id`,`new_id2`)",
		},
		{
			name:     "concat_ws with raw",
			expr:     ConcatWsCol(NewRaw(`'-'`), NewString("", "id"), NewString("", "new_id"), NewRaw("'/'")),
			wantVars: nil,
			want:     "CONCAT_WS('-',`id`,`new_id`,'/')",
		},
		{
			name: "concat_ws with table",
			expr: ConcatWsCol(NewRaw(`'-'`), NewString("user", "id"), NewString("userB", "new_id"), NewString("userC", "new_id2")),
			want: "CONCAT_WS('-',`user`.`id`,`userB`.`new_id`,`userC`.`new_id2`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
