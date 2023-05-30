package assist

import (
	"context"

	"gorm.io/gorm"
)

type Executor[T any] struct {
	db    *gorm.DB
	table func(*gorm.DB) *gorm.DB
	funcs []func(*gorm.DB) *gorm.DB
}

// Executor new executor
func NewExecutor[T any](db *gorm.DB) *Executor[T] {
	return &Executor[T]{
		db:    db,
		table: nil,
		funcs: make([]func(*gorm.DB) *gorm.DB, 0, 16),
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

func (x *Executor[T]) IntoDB() (db *gorm.DB) {
	if x.table == nil {
		var t T

		db = x.db.Model(&t)
	} else {
		db = x.db.Scopes(x.table)
	}
	return db.Scopes(x.funcs...)
}
