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
	err = x.chains().Count(&count).Error
	return count, err
}

func (x *Executor[T]) First(dest any) error {
	return x.chains().First(dest).Error
}

func (x *Executor[T]) FirstOrInit(dest any) error {
	return x.chains().FirstOrInit(dest).Error
}

func (x *Executor[T]) FirstOrCreate(dest any) error {
	return x.chains().FirstOrCreate(dest).Error
}

func (x *Executor[T]) Take(dest any) error {
	return x.chains().Take(dest).Error
}

func (x *Executor[T]) Last(dest any) error {
	return x.chains().Last(dest).Error
}

func (x *Executor[T]) Scan(dest any) error {
	return x.chains().Scan(dest).Error
}

func (x *Executor[T]) Pluck(column string, value any) error {
	return x.chains().Pluck(column, value).Error
}

func (x *Executor[T]) Exist() (exist bool, err error) {
	err = x.chains().
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
	return x.chains().Find(dest).Error
}

func (x *Executor[T]) FindInBatches(dest any, batchSize int, fc func(tx *gorm.DB, batch int) error) error {
	return x.chains().FindInBatches(dest, batchSize, fc).Error
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
	result := x.chains().Updates(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) Update(column string, value any) (rowsAffected int64, err error) {
	result := x.chains().Update(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumns(value *T) (rowsAffected int64, err error) {
	result := x.chains().UpdateColumns(value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumn(column string, value any) (rowsAffected int64, err error) {
	result := x.chains().UpdateColumn(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) Delete() (rowsAffected int64, err error) {
	var t T

	result := x.chains().Delete(&t)
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

func (x *Executor[T]) FirstInt64() (v int64, err error) {
	err = x.First(&v)
	return
}

func (x *Executor[T]) FirstFloat64() (v Float64, err error) {
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

func (x *Executor[T]) TakeInt64() (v int64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) TakeFloat64() (v Float64, err error) {
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
	err = x.Take(&v)
	return
}

func (x *Executor[T]) ScanInt64() (v int64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) ScanFloat64() (v Float64, err error) {
	err = x.Take(&v)
	return
}

func (x *Executor[T]) ScanString() (v string, err error) {
	err = x.Take(&v)
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

func (x *Executor[T]) PluckInt64(column string) (slice []int64, err error) {
	err = x.Pluck(column, &slice)
	return
}

func (x *Executor[T]) PluckString(column string) (slice []string, err error) {
	err = x.Pluck(column, &slice)
	return
}
