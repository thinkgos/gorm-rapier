package rapier

import "gorm.io/gorm/clause"

// IsNull use expr IS NULL
func (e expr) IsNull() Expr {
	e.e = clause.Expr{
		SQL:  "? IS NULL",
		Vars: []any{e.RawExpr()},
	}
	return e
}

// IsNotNull use expr IS NOT NULL
func (e expr) IsNotNull() Expr {
	e.e = clause.Expr{
		SQL:  "? IS NOT NULL",
		Vars: []any{e.RawExpr()},
	}
	return e
}

// Count use COUNT(expr)
func (e expr) Count() Int {
	e.e = clause.Expr{
		SQL:  "COUNT(?)",
		Vars: []any{e.RawExpr()},
	}
	return Int{e}
}

// Distinct use DISTINCT(expr)
func (e expr) Distinct() Int {
	e.e = clause.Expr{
		SQL:  "DISTINCT ?",
		Vars: []any{e.RawExpr()},
	}
	return Int{e}
}

// Length use LENGTH(expr)
func (e expr) Length() Int {
	e.e = clause.Expr{
		SQL:  "LENGTH(?)",
		Vars: []any{e.RawExpr()},
	}
	return Int{e}
}

// Max use MAX(expr)
func (e expr) Max() Float64 {
	e.e = clause.Expr{
		SQL:  "MAX(?)",
		Vars: []any{e.RawExpr()},
	}
	return Float64{e}
}

// Min use MIN(expr)
func (e expr) Min() Float64 {
	e.e = clause.Expr{
		SQL:  "MIN(?)",
		Vars: []any{e.RawExpr()},
	}
	return Float64{e}
}

// Avg use AVG(expr)
func (e expr) Avg() Float64 {
	e.e = clause.Expr{
		SQL:  "AVG(?)",
		Vars: []any{e.RawExpr()},
	}
	return Float64{e}
}

// GroupConcat use GROUP_CONCAT(expr)
func (e expr) GroupConcat() Expr {
	e.e = clause.Expr{
		SQL:  "GROUP_CONCAT(?)",
		Vars: []any{e.RawExpr()},
	}
	return e
}
