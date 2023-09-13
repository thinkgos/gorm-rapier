package rapier

import "testing"

func Test_Bool(t *testing.T) {
	var value1 bool = true
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewBool("", "male").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`male`,?)",
		},
		{
			name:     "eq",
			expr:     NewBool("", "male").Eq(value1),
			wantVars: []any{value1},
			want:     "`male` = ?",
		},
		{
			name:     "eq",
			expr:     NewBool("", "male").Neq(value1),
			wantVars: []any{value1},
			want:     "`male` <> ?",
		},
		{
			name:     "is",
			expr:     NewBool("", "male").Is(value1),
			wantVars: []any{value1},
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
			expr:     NewBool("", "male").Xor(value1),
			wantVars: []any{value1},
			want:     "`male` XOR ?",
		},
		{
			name:     "and",
			expr:     NewBool("", "male").And(value1),
			wantVars: []any{value1},
			want:     "`male` AND ?",
		},
		{
			name:     "or",
			expr:     NewBool("", "male").Or(value1),
			wantVars: []any{value1},
			want:     "`male` OR ?",
		},

		{
			name:     "bit and",
			expr:     NewBool("", "male").BitAnd(value1),
			wantVars: []any{value1},
			want:     "`male`&?",
		},
		{
			name:     "bit or",
			expr:     NewBool("", "male").BitOr(value1),
			wantVars: []any{value1},
			want:     "`male`|?",
		},
		{
			name:     "bit xor",
			expr:     NewBool("", "male").BitXor(value1),
			wantVars: []any{value1},
			want:     "`male`^?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
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
				CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
			})
		})
	}
}
