package assist

import (
	"gorm.io/gorm/clause"
)

func (e expr) value(value any) AssignExpr {
	e.e = clause.Eq{
		Column: e.col.Name,
		Value:  value,
	}
	return e
}
