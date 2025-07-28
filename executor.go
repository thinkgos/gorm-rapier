package rapier

import (
	"context"

	"gorm.io/gorm"
)

type Configure[T any] func(*Executor[T]) *Executor[T]

type Executor[T any] struct {
	db     *gorm.DB
	table  Condition
	scopes []Condition
}

// Executor new executor
func NewExecutor[T any](db *gorm.DB) *Executor[T] {
	return &Executor[T]{
		db: db,
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
	x.db = x.db.Attrs(attrs...)
	return x
}

// Assign provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) Assign(attrs ...any) *Executor[T] {
	x.db = x.db.Assign(attrs...)
	return x
}

// AttrsExpr with AssignExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) AttrsExpr(attrs ...AssignExpr) *Executor[T] {
	x.db = x.db.Attrs(buildAttrsValue(attrs)...)
	return x
}

// AssignExpr with AssignExpr
// provide attributes used in [FirstOrCreate] or [FirstOrInit]
func (x *Executor[T]) AssignExpr(attrs ...AssignExpr) *Executor[T] {
	x.db = x.db.Assign(buildAttrsValue(attrs)...)
	return x
}

// Configure executor middleware
func (x *Executor[T]) Configure(cs ...Configure[T]) *Executor[T] {
	for _, f := range cs {
		x = f(x)
	}
	return x
}

// IntoDB with model or table
func (x *Executor[T]) IntoDB() *gorm.DB {
	if x.table == nil {
		x = x.Model()
	}
	x.db = x.table(x.db)
	return x.IntoRawDB()
}

// IntoRawDB without model or table
func (x *Executor[T]) IntoRawDB() *gorm.DB {
	db := x.db
	for _, f := range x.scopes {
		db = f(db)
	}
	return db
}

func (x *Executor[T]) execute(f Condition) *Executor[T] {
	return x.getInstance(f(x.db))
}

func (x *Executor[T]) getInstance(db *gorm.DB) *Executor[T] {
	x.db = db
	return x
}
