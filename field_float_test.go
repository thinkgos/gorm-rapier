package assist

import "testing"

func Test_Expr_Float_For_Field(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		testExprFloat(
			t,
			NewFloat32,
			func() (float32, float32, float32) {
				return 0, 2.0, 3.0
			},
		)
	})
	t.Run("float64", func(t *testing.T) {
		testExprFloat(
			t,
			NewFloat64,
			func() (float64, float64, float64) {
				return 0, 2.0, 3.0
			},
		)
	})
	t.Run("decimal", func(t *testing.T) {
		testExprFloat(
			t,
			NewDecimal,
			func() (string, string, string) {
				return "0", "2.0", "3.0"
			},
		)
	})
}
