package rapier

import (
	"reflect"
	"testing"

	"gorm.io/gorm/clause"
)

func Test_buildSelectValue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		db := newDb()
		query, args := buildSelectValue(db.Statement)
		if query != "" {
			t.Errorf("SQL expects %v got %v", "", query)
		}
		if len(args) != 0 {
			t.Errorf("Vars expects %+v got %v", nil, args)
		}
	})
	t.Run("fields", func(t *testing.T) {
		db := newDb()
		query, args := buildSelectValue(db.Statement, refDict.Id, refDict.Name.Length().As("name"))
		if want := "`dict`.`id`"; query != want {
			t.Errorf("SQL expects %v got %v", want, query)
		}
		if wantArg1 := any("LENGTH(`dict`.`name`) AS `name`"); len(args) != 1 && args[0] != wantArg1 {
			t.Errorf("Vars expects %+v got %v", []any{wantArg1}, args)
		}
	})
}

func Test_buildColumnsValue(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		db := newDb()
		query := buildColumnsValue(db)
		if query != "" {
			t.Errorf("SQL expects %v got %v", "", query)
		}
	})
	t.Run("fields", func(t *testing.T) {
		db := newDb()
		query := buildColumnsValue(db, refDict.Id, refDict.Name)
		if want := "`dict`.`id`,`dict`.`name`"; query != want {
			t.Errorf("SQL expects %v got %v", want, query)
		}
	})
}

func Test_buildAssignSet(t *testing.T) {
	db := newDb()
	got := buildClauseSet(
		db,
		[]AssignExpr{
			refDict.Pid.Value(100),
			refDict.Score.Add(1),
			refDict.Name.valueEq("name"),
		})
	want := clause.Set{
		{
			Column: clause.Column{Name: "pid"},
			Value:  int64(100),
		},
		{
			Column: clause.Column{Name: "score"},
			Value: clause.Expr{
				SQL: "?+?", Vars: []any{
					clause.Column{Table: "dict", Name: "score"},
					float64(1),
				},
			},
		},
		{
			Column: clause.Column{Name: "name"},
			Value:  "name",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildAssignSet want: %v got: %v", want, got)
	}
}

func Test_buildAttrsValue(t *testing.T) {
	want := []any{
		clause.Eq{
			Column: clause.Column{
				Name: refDict.Pid.ColumnName(),
			},
			Value: int64(100),
		},
		clause.Eq{
			Column: refDict.Name.ColumnName(),
			Value:  "name",
		},
	}
	got := buildAttrsValue(
		[]AssignExpr{
			refDict.Pid.Value(100),
			refDict.Name.valueEq("name"),
		})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildAttrSet want: %v got: %v", want, got)
	}
}

func Test_buildColumnName(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		got := buildColumnName(refDict.Pid, refDict.Score)
		if want := []string{"pid", "score"}; !reflect.DeepEqual(got, want) {
			t.Errorf("column name expects %+v got %v", want, got)
		}
	})
}

func Test_IntoSlice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		want := []int{1, 2, 3}
		got := IntoSlice([]uint8{1, 2, 3}, func(v uint8) int {
			return int(v)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Vars expects %+v got %v", want, got)
		}
	})
}

func Test_IntoIntegerSlice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		want := []int{1, 2, 3}
		got := IntoIntegerSlice[TestInteger, int]([]TestInteger{1, 2, 3})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Vars expects %+v got %v", want, got)
		}
	})
}
