package rapier

import "testing"

func Test_Expr_String(t *testing.T) {
	value1 := ""
	value2 := "lucy"
	value3 := "john"
	value4 := []string{value1, value2, value3}
	value5 := []TestString{TestString(value1), TestString(value2), TestString(value3)}
	value6 := []int{1, 2, 3}

	value4Ptr := &value4

	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "IntoField",
			expr:     NewString("", "name").IntoField().IfNull(value1),
			wantVars: []any{value1},
			want:     "IFNULL(`name`,?)",
		},
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
			name:     "in any current type",
			expr:     NewString("", "name").InAny(value4),
			wantVars: []any{value1, value2, value3},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "in any current type with ptr prt two",
			expr:     NewString("", "name").InAny(&value4Ptr),
			wantVars: []any{value1, value2, value3},
			want:     "`name` IN (?,?,?)",
		},
		{
			name:     "in any under new type",
			expr:     NewString("", "name").InAny(value5),
			wantVars: []any{TestString(value1), TestString(value2), TestString(value3)},
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
			expr:     NewString("", "name").NotIn(value1, value2, value3),
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
			wantVars: []any{TestString(value1), TestString(value2), TestString(value3)},
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
			expr:     NewString("", "name").Like("%%tom%%"),
			wantVars: []any{"%%tom%%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "fuzzy like",
			expr:     NewString("", "name").FuzzyLike("tom"),
			wantVars: []any{"%tom%"},
			want:     "`name` LIKE ?",
		},
		{
			name:     "left like",
			expr:     NewString("", "name").LeftLike("tom"),
			wantVars: []any{"tom%"},
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
			expr:     NewString("", "address").FindInSet("a"),
			wantVars: []any{"a"},
			want:     "FIND_IN_SET(`address`, ?)",
		},
		{
			name:     "find_in_set with",
			expr:     NewString("", "address").FindInSetWith("sh"),
			wantVars: []any{"sh"},
			want:     "FIND_IN_SET(?, `address`)",
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
		{
			name:     "hidden with star",
			expr:     NewString("", "address").Hidden(2, 3, "******"),
			wantVars: []any{2, "******", 3},
			want:     "CONCAT(LEFT(`address`,?),?,RIGHT(`address`,?))",
		},
		{
			name:     "hidden suffix with star",
			expr:     NewString("", "address").HiddenSuffix(2, "******"),
			wantVars: []any{2, "******"},
			want:     "CONCAT(LEFT(`address`,?),?)",
		},
		{
			name:     "hidden prefix with star",
			expr:     NewString("", "address").HiddenPrefix(3, "******"),
			wantVars: []any{"******", 3},
			want:     "CONCAT(?,RIGHT(`address`,?))",
		},
		{
			name:     "trim",
			expr:     NewString("", "address").Trim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(BOTH ? FROM `address`)",
		},
		{
			name:     "leading trim",
			expr:     NewString("", "address").LTrim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(LEADING ? FROM `address`)",
		},
		{
			name:     "trailing trim",
			expr:     NewString("", "address").RTrim("abc"),
			wantVars: []any{"abc"},
			want:     "TRIM(TRAILING ? FROM `address`)",
		},

		{
			name:     "trim space",
			expr:     NewString("", "address").TrimSpace(),
			wantVars: nil,
			want:     "TRIM(`address`)",
		},
		{
			name:     "leading trim space",
			expr:     NewString("", "address").LTrimSpace(),
			wantVars: nil,
			want:     "LTRIM(`address`)",
		},
		{
			name:     "trailing trim space",
			expr:     NewString("", "address").RTrimSpace(),
			wantVars: nil,
			want:     "RTRIM(`address`)",
		},
		{
			name: "add",
			expr: NewString("", "id").AddCol(NewString("", "new_id")),
			want: "`id` + `new_id`",
		},
		{
			name: "add with table",
			expr: NewString("user", "id").AddCol(NewString("userB", "new_id")),
			want: "`user`.`id` + `userB`.`new_id`",
		},
		{
			name: "sub",
			expr: NewString("", "id").SubCol(NewString("", "new_id")),
			want: "`id` - `new_id`",
		},
		{
			name: "sub with table",
			expr: NewString("user", "id").SubCol(NewString("userB", "new_id")),
			want: "`user`.`id` - `userB`.`new_id`",
		},
		{
			name: "mul",
			expr: NewString("", "id").MulCol(NewString("", "new_id")),
			want: "(`id`) * (`new_id`)",
		},
		{
			name: "mul with table",
			expr: NewString("user", "id").MulCol(NewString("userB", "new_id")),
			want: "(`user`.`id`) * (`userB`.`new_id`)",
		},
		{
			name: "mul",
			expr: NewString("", "id").DivCol(NewString("", "new_id")),
			want: "(`id`) / (`new_id`)",
		},
		{
			name: "mul with table",
			expr: NewString("user", "id").DivCol(NewString("userB", "new_id")),
			want: "(`user`.`id`) / (`userB`.`new_id`)",
		},
		{
			name: "concat",
			expr: NewString("", "id").ConcatCol(NewString("", "new_id"), NewString("", "new_id2")),
			want: "Concat(`id`,`new_id`,`new_id2`)",
		},
		{
			name:     "concat with raw",
			expr:     NewString("", "id").ConcatCol(NewString("", "new_id"), NewRaw("'/'")),
			wantVars: nil,
			want:     "Concat(`id`,`new_id`,'/')",
		},
		{
			name: "concat with table",
			expr: NewString("user", "id").ConcatCol(NewString("userB", "new_id"), NewString("userC", "new_id2")),
			want: "Concat(`user`.`id`,`userB`.`new_id`,`userC`.`new_id2`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}

func Test_SetExpr_String(t *testing.T) {
	value1 := "abc"
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name:     "Value",
			expr:     NewString("user", "address").Value(value1),
			wantVars: []any{value1},
			want:     "`address`=?",
		},
		{
			name:     "ValuePointer: null",
			expr:     NewString("user", "address").ValuePointer(nil),
			wantVars: []any{(*string)(nil)},
			want:     "`address`=?",
		},
		{
			name:     "ValuePointer: pointer",
			expr:     NewString("user", "address").ValuePointer(&value1),
			wantVars: []any{&value1},
			want:     "`address`=?",
		},
		{
			name:     "Value",
			expr:     NewString("user", "address").ValueZero(),
			wantVars: []any{""},
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
