package assist

import (
	"reflect"
	"testing"
)

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
