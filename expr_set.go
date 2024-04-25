package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// valueEq same as value, but use clause.Eq
func (e expr) valueEq(value any) AssignExpr {
	e.e = clause.Eq{
		Column: e.col.Name,
		Value:  value,
	}
	return e
}

// value set value, Clause.Set
func (e expr) value(value any) AssignExpr {
	e.e = clause.Set{
		{
			Column: clause.Column{
				Name: e.col.Name,
			},
			Value: value,
		},
	}
	return e
}

// ValueNull set value NULL
func (e expr) ValueNull() AssignExpr {
	e.e = clause.Set{
		{
			Column: clause.Column{
				Name: e.col.Name,
			},
			Value: nil,
		},
	}
	return e
}

// ValueSubQuery set with subQuery
// same as SetSubQuery.
func (e expr) ValueSubQuery(subQuery *gorm.DB) AssignExpr {
	return e.SetSubQuery(subQuery)
}

// ValueAny set any value.
func (e expr) ValueAny(value any) AssignExpr {
	e.e = clause.Set{
		{
			Column: clause.Column{
				Name: e.col.Name,
			},
			Value: value,
		},
	}
	return e
}
