package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EqSubQuery use expr1 = (subQuery)
func (e expr) EqSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? = (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// NeqSubQuery use expr1 <> (subQuery)
func (e expr) NeqSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? <> (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// GtSubQuery use expr1 > (subQuery)
func (e expr) GtSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? > (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// GteSubQuery use expr1 >= (subQuery)
func (e expr) GteSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? >= (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// LtSubQuery use expr1 < (subQuery)
func (e expr) LtSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? < (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// LteSubQuery use expr1 <= (subQuery)
func (e expr) LteSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? <= (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// InSubQuery use expr1 IN (subQuery)
func (e expr) InSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? IN (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// NotInSubQuery use expr1 NOT IN (subQuery)
func (e expr) NotInSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "? NOT IN (?)",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// FindInSetSubQuery FIND_IN_SET(column, (subQuery))
func (e expr) FindInSetSubQuery(subQuery *gorm.DB) Expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, (?))",
		Vars: []any{e.RawExpr(), subQuery},
	}
	return e
}

// setSubQuery set with subQuery
func (e expr) SetSubQuery(subQuery *gorm.DB) SetExpr {
	e.e = clause.Set{
		clause.Assignment{
			Column: clause.Column{Name: e.col.Name},
			Value:  gorm.Expr("(?)", subQuery),
		},
	}
	return e
}
