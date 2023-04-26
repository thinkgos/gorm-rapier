package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type expression any

// Expr a query expression about field
type Expr interface {
	clause.Expression

	As(alias string) Expr
	AsWithPrefix(prefix string) Expr
	ColumnName() string
	Expression() clause.Expression
	RawExpr() expression
	BuildColumn(*gorm.Statement, ...BuildOption) string
	BuildWithArgs(*gorm.Statement) (query string, args []any)

	// col operate expression
	AddCol(col Expr) Expr
	SubCol(col Expr) Expr
	MulCol(col Expr) Expr
	DivCol(col Expr) Expr
	ConcatCol(cols ...Expr) Expr
}
