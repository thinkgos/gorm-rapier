package assist

import "testing"

func Test_Expr_String(t *testing.T) {
	var value1 = ""
	var value2 = "lucy"
	var value3 = "john"

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IfNull",
			expr:     NewString("", "name").IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`name`,?)",
		},
		{
			name:     "eq",
			expr:     NewString("", "name").Eq(value1),
			wantVars: []any{value1},
			want:     "`name` = ?",
		},
		{
			name:     "neq",
			expr:     NewString("", "name").Neq(value1),
			wantVars: []any{value1},
			want:     "`name` <> ?",
		},
		{
			name:     "gt",
			expr:     NewString("", "name").Gt(value1),
			wantVars: []any{value1},
			want:     "`name` > ?",
		},
		{
			name:     "gte",
			expr:     NewString("", "name").Gte(value1),
			wantVars: []any{value1},
			want:     "`name` >= ?",
		},
		{
			name:     "lt",
			expr:     NewString("", "name").Lt(value1),
			wantVars: []any{value1},
			want:     "`name` < ?",
		},
		{
			name:     "lte",
			expr:     NewString("", "name").Lte(value1),
			wantVars: []any{value1},
			want:     "`name` <= ?",
		},
		{
			name:     "between",
			expr:     NewString("", "name").Between(value1, value2),
			wantVars: []any{value1, value2},
			want:     "`name` BETWEEN ? AND ?",
		},
		{
			name:     "not between",
			expr:     NewString("", "name").NotBetween(value1, value2),
			wantVars: []any{value1, value2},
			want:     "NOT (`name` BETWEEN ? AND ?)",
		},
		{
			name:     "in",
			expr:     NewString("", "name").In(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "not in",
			expr:     NewString("", "name").NotIn(value1, value2, value3),
			wantVars: []any{value1, value2, value3},
			want:     "`name` NOT IN (?,?,?)",
		},
		{
			name:     "like",
			expr:     NewString("", "name").Like("%%tom%%"),
			wantVars: []any{"%%tom%%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "not like",
			expr:     NewString("", "name").NotLike("%%tom%%"),
			wantVars: []any{"%%tom%%"},
			want:     "`name` NOT LIKE ?",
		},
		{
			name:     "regexp",
			expr:     NewString("", "name").Regexp(".*"),
			wantVars: []any{".*"},
			want:     "`name` REGEXP ?",
		},
		{
			name:     "not regexp",
			expr:     NewString("", "name").NotRegxp(".*"),
			wantVars: []any{".*"},
			want:     "NOT `name` REGEXP ?",
		},
		{
			name:     "find_in_set",
			expr:     NewString("", "address").FindInSet("sh"),
			wantVars: []any{"sh"},
			want:     "FIND_IN_SET(`address`,?)",
		},
		{
			name:     "find_in_set with",
			expr:     NewString("", "address").FindInSetWith("sh"),
			wantVars: []any{"sh"},
			want:     "FIND_IN_SET(?,`address`)",
		},
		{
			name:     "SUBSTRING_INDEX",
			expr:     NewString("", "address").SubstringIndex(",", 2),
			wantVars: nil,
			want:     "SUBSTRING_INDEX(`address`,\",\",2)",
		},
		{
			name:     "replace",
			expr:     NewString("", "address").Replace("address", "path"),
			wantVars: []any{"address", "path"},
			want:     "REPLACE(`address`,?,?)",
		},
		{
			name:     "concat with '',''",
			expr:     NewString("", "address").Concat("", ""),
			wantVars: nil,
			want:     "`address`",
		},
		{
			name:     "concat with '[',']'",
			expr:     NewString("", "address").Concat("[", "]"),
			wantVars: []any{"[", "]"},
			want:     "CONCAT(?,`address`,?)",
		},
		{
			name:     "concat with '','_'",
			expr:     NewString("", "address").Concat("", "_"),
			wantVars: []any{"_"},
			want:     "CONCAT(`address`,?)",
		},
		{
			name:     "concat with '_',''",
			expr:     NewString("", "address").Concat("_", ""),
			wantVars: []any{"_"},
			want:     "CONCAT(?,`address`)",
		},
		{
			name:     "replace then concat with '[',']'",
			expr:     NewString("", "address").Replace("address", "path").Concat("[", "]"),
			wantVars: []any{"[", "address", "path", "]"},
			want:     "CONCAT(?,REPLACE(`address`,?,?),?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
