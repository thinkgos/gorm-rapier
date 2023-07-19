package assist

import "testing"

func Test_Expr_Integer_For_Field(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		testExprInteger(t, NewInt)
	})
	t.Run("int8", func(t *testing.T) {
		testExprInteger(t, NewInt8)
	})
	t.Run("int16", func(t *testing.T) {
		testExprInteger(t, NewInt16)
	})
	t.Run("int32", func(t *testing.T) {
		testExprInteger(t, NewInt32)
	})
	t.Run("int64", func(t *testing.T) {
		testExprInteger(t, NewInt64)
	})
	t.Run("uint", func(t *testing.T) {
		testExprInteger(t, NewUint)
	})
	t.Run("uint8", func(t *testing.T) {
		testExprInteger(t, NewUint8)
	})
	t.Run("uint16", func(t *testing.T) {
		testExprInteger(t, NewUint16)
	})
	t.Run("uint32", func(t *testing.T) {
		testExprInteger(t, NewUint32)
	})
	t.Run("uint64", func(t *testing.T) {
		testExprInteger(t, NewUint64)
	})
}

func Test_AssignExpr_Integer_For_Field(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[int])
	})
	t.Run("int8", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[int8])
	})
	t.Run("int16", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[int16])
	})
	t.Run("int32", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[int32])
	})
	t.Run("int64", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[int64])
	})
	t.Run("uint", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[uint])
	})
	t.Run("uint8", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[uint8])
	})
	t.Run("uint16", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[uint16])
	})
	t.Run("uint32", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[uint32])
	})
	t.Run("uint64", func(t *testing.T) {
		testAssignExprInteger(t, NewInteger[uint64])
	})
}
