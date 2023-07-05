package assist

import (
	"strings"

	"gorm.io/gorm/clause"
)

// EqCol use expr1 = expr2
func (e expr) EqCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? = ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// NeqCol use expr1 <> expr2
func (e expr) NeqCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? <> ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// LtCol use expr1 < expr2
func (e expr) LtCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? < ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// LteCol use expr1 <= expr2
func (e expr) LteCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? <= ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// GtCol use expr1 > expr2
func (e expr) GtCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? > ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// GteCol use expr1 >= expr2
func (e expr) GteCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "? >= ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// addCol use expr1 + expr2
func (e expr) addCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "? + ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// subCol use expr1 - expr2
func (e expr) subCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "? - ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// mulCol use (expr1) * (expr2)
func (e expr) mulCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) * (?)",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// divCol use (expr1) / (expr2)
func (e expr) divCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) / (?)",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// concatCol use CONCAT(expr1,exp2...exprN)
func (e expr) concatCol(es ...Expr) expr {
	sqlBuilder := strings.Builder{}
	sqlBuilder.Grow(8 + 2*len(es) + 1)
	vars := make([]any, 0, len(es)+1)

	sqlBuilder.WriteString("Concat(?")
	vars = append(vars, e.RawExpr())
	for _, ee := range es {
		sqlBuilder.WriteString(",?")
		vars = append(vars, ee.RawExpr())
	}
	sqlBuilder.WriteString(")")
	e.e = clause.Expr{
		SQL:  sqlBuilder.String(),
		Vars: vars,
	}
	return e
}
