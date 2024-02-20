package rapier

import "gorm.io/gorm/clause"

// As use expr AS ?
func (e expr) As(alias string) Expr {
	if e.e != nil {
		e.e = clause.Expr{
			SQL:  "? AS ?",
			Vars: []any{e.e, clause.Column{Name: alias}},
		}
		return e
	}
	e.col.Alias = alias
	return e
}

// Desc use expr DESC
func (e expr) Desc() Expr {
	e.e = clause.Expr{
		SQL:  "? DESC",
		Vars: []any{e.RawExpr()},
	}
	return e
}

// Asc use expr ASC
func (e expr) Asc() Expr {
	e.e = clause.Expr{
		SQL:  "? ASC",
		Vars: []any{e.RawExpr()},
	}
	return e
}
