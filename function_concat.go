package rapier

import (
	"strings"

	"gorm.io/gorm/clause"
)

// CONCAT(expr1,expr2,...exprN)
func ConcatCol(e Expr, es ...Expr) Field {
	return Field{
		expr: expr{
			e: concatCol(e, es...),
		},
	}
}

// CONCAT_WS(separator,expr1,expr2,...exprN)
func ConcatWsCol(separator Expr, e Expr, es ...Expr) Field {
	return Field{
		expr: expr{
			e: concatWsCol(separator, e, es...),
		},
	}
}

// CONCAT(expr1,expr2,...exprN)
func concatCol(e Expr, es ...Expr) expr {
	sqlBuilder := strings.Builder{}
	sqlBuilder.Grow(8 + 2*len(es) + 1)
	vars := make([]any, 0, len(es)+1)

	sqlBuilder.WriteString("CONCAT(?")
	vars = append(vars, e.RawExpr())
	for _, ee := range es {
		sqlBuilder.WriteString(",?")
		vars = append(vars, ee.RawExpr())
	}
	sqlBuilder.WriteString(")")
	return expr{
		e: clause.Expr{
			SQL:  sqlBuilder.String(),
			Vars: vars,
		},
	}
}

// CONCAT_WS(separator,expr1,expr2,...exprN)
func concatWsCol(separator Expr, e Expr, es ...Expr) expr {
	sqlBuilder := strings.Builder{}
	sqlBuilder.Grow(13 + 2*len(es) + 1)
	vars := make([]any, 0, 1+len(es)+1)

	sqlBuilder.WriteString("CONCAT_WS(?,?")
	vars = append(vars, separator.RawExpr(), e.RawExpr())
	for _, ee := range es {
		sqlBuilder.WriteString(",?")
		vars = append(vars, ee.RawExpr())
	}
	sqlBuilder.WriteString(")")
	return expr{
		e: clause.Expr{
			SQL:  sqlBuilder.String(),
			Vars: vars,
		},
	}
}
