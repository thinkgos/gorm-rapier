package assist

import (
	"gorm.io/gorm/clause"
)

// Asterisk a type of xxx.*
type Asterisk struct{ expr }

// NewAsterisk new * field
func NewAsterisk(table string, opts ...Option) Asterisk {
	return Asterisk{
		expr{
			col: intoClauseColumn(table, "*", opts...),
		},
	}
}

// Count use COUNT(expr)
func (a Asterisk) Count() Asterisk {
	var e clause.Expression
	switch {
	case a.e != nil:
		e = &clause.Expr{
			SQL:  "COUNT(?)",
			Vars: []any{a.e},
		}
	case a.col.Table == "":
		e = &clause.Expr{SQL: "COUNT(*)"}
	default:
		e = &clause.Expr{
			SQL:  "COUNT(?.*)",
			Vars: []any{clause.Table{Name: a.col.Table}},
		}
	}
	return Asterisk{
		expr: expr{
			col:       a.col,
			e:         e,
			buildOpts: a.buildOpts,
		},
	}
}

// Distinct use DISTINCT expr
func (a Asterisk) Distinct() Asterisk {
	var e clause.Expression
	if a.col.Table == "" {
		e = &clause.Expr{SQL: "DISTINCT *"}
	} else {
		e = &clause.Expr{
			SQL:  "DISTINCT ?.*",
			Vars: []any{clause.Table{Name: a.col.Table}},
		}
	}
	return Asterisk{
		expr: expr{
			col:       a.col,
			e:         e,
			buildOpts: a.buildOpts,
		},
	}
}
