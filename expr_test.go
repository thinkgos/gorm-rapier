package rapier

import (
	"testing"
)

func Test_Expr(t *testing.T) {
	t.Run("column name", func(t *testing.T) {
		got := NewField("table", "id").ColumnName()
		if want := "id"; got != want {
			t.Errorf("ColumnName want %+v, got %v", want, got)
		}
	})
	t.Run("field name", func(t *testing.T) {
		got := NewField("table", "id").FieldName()
		if want := "id"; got != want {
			t.Errorf("FieldName want %+v, got %v", want, got)
		}
		got = NewField("table", "id").FieldName("tb")
		if want := "tb_id"; got != want {
			t.Errorf("FieldName want %+v, got %v", want, got)
		}
	})
}

func Test_Expr_BuildColumn(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		buildOpt []BuildOption
		want     string
	}{
		{
			name:     "BuildOpt - empty",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{},
			want:     "`id`",
		},
		{
			name:     "BuildOpt - WithTable",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{WithTable},
			want:     "`user`.`id`",
		},
		{
			name:     "BuildOpt - WithoutQuote",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{WithoutQuote},
			want:     "id",
		},
		{
			name:     "BuildOpt - WithAll",
			expr:     NewUint("user", "id").As("user_id"),
			buildOpt: []BuildOption{WithAll},
			want:     "`user`.`id` AS `user_id`",
		},
		{
			name:     "BuildOpt - WithAll(not alias)",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{WithAll},
			want:     "`user`.`id`",
		},
		{
			name:     "BuildOpt - WithTable + WithoutQuote",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{WithTable, WithoutQuote},
			want:     "user.id",
		},
		{
			name:     "BuildOpt - WithAll + WithoutQuote(not alias)",
			expr:     NewUint("user", "id"),
			buildOpt: []BuildOption{WithAll, WithoutQuote},
			want:     "user.id",
		},
		{
			name:     "BuildOpt - WithAll + WithoutQuote",
			expr:     NewUint("user", "id").As("user_id"),
			buildOpt: []BuildOption{WithAll, WithoutQuote},
			want:     "user.id AS user_id",
		},
		{
			name:     "BuildOpt - WithoutQuote use withAppendBuildOpts",
			expr:     NewUint("user", "id").withAppendBuildOpts(WithoutQuote),
			buildOpt: nil,
			want:     "id",
		},
		// star(*): all columns
		{
			name:     "Star: BuildOpt - empty",
			expr:     NewAsterisk(""),
			buildOpt: []BuildOption{},
			want:     "*",
		},
		{
			name:     "Star: BuildOpt - WithTable",
			expr:     NewAsterisk("user"),
			buildOpt: []BuildOption{WithTable},
			want:     "`user`.*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := NewStatement()
			got := tt.expr.BuildColumn(stmt, tt.buildOpt...)
			if tt.want != got {
				t.Errorf("BuildColumn() got: %q, except: %q", got, tt.want)
			}
		})
	}
}

func BenchmarkExpr_Count(b *testing.B) {
	id := NewUint("", "id")
	for i := 0; i < b.N; i++ {
		n := id.Count()
		_ = n
	}
}
