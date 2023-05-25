package assist

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
