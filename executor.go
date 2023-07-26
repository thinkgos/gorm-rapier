package assist

import (
	"context"

	"gorm.io/gorm"
)

type Executor[T any] struct {
	db         *gorm.DB
	table      Condition
	attrs      Condition // for [FirstOrInit|FirstOrCreate]
	assigns    Condition // for [FirstOrInit|FirstOrCreate]
	conditions *Conditions
}

// Executor new executor
func NewExecutor[T any](db *gorm.DB) *Executor[T] {
	return &Executor[T]{
		db:         db,
		table:      nil,
		attrs:      nil,
		assigns:    nil,
		conditions: NewConditions(),
	}
}

func (x *Executor[T]) Session(config *gorm.Session) *Executor[T] {
	x.db = x.db.Session(config)
	return x
}

func (x *Executor[T]) WithContext(ctx context.Context) *Executor[T] {
	x.db = x.db.WithContext(ctx)
	return x
}

func (x *Executor[T]) Debug() *Executor[T] {
	x.db = x.db.Debug()
	return x
}

// Attrs provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) Attrs(attrs ...any) *Executor[T] {
	x.attrs = innerAttrs(attrs...)
	return x
}

// AttrsExpr  provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) Assign(attrs ...any) *Executor[T] {
	x.assigns = innerAssign(attrs...)
	return x
}

// AttrsExpr with SetExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) AttrsExpr(attrs ...SetExpr) *Executor[T] {
	x.attrs = innerAttrsExpr(attrs...)
	return x
}

// AssignExpr with SetExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) AssignExpr(attrs ...SetExpr) *Executor[T] {
	x.assigns = innerAssignExpr(attrs...)
	return x
}

// IntoDB with model or table
func (x *Executor[T]) IntoDB() *gorm.DB {
	if x.table == nil {
		var model T

		x.db = innerModel(model)(x.db)
	} else {
		x.db = x.table(x.db)
	}
	return x.IntoRawDB()
}

// IntoRawDB without model or table
func (x *Executor[T]) IntoRawDB() *gorm.DB {
	db := x.db
	for _, f := range x.conditions.Build() {
		db = f(db)
	}
	return db
}

/****************************************************************************/

func innerAttrs(attrs ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Attrs(attrs...)
	}
}

func innerAttrsExpr(attrs ...SetExpr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Attrs(buildAttrsValue(attrs)...)
	}
}

func innerAssign(attrs ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Assign(attrs...)
	}
}

func innerAssignExpr(attrs ...SetExpr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Assign(buildAttrsValue(attrs)...)
	}
}

func innerModel[T any](t T) Condition {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Model(&t)
		err := db.Statement.Parse(t)
		if err != nil {
			_ = db.AddError(err)
		}
		return db
	}
}
