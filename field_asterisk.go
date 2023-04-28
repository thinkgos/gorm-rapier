package assist

import (
	"gorm.io/gorm/clause"
)

// Asterisk a type of xxx.*
type Asterisk struct{ expr }

// NewAsterisk new * field
func NewAsterisk(table string, opts ...Option) Asterisk {
	return Asterisk{
		expr{col: intoClauseColumn(table, "*", opts...)},
	}
}

// Count use COUNT(expr)
func (a Asterisk) Count() Asterisk {
	var expr *clause.Expr
	switch {
	case a.e != nil:
		expr = &clause.Expr{
			SQL:  "COUNT(?)",
			Vars: []any{a.e},
		}
	case a.col.Table == "":
		expr = &clause.Expr{SQL: "COUNT(*)"}
	default:
		expr = &clause.Expr{
			SQL:  "COUNT(?.*)",
			Vars: []any{clause.Table{Name: a.col.Table}},
		}
	}
	a.e = expr
	return Asterisk{expr: a.expr}
}

// Distinct use DISTINCT expr
func (a Asterisk) Distinct() Asterisk {
	var expr *clause.Expr
	if a.col.Table == "" {
		expr = &clause.Expr{SQL: "DISTINCT *"}
	} else {
		expr = &clause.Expr{
			SQL:  "DISTINCT ?.*",
			Vars: []any{clause.Table{Name: a.col.Table}},
		}
	}
	a.e = expr
	return Asterisk{expr: a.expr}
}
