package assist

import (
	"gorm.io/gorm/clause"
)

// valueEq same as value, but use clause.Eq
func (e expr) valueEq(value any) SetExpr {
	e.e = clause.Eq{
		Column: e.col.Name,
		Value:  value,
	}
	return e
}

// value set value, Clause.Set
func (e expr) value(value any) SetExpr {
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
func (e expr) ValueNull() SetExpr {
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

// ValueAny set any value.
// Deprecated: use other ValueXXX instead.
func (e expr) ValueAny(value any) SetExpr {
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
