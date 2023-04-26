package assist

import "gorm.io/gorm/clause"

// As use expr AS ?
func (e expr) As(alias string) Expr {
	if e.e != nil {
		e.e = clause.Expr{SQL: "? AS ?", Vars: []any{e.e, clause.Column{Name: alias}}}
		return e
	}
	e.col.Alias = alias
	return e
}

// AsWithPrefix use expr AS {prefix}_{ColumnName}
func (e expr) AsWithPrefix(prefix string) Expr {
	alias := prefix + "_" + e.col.Name
	if e.e != nil {
		e.e = clause.Expr{SQL: "? AS ?", Vars: []any{e.e, clause.Column{Name: alias}}}
		return e
	}
	e.col.Alias = alias
	return e
}

// Desc use expr DESC
func (e expr) Desc() Expr {
	e.e = clause.Expr{SQL: "? DESC", Vars: []any{e.RawExpr()}}
	return e
}
