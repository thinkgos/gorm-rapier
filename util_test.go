package assist

import (
	"reflect"
	"testing"
	"time"
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
		query, args := buildSelectValue(db.Statement, xDict.Id, xDict.Name.Length().As("name"))
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
		query := buildColumnsValue(db, xDict.Id, xDict.Name)
		if want := "`dict`.`id`,`dict`.`name`"; query != want {
			t.Errorf("SQL expects %v got %v", want, query)
		}
	})
}

func Test_buildAssignSet(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		db := newDb()
		query := buildAssignSet(
			db,
			[]SetExpr{
				xDict.Pid.Value(100),
				xDict.Score.Add(1),
			})
		t.Logf("%+v", query)
	})
}

func Test_buildColumnName(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		got := buildColumnName(
			xDict.Pid,
			xDict.Score,
		)
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

type TestInteger int32
type TestFloat float64
type TestString string
type TestBytes []byte
type TestTime time.Time

func Test_IntoIntegerSlice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		want := []int{1, 2, 3}
		got := IntoIntegerSlice[TestInteger, int]([]TestInteger{1, 2, 3})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Vars expects %+v got %v", want, got)
		}
	})
}
