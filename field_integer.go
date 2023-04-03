package assist

type Int = Integer[int]
type Int8 = Integer[int8]
type Int16 = Integer[int16]
type Int32 = Integer[int32]
type Int64 = Integer[int64]
type Uint = Integer[uint]
type Uint8 = Integer[uint8]
type Uint16 = Integer[uint16]
type Uint32 = Integer[uint32]
type Uint64 = Integer[uint64]

// NewField new int field
func NewInt(table, column string, opts ...Option) Int {
	return NewInteger[int](table, column, opts...)
}

// NewField new int8 field
func NewInt8(table, column string, opts ...Option) Int8 {
	return NewInteger[int8](table, column, opts...)
}

// NewField new int16 field
func NewInt16(table, column string, opts ...Option) Int16 {
	return NewInteger[int16](table, column, opts...)
}

// NewField new int32 field
func NewInt32(table, column string, opts ...Option) Int32 {
	return NewInteger[int32](table, column, opts...)
}

// NewField new int64 field
func NewInt64(table, column string, opts ...Option) Int64 {
	return NewInteger[int64](table, column, opts...)
}

// NewField new uint field
func NewUint(table, column string, opts ...Option) Uint {
	return NewInteger[uint](table, column, opts...)
}

// NewField new uint8 field
func NewUint8(table, column string, opts ...Option) Uint8 {
	return NewInteger[uint8](table, column, opts...)
}

// NewField new uint16 field
func NewUint16(table, column string, opts ...Option) Uint16 {
	return NewInteger[uint16](table, column, opts...)
}

// NewField new uint32 field
func NewUint32(table, column string, opts ...Option) Uint32 {
	return NewInteger[uint32](table, column, opts...)
}

// NewField new uint64 field
func NewUint64(table, column string, opts ...Option) Uint64 {
	return NewInteger[uint64](table, column, opts...)
}
