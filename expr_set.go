package assist

import (
	"gorm.io/gorm/clause"
)

// setValue same sa value, but use clause.Eq
func (e expr) setValue(value any) SetExpr {
	e.e = clause.Eq{
		Column: e.col.Name,
		Value:  value,
	}
	return e
}

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
