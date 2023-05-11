package assist

import (
	"reflect"
	"testing"
)

func Test_buildSelectValue(t *testing.T) {
	t.Run("", func(t *testing.T) {
		db := newDb()
		query, args := buildSelectValue(db.Statement)
		if query != "" {
			t.Errorf("SQL expects %v got %v", "", query)
		}
		if len(args) != 0 {
			t.Errorf("Vars expects %+v got %v", nil, args)
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

func Test_IntoIntegerSlice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		want := []int{1, 2, 3}
		got := IntoIntegerSlice[TestInteger, int]([]TestInteger{1, 2, 3})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Vars expects %+v got %v", want, got)
		}
	})
}
