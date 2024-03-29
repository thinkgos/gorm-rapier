package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

type SetExpr interface {
	Expr

	SetExpr() any
}

type AssignExpr = SetExpr
