package assist

import "testing"

func Test_Bytes(t *testing.T) {
	value1 := []byte("tom")
	value2 := []byte("lucy")
	value3 := []byte("john")
	value4 := [][]byte{value1, value2, value3}
	value5 := []TestBytes{value1, value2, value3}
	value6 := []int{1, 2, 3}

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewBytes("", "name").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`name`,?)",
		},
		{
			name:     "eq",
			expr:     NewBytes("", "name").Eq(value1),
			wantVars: []any{value1},
			want:     "`name` = ?",
		},
		{
			name:     "neq",
			expr:     NewBytes("", "name").Neq(value1),
			wantVars: []any{value1},
			want:     "`name` <> ?",
		},
		{
			name:     "gt",
			expr:     NewBytes("", "name").Gt(value1),
			wantVars: []any{value1},
			want:     "`name` > ?",
		},
		{
			name:     "gte",
			expr:     NewBytes("", "name").Gte(value1),
			wantVars: []any{value1},
			want:     "`name` >= ?",
		},
		{
			name:     "lt",
			expr:     NewBytes("", "name").Lt(value1),
			wantVars: []any{value1},
			want:     "`name` < ?",
		},
		{
			name:     "lte",
			expr:     NewBytes("", "name").Lte(value1),
			wantVars: []any{value1},
			want:     "`name` <= ?",
		},
		{
			name:     "between",
			expr:     NewBytes("", "name").Between(value1, value2),
			wantVars: []any{value1, value2},
			want:     "`name` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     NewBytes("", "name").NotBetween(value1, value2),
			wantVars: []any{value1, value2},
			want:     "NOT (`name` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     NewBytes("", "name").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`name` IN (?,?,?)",
		},

		{
			name:     "in any current type",
			expr:     NewString("", "name").InAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "in any under new type",
			expr:     NewString("", "name").InAny(value5),
			wantVars: []any{TestBytes(value1), TestBytes(value2), TestBytes(value3)},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "in any under type string",
			expr:     NewString("", "name").InAny(value6),
			wantVars: []any{1, 2, 3},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "in any but not a array/slice",
			expr:     NewString("", "name").InAny(1),
			wantVars: nil,
			want:     "",
		},

		{
			name:     "not in",
			expr:     NewBytes("", "name").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`name` NOT IN (?,?,?)",
		},
		{
			name:     "not in any current type",
			expr:     NewString("", "name").NotInAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`name` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under new type",
			expr:     NewString("", "name").NotInAny(value5),
			wantVars: []any{TestBytes(value1), TestBytes(value2), TestBytes(value3)},
			want:     "`name` NOT IN (?,?,?)",
		},
		{
			name:     "not in any under type string",
			expr:     NewString("", "name").NotInAny(value6),
			wantVars: []any{1, 2, 3},
			want:     "`name` NOT IN (?,?,?)",
		},
		{
			name:     "not in any but not a array/slice",
			expr:     NewString("", "name").NotInAny(1),
			wantVars: nil,
			want:     "NOT",
		},
		{
			name:     "like",
			expr:     NewBytes("", "name").Like("%%tom%%"),
			wantVars: []any{"%%tom%%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "fuzzy like",
			expr:     NewBytes("", "name").FuzzyLike("tom"),
			wantVars: []any{"%tom%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "left like",
			expr:     NewBytes("", "name").LeftLike("tom"),
			wantVars: []any{"tom%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "not like",
			expr:     NewBytes("", "name").NotLike("%%tom%%"),
			wantVars: []any{"%%tom%%"},
			want:     "`name` NOT LIKE ?",
		},
		{
			name:     "regexp",
			expr:     NewBytes("", "name").Regexp(".*"),
			wantVars: []any{".*"},
			want:     "`name` REGEXP ?",
		},
		{
			name:     "not regexp",
			expr:     NewBytes("", "name").NotRegxp(".*"),
			wantVars: []any{".*"},
			want:     "NOT `name` REGEXP ?",
		},
		{
			name:     "find_in_set",
			expr:     NewBytes("", "address").FindInSet("a"),
			wantVars: []any{"a"},
			want:     "FIND_IN_SET(`address`, ?)",
		},
		{
			name:     "find_in_set with",
			expr:     NewBytes("", "address").FindInSetWith("sh"),
			wantVars: []any{"sh"},
			want:     "FIND_IN_SET(?, `address`)",
		},
		{
			name:     "SUBSTRING_INDEX",
			expr:     NewBytes("", "address").SubstringIndex(",", 2),
			wantVars: nil,
			want:     "SUBSTRING_INDEX(`address`,\",\",2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
