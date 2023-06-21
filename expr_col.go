package assist

import (
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

// EqCol use expr1 = expr2
func (e expr) EqCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? = ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// NeqCol use expr1 <> expr2
func (e expr) NeqCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? <> ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// LtCol use expr1 < expr2
func (e expr) LtCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? < ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// LteCol use expr1 <= expr2
func (e expr) LteCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? <= ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// GtCol use expr1 > expr2
func (e expr) GtCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? > ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// GteCol use expr1 >= expr2
func (e expr) GteCol(col Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? >= ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// addCol use expr1 + expr2
func (e expr) addCol(col Expr) expr {
	e.e = clause.Expr{
		SQL:  "? + ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// subCol use expr1 - expr2
func (e expr) subCol(col Expr) expr {
	e.e = clause.Expr{
		SQL:  "? - ?",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// mulCol use (expr1) * (expr2)
func (e expr) mulCol(col Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) * (?)",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// divCol use (expr1) / (expr2)
func (e expr) divCol(col Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) / (?)",
		Vars: []any{e.RawExpr(), col.RawExpr()},
	}
	return e
}

// concatCol use CONCAT(expr1,exp2...exprN)
func (e expr) concatCol(cols ...Expr) expr {
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
