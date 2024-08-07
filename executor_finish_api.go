package rapier

import "gorm.io/gorm"

type Result[T any] struct {
	RowsAffected int64
	Data         *T
}

func (x *Executor[T]) FirstOne() (*T, error) {
	var dest T

	err := x.First(&dest)
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (x *Executor[T]) TakeOne() (*T, error) {
	var dest T

	err := x.Take(&dest)
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (x *Executor[T]) LastOne() (*T, error) {
	var dest T

	err := x.Last(&dest)
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (x *Executor[T]) ScanOne() (*T, error) {
	var dest T

	err := x.Scan(&dest)
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (x *Executor[T]) Count() (count int64, err error) {
	err = x.IntoDB().Count(&count).Error
	return count, err
}

func (x *Executor[T]) First(dest any) error {
	return x.IntoDB().First(dest).Error
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

func (x *Executor[T]) PluckExpr(column Expr, value any) error {
	return x.IntoDB().Pluck(column.ColumnName(), value).Error
}

func (x *Executor[T]) Exist() (exist bool, err error) {
	err = x.IntoDB().
		Select("1").
		Limit(1).
		Scan(&exist).Error
	return exist, err
}

func (x *Executor[T]) FindAll() (rows []*T, err error) {
	err = x.Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (x *Executor[T]) FindAllByPage(offset, limit int) (rows []*T, count int64, err error) {
	count, err = x.FindByPage(&rows, offset, limit)
	return rows, count, err
}

func (x *Executor[T]) FindAllPaginate(page, perPage int64, maxPerPages ...int64) (rows []*T, count int64, err error) {
	count, err = x.FindPaginate(&rows, page, perPage, maxPerPages...)
	return rows, count, err
}

func (x *Executor[T]) Find(dest any) error {
	return x.IntoDB().Find(dest).Error
}

func (x *Executor[T]) FindByPage(dest any, offset, limit int) (count int64, err error) {
	db := x.IntoDB()
	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count > 0 {
		err = db.Offset(offset).Limit(limit).Find(dest).Error
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (x *Executor[T]) FindPaginate(dest any, page, perPage int64, maxPerPages ...int64) (count int64, err error) {
	db := x.IntoDB()
	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count > 0 {
		err = db.Scopes(Pagination(page, perPage, maxPerPages...)).Find(dest).Error
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (x *Executor[T]) FindInBatches(dest any, batchSize int, fc func(tx *gorm.DB, batch int) error) error {
	return x.IntoDB().FindInBatches(dest, batchSize, fc).Error
}

func (x *Executor[T]) Create(values ...*T) (rowsAffected int64, err error) {
	if len(values) == 0 {
		return 0, nil
	}
	res := x.IntoRawDB().Create(values)
	return res.RowsAffected, res.Error
}

func (x *Executor[T]) CreateInBatches(value []*T, batchSize int) (rowsAffected int64, err error) {
	res := x.IntoRawDB().CreateInBatches(value, batchSize)
	return res.RowsAffected, res.Error
}

func (x *Executor[T]) FirstOrInit() (Result[T], error) {
	var dest T
	var result Result[T]

	res := x.IntoDB().FirstOrInit(&dest)
	if err := res.Error; err != nil {
		return result, err
	}
	return Result[T]{RowsAffected: res.RowsAffected, Data: &dest}, nil
}

func (x *Executor[T]) FirstOrCreate() (Result[T], error) {
	var dest T
	var result Result[T]

	res := x.IntoDB().FirstOrCreate(&dest)
	if err := res.Error; err != nil {
		return result, err
	}
	return Result[T]{RowsAffected: res.RowsAffected, Data: &dest}, nil
}

func (x *Executor[T]) Save(value *T) (rowsAffected int64, err error) {
	result := x.IntoRawDB().Save(value)
	return result.RowsAffected, result.Error
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

func (x *Executor[T]) UpdateExpr(column Expr, value any) (rowsAffected int64, err error) {
	result := x.updateExpr(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdatesExpr(columns ...AssignExpr) (rowsAffected int64, err error) {
	result := x.updatesExpr(columns...)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumnExpr(column Expr, value any) (rowsAffected int64, err error) {
	result := x.updateColumnExpr(column, value)
	return result.RowsAffected, result.Error
}

func (x *Executor[T]) UpdateColumnsExpr(columns ...AssignExpr) (rowsAffected int64, err error) {
	result := x.updateColumnsExpr(columns...)
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

func (x *Executor[T]) PluckExprBool(column Expr) (slice []bool, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprInt(column Expr) (slice []int, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprInt8(column Expr) (slice []int8, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprInt16(column Expr) (slice []int16, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprInt32(column Expr) (slice []int32, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprInt64(column Expr) (slice []int64, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprUint(column Expr) (slice []uint, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprUint8(column Expr) (slice []uint8, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprUint16(column Expr) (slice []uint16, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprUint32(column Expr) (slice []uint32, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprUint64(column Expr) (slice []uint64, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

func (x *Executor[T]) PluckExprString(column Expr) (slice []string, err error) {
	err = x.PluckExpr(column, &slice)
	return
}

// ------------------------- inner finish api ---------------------------------------

func (x *Executor[T]) updateExpr(column Expr, value any) *gorm.DB {
	db := x.IntoDB()
	columnName := column.BuildColumn(db.Statement, WithoutQuote)

	switch v := value.(type) {
	case AssignExpr:
		return db.Update(columnName, v.SetExpr())
	default:
		return db.Update(columnName, v)
	}
}

func (x *Executor[T]) updatesExpr(columns ...AssignExpr) *gorm.DB {
	db := x.IntoDB()
	return db.
		Clauses(buildClauseSet(db, columns)).
		Omit("*").
		Updates(map[string]any{})
}

func (x *Executor[T]) updateColumnExpr(column Expr, value any) *gorm.DB {
	db := x.IntoDB()
	columnName := column.BuildColumn(db.Statement, WithoutQuote)
	switch v := value.(type) {
	case AssignExpr:
		return db.UpdateColumn(columnName, v.SetExpr())
	default:
		return db.UpdateColumn(columnName, v)
	}
}

func (x *Executor[T]) updateColumnsExpr(columns ...AssignExpr) *gorm.DB {
	db := x.IntoDB()
	db.Statement.SkipHooks = true
	return db.Clauses(buildClauseSet(db, columns)).
		Omit("*").
		UpdateColumns(map[string]any{})
}
