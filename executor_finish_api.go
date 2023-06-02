package assist

import "gorm.io/gorm"

func (x *Executor[T]) FirstOne() (*T, error) {
	var row T

	err := x.First(&row)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (x *Executor[T]) TakeOne() (*T, error) {
	var row T

	err := x.Take(&row)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (x *Executor[T]) LastOne() (*T, error) {
	var row T

	err := x.Last(&row)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (x *Executor[T]) ScanOne() (*T, error) {
	var row T

	err := x.Scan(&row)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (x *Executor[T]) Count() (count int64, err error) {
	err = x.IntoDB().Count(&count).Error
	return count, err
}

func (x *Executor[T]) First(dest any) error {
	return x.IntoDB().First(dest).Error
}

func (x *Executor[T]) FirstOrInit(dest any) error {
	return x.IntoDB().FirstOrInit(dest).Error
}

func (x *Executor[T]) FirstOrCreate(dest any) error {
	return x.IntoDB().FirstOrCreate(dest).Error
}

func (x *Executor[T]) Take(dest any) error {
	return x.IntoDB().Take(dest).Error
}

func (x *Executor[T]) Last(dest any) error {
	return x.IntoDB().Last(dest).Error
}

func (x *Executor[T]) Scan(dest any) error {
	return x.IntoDB().Scan(dest).Error
}

func (x *Executor[T]) Pluck(column string, value any) error {
	return x.IntoDB().Pluck(column, value).Error
}

func (x *Executor[T]) Exist() (exist bool, err error) {
	err = x.IntoDB().
		Select("1").
		Limit(1).
		Scan(&exist).Error
	return exist, err
}

func (x *Executor[T]) FindAll() ([]*T, error) {
	var rows []*T

	err := x.Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (x *Executor[T]) Find(dest any) error {
	return x.IntoDB().Find(dest).Error
}

func (x *Executor[T]) FindInBatches(dest any, batchSize int, fc func(tx *gorm.DB, batch int) error) error {
	return x.IntoDB().FindInBatches(dest, batchSize, fc).Error
}

func (x *Executor[T]) Create(value any) error {
	return x.db.Scopes(x.funcs...).Create(value).Error
}

func (x *Executor[T]) CreateInBatches(value any, batchSize int) error {
	return x.db.Scopes(x.funcs...).CreateInBatches(value, batchSize).Error
}

func (x *Executor[T]) Save(value any) error {
	return x.db.Scopes(x.funcs...).Save(value).Error
}

func (x *Executor[T]) Updates(value *T) (rowsAffected int64, err error) {
	result := x.IntoDB().Updates(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdatesMap(value map[string]any) (rowsAffected int64, err error) {
	result := x.IntoDB().Updates(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) Update(column string, value any) (rowsAffected int64, err error) {
	result := x.IntoDB().Update(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumns(value *T) (rowsAffected int64, err error) {
	result := x.IntoDB().UpdateColumns(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumnsMap(value map[string]any) (rowsAffected int64, err error) {
	result := x.IntoDB().UpdateColumns(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumn(column string, value any) (rowsAffected int64, err error) {
	result := x.IntoDB().UpdateColumn(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) Delete() (rowsAffected int64, err error) {
	var t T

	result := x.IntoDB().Delete(&t)
	return result.RowsAffected, result.Error
}

/**************************** 辅助 api *************************************/

func (x *Executor[T]) FirstBool() (v bool, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstInt() (v int, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstInt8() (v int8, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstInt16() (v int16, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstInt32() (v int32, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstInt64() (v int64, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstUint() (v uint, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstUint8() (v uint8, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstUint16() (v uint16, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstUint32() (v uint32, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstUint64() (v uint64, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstFloat32() (v float32, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstFloat64() (v float64, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstString() (v string, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) TakeBool() (v bool, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeInt() (v int, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeInt8() (v int8, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeInt16() (v int16, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeInt32() (v int32, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeInt64() (v int64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeUint() (v uint, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeUint8() (v uint8, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeUint16() (v uint16, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeUint32() (v uint32, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeUint64() (v uint64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeFloat32() (v float32, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeFloat64() (v float64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeString() (v string, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) ScanBool() (v bool, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanInt() (v int, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanInt8() (v int8, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanInt16() (v int16, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanInt32() (v int32, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanInt64() (v int64, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanUint() (v uint, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanUint8() (v uint8, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanUint16() (v uint16, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanUint32() (v uint32, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanUint64() (v uint64, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanFloat32() (v float32, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanFloat64() (v float64, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) ScanString() (v string, err error) {
	err = x.Scan(&v)
	return
}

func (x *Executor[T]) PluckBool(column string) (slice []bool, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckInt(column string) (slice []int, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckInt8(column string) (slice []int8, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckInt16(column string) (slice []int16, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckInt32(column string) (slice []int32, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckInt64(column string) (slice []int64, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckUint(column string) (slice []uint, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckUint8(column string) (slice []uint8, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckUint16(column string) (slice []uint16, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckUint32(column string) (slice []uint32, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckUint64(column string) (slice []uint64, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckString(column string) (slice []string, err error) {
	err = x.Pluck(column, &slice)
	return
}
