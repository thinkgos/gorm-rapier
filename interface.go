package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SetExpr = AssignExpr

// Expr a query expression about field
type Expr interface {
	clause.Expression

	As(alias string) Expr
	ColumnName() string
	Expression() clause.Expression
	RawExpr() any
	BuildColumn(*gorm.Statement, ...BuildOption) string
	BuildWithArgs(*gorm.Statement) (query string, args []any)
}

type AssignExpr interface {
	Expr

	SetExpr() any
}
