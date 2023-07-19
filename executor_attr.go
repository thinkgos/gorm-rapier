package assist

import "gorm.io/gorm"

type AttrExecutor struct {
	db *gorm.DB
}

// AttrsExpr with SetExpr
func (x *Executor[T]) Attrs(attrs ...any) *AttrExecutor {
	return &AttrExecutor{
		db: x.IntoDB().Attrs(attrs...),
	}
}

// AssignExpr with SetExpr
func (x *Executor[T]) Assign(attrs ...any) *AttrExecutor {
	return &AttrExecutor{
		db: x.IntoDB().Assign(attrs...),
	}
}

// AttrsExpr with SetExpr
func (x *Executor[T]) AttrsExpr(attrs ...SetExpr) *AttrExecutor {
	return &AttrExecutor{
		db: x.IntoDB().Attrs(buildAttrsValue(attrs)...),
	}
}

// AssignExpr with SetExpr
func (x *Executor[T]) AssignExpr(attrs ...SetExpr) *AttrExecutor {
	return &AttrExecutor{
		db: x.IntoDB().Assign(buildAttrsValue(attrs)...),
	}
}

func (x *AttrExecutor) FirstOrInit(dest any) (rowsAffected int64, err error) {
	result := x.db.FirstOrInit(dest)
	return result.RowsAffected, result.Error
}

func (x *AttrExecutor) FirstOrCreate(dest any) (rowsAffected int64, err error) {
	result := x.db.FirstOrCreate(dest)
	return result.RowsAffected, result.Error
}
