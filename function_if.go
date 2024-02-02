package rapier

import "gorm.io/gorm/clause"

func IF(condition, value_if_true, value_if_false Expr) Field {
	return Field{
		expr: expr{
			e: clause.Expr{
				SQL:                "IF(?,?,?)",
				Vars:               []any{condition, value_if_true, value_if_false},
				WithoutParentheses: false,
			},
		},
	}
}
