package rapier

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

// FindInSetCol use FIND_IN_SET(expr1, expr2)
func (e expr) FindInSetCol(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// FindInSetColWith use FIND_IN_SET(expr2, expr1)
func (e expr) FindInSetColWith(e2 Expr) Expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{e2.RawExpr(), e.RawExpr()},
	}
	return e
}

// SetCol expr1=expr2
func (e expr) SetCol(e2 Expr) SetExpr {
	e.e = clause.Set{
		{
			Column: clause.Column{
				Name: e.col.Name,
			},
			Value: e2.RawExpr(),
		},
	}
	return e
}

// expr1 + expr2
func (e expr) innerAddCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "? + ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// expr1 - expr2
func (e expr) innerSubCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "? - ?",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// use (expr1) * (expr2)
func (e expr) innerMulCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) * (?)",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// (expr1) / (expr2)
func (e expr) innerDivCol(e2 Expr) expr {
	e.e = clause.Expr{
		SQL:  "(?) / (?)",
		Vars: []any{e.RawExpr(), e2.RawExpr()},
	}
	return e
}

// CONCAT(expr1,exp2...exprN)
func (e expr) innerConcatCol(es ...Expr) expr {
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
