package rapier

import "testing"

func Test_Bool(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewBool("", "male").IfNull(true),
			wantVars: []any{true},
			want:     "IFNULL(`male`,?)",
		},
		{
			name:     "NullIf",
			expr:     NewBool("", "male").NullIf(true),
			wantVars: []any{true},
			want:     "NULLIF(`male`,?)",
		},
		{
			name:     "eq",
			expr:     NewBool("", "male").Eq(true),
			wantVars: []any{true},
			want:     "`male` = ?",
		},
		{
			name:     "eq",
			expr:     NewBool("", "male").Neq(true),
			wantVars: []any{true},
			want:     "`male` <> ?",
		},
		{
			name:     "is",
			expr:     NewBool("", "male").Is(true),
			wantVars: []any{true},
			want:     "`male` = ?",
		},
		{
			name:     "find_in_set",
			expr:     NewBool("", "male").FindInSet("1,2,3"),
			wantVars: []any{"1,2,3"},
			want:     "FIND_IN_SET(`male`, ?)",
		},
		{
			name:     "not",
			expr:     NewBool("", "male").Not(),
			wantVars: nil,
			want:     "NOT `male`",
		},
		{
			name:     "xor",
			expr:     NewBool("", "male").Xor(true),
			wantVars: []any{true},
			want:     "`male` XOR ?",
		},
		{
			name:     "and",
			expr:     NewBool("", "male").And(true),
			wantVars: []any{true},
			want:     "`male` AND ?",
		},
		{
			name:     "or",
			expr:     NewBool("", "male").Or(true),
			wantVars: []any{true},
			want:     "`male` OR ?",
		},

		{
			name:     "bit and",
			expr:     NewBool("", "male").BitAnd(true),
			wantVars: []any{true},
			want:     "`male`&?",
		},
		{
			name:     "bit or",
			expr:     NewBool("", "male").BitOr(true),
			wantVars: []any{true},
			want:     "`male`|?",
		},
		{
			name:     "bit xor",
			expr:     NewBool("", "male").BitXor(true),
			wantVars: []any{true},
			want:     "`male`^?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_SetExpr_Bool(t *testing.T) {
	value := true
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     NewBool("user", "male").Value(true),
			wantVars: []any{true},
			want:     "`male`=?",
		},
		{
			name:     "ValuePointer: null",
			expr:     NewBool("user", "male").ValuePointer(nil),
			wantVars: []any{(*bool)(nil)},
			want:     "`male`=?",
		},
		{
			name:     "ValuePointer: pointer",
			expr:     NewBool("user", "male").ValuePointer(&value),
			wantVars: []any{&value},
			want:     "`male`=?",
		},
		{
			name:     "Value",
			expr:     NewBool("user", "male").ValueZero(),
			wantVars: []any{false},
			want:     "`male`=?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				ReviewBuildExpr(t, tt.expr, tt.want, tt.wantVars)
			})
		})
	}
}
