package assist

import "gorm.io/gorm"

type AttrExecutor struct {
	db *gorm.DB
}

// IntoAttrExecutor into AttrExecutor
// [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) IntoAttrExecutor() *AttrExecutor {
	return &AttrExecutor{db: x.IntoDB()}
}

// Attrs provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *AttrExecutor) Attrs(attrs ...any) *AttrExecutor {
	x.db = x.db.Attrs(attrs...)
	return x
}

// Assign provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *AttrExecutor) Assign(attrs ...any) *AttrExecutor {
	x.db = x.db.Assign(attrs...)
	return x
}

// AttrsExpr with SetExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *AttrExecutor) AttrsExpr(attrs ...SetExpr) *AttrExecutor {
	x.db = x.db.Attrs(buildAttrsValue(attrs)...)
	return x
}

// AssignExpr with SetExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *AttrExecutor) AssignExpr(attrs ...SetExpr) *AttrExecutor {
	x.db = x.db.Assign(buildAttrsValue(attrs)...)
	return x
}

func (x *AttrExecutor) FirstOrInit(dest any) (rowsAffected int64, err error) {
	result := x.db.FirstOrInit(dest)
	return result.RowsAffected, result.Error
}

func (x *AttrExecutor) FirstOrCreate(dest any) (rowsAffected int64, err error) {
	result := x.db.FirstOrCreate(dest)
	return result.RowsAffected, result.Error
}
