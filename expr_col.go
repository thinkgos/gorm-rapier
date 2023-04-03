package assist

import (
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

// EqCol use expr1 = expr2
func (e expr) EqCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? = ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// NeqCol use expr1 <> expr2
func (e expr) NeqCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? <> ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// LtCol use expr1 < expr2
func (e expr) LtCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? < ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// LteCol use expr1 <= expr2
func (e expr) LteCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? <= ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// GtCol use expr1 > expr2
func (e expr) GtCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? > ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// GteCol use expr1 >= expr2
func (e expr) GteCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? >= ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// AddCol use expr1 + expr2
func (e expr) AddCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? + ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// SubCol use expr1 - expr2
func (e expr) SubCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "? - ?", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// MulCol use (expr1) * (expr2)
func (e expr) MulCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "(?) * (?)", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// DivCol use (expr1) / (expr2)
func (e expr) DivCol(col Expr) Expr {
	e.e = clause.Expr{SQL: "(?) / (?)", Vars: []any{e.RawExpr(), col.RawExpr()}}
	return e
}

// ConcatCol use CONCAT(expr1,exp2...exprN)
func (e expr) ConcatCol(cols ...Expr) Expr {
	placeholders := []string{"?"}
	vars := []any{e.RawExpr()}
	for _, col := range cols {
		placeholders = append(placeholders, "?")
		vars = append(vars, col.RawExpr())
	}
	e.e = clause.Expr{
		SQL:  fmt.Sprintf("Concat(%s)", strings.Join(placeholders, ",")),
		Vars: vars,
	}
	return e
}
