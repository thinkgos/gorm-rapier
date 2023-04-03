package assist

type Float32 = Float[float32]
type Float64 = Float[float64]
type Decimal = Float[string]

// NewFloat32 new float32 field
func NewFloat32(table, column string, opts ...Option) Float32 {
	return NewFloat[float32](table, column, opts...)
}

// NewFloat64 new float64 field
func NewFloat64(table, column string, opts ...Option) Float64 {
	return NewFloat[float64](table, column, opts...)
}

// NewDecimal new decimal field
func NewDecimal(table, column string, opts ...Option) Decimal {
	return NewFloat[string](table, column, opts...)
}
